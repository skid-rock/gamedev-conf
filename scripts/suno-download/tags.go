package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/abema/go-mp4"
)

// iTunes-типы данных в атоме data (QTFF). В go-mp4 JPEG назван 14,
// у Apple JPEG=13, PNG=14 — пишем числа по спецификации, не по имени константы.
const (
	dataTypeUTF8 = mp4.DataTypeStringUTF8
	dataTypeJPEG = 13
	dataTypePNG  = 14
)

var (
	boxTypeNam  = mp4.BoxType{0xA9, 'n', 'a', 'm'}
	boxTypeART  = mp4.BoxType{0xA9, 'A', 'R', 'T'}
	boxTypeGen  = mp4.BoxType{0xA9, 'g', 'e', 'n'}
	boxTypeDay  = mp4.BoxType{0xA9, 'd', 'a', 'y'}
	boxTypeCmt  = mp4.BoxType{0xA9, 'c', 'm', 't'}
	boxTypeCovr = mp4.StrToBoxType("covr")

	errFFmpegMissing = errors.New("ffmpeg не найден")
)

// Атомы, которые перезаписываем своими значениями (©too и прочее не трогаем).
func replaceableIlstType(t mp4.BoxType) bool {
	switch t {
	case boxTypeNam, boxTypeART, boxTypeGen, boxTypeDay, boxTypeCmt, boxTypeCovr:
		return true
	}
	return false
}

type itunesTags struct {
	Title   string
	Artist  string
	Genre   string
	Date    string
	Comment string
	Cover   []byte
}

func (c clipJSON) modelLabel() string {
	return strings.TrimSpace(strings.TrimSpace(c.MajorModelVersion) + " " + strings.TrimSpace(c.ModelName))
}

func mergeComment(existing, model, created, id string) string {
	c := strings.TrimSpace(existing)
	if c == "" {
		parts := []string{"made with suno"}
		if created != "" {
			parts = append(parts, "created="+created)
		}
		if id != "" {
			parts = append(parts, "id="+id)
		}
		if model != "" {
			parts = append(parts, "model="+model)
		}
		return strings.Join(parts, "; ")
	}
	if model != "" && !strings.Contains(strings.ToLower(c), "model=") {
		c = strings.TrimRight(c, "; ") + "; model=" + model
	}
	return c
}

func fetchCover(imageURL string, ref clipRef, c creds) ([]byte, error) {
	status, body, err := httpRequest(http.MethodGet, imageURL, nil, baseHeaders(ref, c))
	if err != nil {
		return nil, fmt.Errorf("не скачалась обложка: %w", err)
	}
	if status != 200 {
		return nil, fmt.Errorf("обложка HTTP %d", status)
	}
	if len(body) == 0 {
		return nil, errors.New("обложка пустая")
	}
	return body, nil
}

func coverDataType(img []byte) uint32 {
	if len(img) >= 3 && img[0] == 0xff && img[1] == 0xd8 && img[2] == 0xff {
		return dataTypeJPEG
	}
	if len(img) >= 8 && string(img[:8]) == "\x89PNG\r\n\x1a\n" {
		return dataTypePNG
	}
	return mp4.DataTypeBinary
}

func readIlstComment(r io.ReadSeeker) (string, error) {
	var comment string
	_, err := mp4.ReadBoxStructure(r, func(h *mp4.ReadHandle) (interface{}, error) {
		switch h.BoxInfo.Type {
		case mp4.BoxTypeMoov(), mp4.BoxTypeUdta(), mp4.BoxTypeMeta(), mp4.BoxTypeIlst(), boxTypeCmt:
			_, err := h.Expand()
			return nil, err
		case mp4.BoxTypeData():
			if !h.BoxInfo.Context.UnderIlstMeta || !pathEndsWith(h.Path, boxTypeCmt) {
				return nil, nil
			}
			box, _, err := h.ReadPayload()
			if err != nil {
				return nil, err
			}
			if d, ok := box.(*mp4.Data); ok {
				comment = string(d.Data)
			}
			return nil, nil
		default:
			return nil, nil
		}
	})
	return comment, err
}

func pathEndsWith(path mp4.BoxPath, parent mp4.BoxType) bool {
	if len(path) < 2 {
		return false
	}
	return path[len(path)-2] == parent
}

func writeTaggedM4A(plaintext []byte, outPath string, tags itunesTags) error {
	if err := writeItunesTagsGo(plaintext, outPath, tags); err == nil {
		return nil
	} else if ferr := writeItunesTagsFFmpeg(plaintext, outPath, tags); ferr != nil {
		if errors.Is(ferr, errFFmpegMissing) {
			return err
		}
		return fmt.Errorf("атомы ilst (go-mp4): %v; запасной ffmpeg: %w", err, ferr)
	}
	fmt.Println("теги: go-mp4 не вышло, дописал ffmpeg -c copy")
	return nil
}

func writeItunesTagsGo(plaintext []byte, outPath string, tags itunesTags) error {
	dir := filepath.Dir(outPath)
	tmp, err := os.CreateTemp(dir, "suno-tag-*.m4a")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	ok := false
	defer func() {
		tmp.Close()
		if !ok {
			os.Remove(tmpName)
		}
	}()

	r := bytes.NewReader(plaintext)
	w := mp4.NewWriter(tmp)
	var ilstExists bool
	var mdatOffsetDiff int64
	var stcoOffsets []int64
	var co64Offsets []int64

	_, err = mp4.ReadBoxStructure(r, func(h *mp4.ReadHandle) (interface{}, error) {
		switch h.BoxInfo.Type {
		case mp4.BoxTypeMoov(),
			mp4.BoxTypeTrak(),
			mp4.BoxTypeMdia(),
			mp4.BoxTypeMinf(),
			mp4.BoxTypeStbl(),
			mp4.BoxTypeUdta(),
			mp4.BoxTypeIlst():
			if _, err := w.StartBox(&h.BoxInfo); err != nil {
				return nil, err
			}
			if _, err := h.Expand(); err != nil {
				return nil, err
			}
			if h.BoxInfo.Type == mp4.BoxTypeIlst() {
				ctx := h.BoxInfo.Context
				ctx.UnderIlst = true
				if err := writeIlstItems(w, ctx, tags); err != nil {
					return nil, err
				}
				ilstExists = true
			}
			if h.BoxInfo.Type == mp4.BoxTypeMoov() && !ilstExists {
				if err := writeNewUdtaIlst(w, h.BoxInfo.Context, tags); err != nil {
					return nil, err
				}
				ilstExists = true
			}
			_, err := w.EndBox()
			return nil, err

		case mp4.BoxTypeMeta():
			if _, err := w.StartBox(&h.BoxInfo); err != nil {
				return nil, err
			}
			box, n, err := h.ReadPayload()
			if err != nil {
				return nil, err
			}
			if n > 0 {
				if _, err := mp4.Marshal(w, box, h.BoxInfo.Context); err != nil {
					return nil, err
				}
			}
			if _, err := h.Expand(); err != nil {
				return nil, err
			}
			_, err = w.EndBox()
			return nil, err

		default:
			if replaceableIlstType(h.BoxInfo.Type) && h.BoxInfo.Context.UnderIlst && !h.BoxInfo.Context.UnderIlstMeta {
				return nil, nil
			}
			switch h.BoxInfo.Type {
			case mp4.BoxTypeStco():
				off, err := w.Seek(0, io.SeekCurrent)
				if err != nil {
					return nil, err
				}
				stcoOffsets = append(stcoOffsets, off)
			case mp4.BoxTypeCo64():
				off, err := w.Seek(0, io.SeekCurrent)
				if err != nil {
					return nil, err
				}
				co64Offsets = append(co64Offsets, off)
			case mp4.BoxTypeMdat():
				iOffset := int64(h.BoxInfo.Offset)
				oOffset, err := w.Seek(0, io.SeekCurrent)
				if err != nil {
					return nil, err
				}
				mdatOffsetDiff = oOffset - iOffset
			}
			return nil, w.CopyBox(r, &h.BoxInfo)
		}
	})
	if err != nil {
		return err
	}
	if mdatOffsetDiff != 0 {
		if err := patchChunkOffsets(tmp, stcoOffsets, co64Offsets, mdatOffsetDiff); err != nil {
			return err
		}
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmpName, outPath); err != nil {
		return err
	}
	ok = true
	return nil
}

func writeNewUdtaIlst(w *mp4.Writer, ctx mp4.Context, tags itunesTags) error {
	if _, err := w.StartBox(&mp4.BoxInfo{Type: mp4.BoxTypeUdta()}); err != nil {
		return err
	}
	if _, err := w.StartBox(&mp4.BoxInfo{Type: mp4.BoxTypeMeta()}); err != nil {
		return err
	}
	ctx.UnderUdta = true
	if _, err := mp4.Marshal(w, &mp4.Meta{}, ctx); err != nil {
		return err
	}
	if _, err := w.StartBox(&mp4.BoxInfo{Type: mp4.BoxTypeHdlr()}); err != nil {
		return err
	}
	hdlr := &mp4.Hdlr{HandlerType: [4]byte{'m', 'd', 'i', 'r'}}
	if _, err := mp4.Marshal(w, hdlr, ctx); err != nil {
		return err
	}
	if _, err := w.EndBox(); err != nil {
		return err
	}
	if _, err := w.StartBox(&mp4.BoxInfo{Type: mp4.BoxTypeIlst()}); err != nil {
		return err
	}
	ctx.UnderIlst = true
	if err := writeIlstItems(w, ctx, tags); err != nil {
		return err
	}
	if _, err := w.EndBox(); err != nil {
		return err
	}
	if _, err := w.EndBox(); err != nil {
		return err
	}
	_, err := w.EndBox()
	return err
}

func writeIlstItems(w *mp4.Writer, ctx mp4.Context, tags itunesTags) error {
	pairs := []struct {
		typ  mp4.BoxType
		text string
	}{
		{boxTypeNam, tags.Title},
		{boxTypeART, tags.Artist},
		{boxTypeGen, tags.Genre},
		{boxTypeDay, tags.Date},
		{boxTypeCmt, tags.Comment},
	}
	for _, p := range pairs {
		if strings.TrimSpace(p.text) == "" {
			continue
		}
		if err := writeIlstData(w, ctx, p.typ, dataTypeUTF8, []byte(p.text)); err != nil {
			return err
		}
	}
	if len(tags.Cover) == 0 {
		return nil
	}
	return writeIlstData(w, ctx, boxTypeCovr, coverDataType(tags.Cover), tags.Cover)
}

func writeIlstData(w *mp4.Writer, ctx mp4.Context, typ mp4.BoxType, dataType uint32, payload []byte) error {
	if _, err := w.StartBox(&mp4.BoxInfo{Type: typ}); err != nil {
		return err
	}
	if _, err := w.StartBox(&mp4.BoxInfo{Type: mp4.BoxTypeData()}); err != nil {
		return err
	}
	dataCtx := ctx
	dataCtx.UnderIlstMeta = true
	boxData := &mp4.Data{
		DataType: dataType,
		Data:     payload,
	}
	if _, err := mp4.Marshal(w, boxData, dataCtx); err != nil {
		return err
	}
	if _, err := w.EndBox(); err != nil {
		return err
	}
	_, err := w.EndBox()
	return err
}

func patchChunkOffsets(f *os.File, stcoOffsets, co64Offsets []int64, diff int64) error {
	for _, off := range stcoOffsets {
		if _, err := f.Seek(off, io.SeekStart); err != nil {
			return err
		}
		bi, err := mp4.ReadBoxInfo(f)
		if err != nil {
			return err
		}
		var stco mp4.Stco
		if _, err := mp4.Unmarshal(f, bi.Size-bi.HeaderSize, &stco, bi.Context); err != nil {
			return err
		}
		for i, v := range stco.ChunkOffset {
			n := int64(v) + diff
			if n < 0 || n > int64(^uint32(0)) {
				return fmt.Errorf("stco: смещение %d не влезает в uint32", n)
			}
			stco.ChunkOffset[i] = uint32(n)
		}
		if _, err := bi.SeekToPayload(f); err != nil {
			return err
		}
		if _, err := mp4.Marshal(f, &stco, bi.Context); err != nil {
			return err
		}
	}
	for _, off := range co64Offsets {
		if _, err := f.Seek(off, io.SeekStart); err != nil {
			return err
		}
		bi, err := mp4.ReadBoxInfo(f)
		if err != nil {
			return err
		}
		var co64 mp4.Co64
		if _, err := mp4.Unmarshal(f, bi.Size-bi.HeaderSize, &co64, bi.Context); err != nil {
			return err
		}
		for i, v := range co64.ChunkOffset {
			if diff >= 0 {
				co64.ChunkOffset[i] = v + uint64(diff)
			} else {
				d := uint64(-diff)
				if v < d {
					return fmt.Errorf("co64: смещение ушло в минус")
				}
				co64.ChunkOffset[i] = v - d
			}
		}
		if _, err := bi.SeekToPayload(f); err != nil {
			return err
		}
		if _, err := mp4.Marshal(f, &co64, bi.Context); err != nil {
			return err
		}
	}
	return nil
}

func writeItunesTagsFFmpeg(plaintext []byte, outPath string, tags itunesTags) error {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return errFFmpegMissing
	}
	dir := filepath.Dir(outPath)
	inFile, err := os.CreateTemp(dir, "suno-in-*.m4a")
	if err != nil {
		return err
	}
	inName := inFile.Name()
	defer os.Remove(inName)
	if _, err := inFile.Write(plaintext); err != nil {
		inFile.Close()
		return err
	}
	if err := inFile.Close(); err != nil {
		return err
	}

	args := []string{"-y", "-i", inName}
	var coverName string
	if len(tags.Cover) > 0 {
		cf, err := os.CreateTemp(dir, "suno-cover-*.jpg")
		if err != nil {
			return err
		}
		coverName = cf.Name()
		defer os.Remove(coverName)
		if _, err := cf.Write(tags.Cover); err != nil {
			cf.Close()
			return err
		}
		if err := cf.Close(); err != nil {
			return err
		}
		args = append(args, "-i", coverName, "-map", "0:a", "-map", "1", "-c", "copy", "-disposition:v:0", "attached_pic")
	} else {
		args = append(args, "-c", "copy")
	}
	addMeta := func(key, val string) {
		if strings.TrimSpace(val) == "" {
			return
		}
		args = append(args, "-metadata", key+"="+val)
	}
	addMeta("title", tags.Title)
	addMeta("artist", tags.Artist)
	addMeta("genre", tags.Genre)
	addMeta("date", tags.Date)
	addMeta("comment", tags.Comment)
	args = append(args, outPath)

	cmd := exec.Command("ffmpeg", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		snip := out
		if len(snip) > 400 {
			snip = snip[:400]
		}
		return fmt.Errorf("ffmpeg: %w: %s", err, snip)
	}
	return nil
}
