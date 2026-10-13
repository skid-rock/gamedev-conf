package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path"
	"regexp"
	"strings"
)

// Склейка: результат строится поверх первого файла, слайды остальных переносятся в конец.
type merger struct {
	out    *pkg
	outCT  *contentTypes
	target slideSize
	scale  bool

	// Что уже перенесено, чтобы не копировать дважды.
	mediaByHash map[string]string // sha256 содержимого -> путь в результате
	copied      map[string]string // "<файл>\x00<часть>" -> путь в результате

	notesMaster string // мастер заметок результата; заметки всех источников ведут на него

	st stats
}

// Итоги работы — для отчёта.
type stats struct {
	slides      int
	layouts     int
	masters     int
	notes       int
	mediaCopied int
	mediaReused int
	scaledAttrs int
	warnings    []string
}

func (s *stats) warn(format string, a ...any) {
	s.warnings = append(s.warnings, fmt.Sprintf(format, a...))
}

func newMerger(base *pkg, target slideSize, scale bool) (*merger, error) {
	out := base.clone()
	ct, err := readContentTypes(out)
	if err != nil {
		return nil, err
	}
	nm, err := notesMasterOf(out)
	if err != nil {
		return nil, err
	}
	return &merger{
		out:         out,
		outCT:       ct,
		target:      target,
		scale:       scale,
		mediaByHash: map[string]string{},
		copied:      map[string]string{},
		notesMaster: nm,
	}, nil
}

// Коэффициент пересчёта размеров для источника.
func (m *merger) factor(src slideSize) float64 {
	if !m.scale || src.cx == 0 || src.cx == m.target.cx {
		return 1
	}
	return float64(m.target.cx) / float64(src.cx)
}

// Привести слайды самого первого файла к целевому размеру.
func (m *merger) rescaleBase(from slideSize) error {
	k := m.factor(from)
	if k == 1 {
		return nil
	}
	for _, part := range m.out.order {
		if !kindOf(m.outCT, part).scalable() {
			continue
		}
		data, n, err := scaleXML(m.out.parts[part], k)
		if err != nil {
			return fmt.Errorf("%s: %w", part, err)
		}
		m.out.set(part, data)
		m.st.scaledAttrs += n
	}
	return nil
}

// Перенести все слайды одного источника.
func (m *merger) mergeFile(src *pkg, srcSize slideSize) error {
	k := m.factor(srcSize)
	srcCT, err := readContentTypes(src)
	if err != nil {
		return err
	}
	slides, err := slideOrder(src)
	if err != nil {
		return err
	}
	if len(slides) == 0 {
		m.st.warn("%s: не нашлось ни одного слайда", path.Base(src.src))
	}
	for _, slide := range slides {
		newName, err := m.copyPart(src, srcCT, slide, k)
		if err != nil {
			return fmt.Errorf("%s: %w", slide, err)
		}
		if err := registerInPresentation(m.out, newName, relSlide, "sldIdLst", "sldId"); err != nil {
			return err
		}
		m.st.slides++
	}
	return nil
}

// Скопировать часть вместе со всем, на что она ссылается.
// Имя в результате резервируется до обхода связей: OOXML допускает циклы
// (макет ссылается на образец, образец — обратно на свои макеты).
func (m *merger) copyPart(src *pkg, srcCT *contentTypes, part string, k float64) (string, error) {
	key := src.src + "\x00" + part
	if name, ok := m.copied[key]; ok {
		return name, nil
	}
	if !src.has(part) {
		return "", fmt.Errorf("в исходнике нет части %s", part)
	}
	kind := kindOf(srcCT, part)
	if kind == kindMedia {
		return m.copyMedia(src, srcCT, part)
	}
	// Заметки всех источников ведут на мастер заметок результата:
	// presentation.xml допускает только один notesMaster.
	if kind == kindNotesMaster && m.notesMaster != "" {
		m.copied[key] = m.notesMaster
		return m.notesMaster, nil
	}

	dir, base, ext := splitPartName(part)
	newName := m.out.freeName(dir, base, ext)
	m.copied[key] = newName
	m.out.set(newName, nil) // бронь имени на время обхода связей

	data := src.parts[part]
	if k != 1 && kind.scalable() {
		scaled, n, err := scaleXML(data, k)
		if err != nil {
			return "", fmt.Errorf("пересчёт размеров: %w", err)
		}
		data = scaled
		m.st.scaledAttrs += n
	}
	if k != 1 && kind == kindSlide && hasUnscalableGraphics(data) {
		m.st.warn("%s ← %s: есть SmartArt или диаграмма, её внутренняя разметка не пересчитывается — проверьте глазами",
			shortPart(newName), path.Base(src.src))
	}
	m.out.set(newName, data)
	m.outCT.setOverride(newName, srcCT.typeOf(part))

	switch kind {
	case kindLayout:
		m.st.layouts++
	case kindMaster:
		m.st.masters++
	case kindNotesSlide:
		m.st.notes++
	}

	rels, err := readRels(src, part)
	if err != nil {
		return "", err
	}
	var newRels []rel
	for _, r := range rels {
		if r.external() || r.Type == relHyperlink {
			newRels = append(newRels, r)
			continue
		}
		target := resolveTarget(part, r.Target)
		dst, err := m.copyPart(src, srcCT, target, k)
		if err != nil {
			return "", err
		}
		newRels = append(newRels, rel{ID: r.ID, Type: r.Type, Target: relTarget(newName, dst)})
	}
	if len(newRels) > 0 {
		writeRels(m.out, newName, newRels)
	}

	// Образцы слайдов и мастер заметок обязаны быть перечислены в presentation.xml.
	if kind == kindMaster {
		if err := registerInPresentation(m.out, newName, relSlideMaster, "sldMasterIdLst", "sldMasterId"); err != nil {
			return "", err
		}
	}
	if kind == kindNotesMaster && m.notesMaster == "" {
		if err := registerInPresentation(m.out, newName, relNotesMaster, "notesMasterIdLst", "notesMasterId"); err != nil {
			return "", err
		}
		m.notesMaster = newName
	}
	return newName, nil
}

// Медиа переносим с дедупликацией по содержимому: одна и та же картинка
// в нескольких презентациях ляжет в результат один раз.
func (m *merger) copyMedia(src *pkg, srcCT *contentTypes, part string) (string, error) {
	key := src.src + "\x00" + part
	data := src.parts[part]
	sum := sha256.Sum256(data)
	hash := hex.EncodeToString(sum[:])

	if name, ok := m.mediaByHash[hash]; ok {
		m.copied[key] = name
		m.st.mediaReused++
		return name, nil
	}
	dir, base, ext := splitPartName(part)
	newName := m.out.freeName(dir, base, ext)
	m.out.set(newName, data)
	m.copied[key] = newName
	m.mediaByHash[hash] = newName
	m.outCT.ensureDefault(ext, srcCT.typeOf(part))
	m.st.mediaCopied++
	return newName, nil
}

// Учесть медиа, которые уже лежат в первом файле, чтобы не копировать их повторно.
func (m *merger) indexBaseMedia() {
	for _, part := range m.out.order {
		if kindOf(m.outCT, part) != kindMedia {
			continue
		}
		sum := sha256.Sum256(m.out.parts[part])
		hash := hex.EncodeToString(sum[:])
		if _, ok := m.mediaByHash[hash]; !ok {
			m.mediaByHash[hash] = part
		}
	}
}

// Завершить сборку: записать таблицу типов и проставить целевой размер слайда.
func (m *merger) finish() error {
	m.outCT.write(m.out)
	return setSlideSize(m.out, m.target)
}

var numberedRe = regexp.MustCompile(`^(.*?)(\d+)?$`)

// Разложить путь части на каталог, имя без числового суффикса и расширение:
// "ppt/slides/slide12.xml" -> ("ppt/slides", "slide", "xml").
func splitPartName(part string) (dir, base, ext string) {
	dir = dirOf(part)
	name := path.Base(part)
	ext = strings.TrimPrefix(path.Ext(name), ".")
	stem := strings.TrimSuffix(name, path.Ext(name))
	if m := numberedRe.FindStringSubmatch(stem); m != nil && m[1] != "" {
		return dir, m[1], ext
	}
	return dir, stem, ext
}

// Роль части в презентации.
//
// Определяется по типу содержимого, а не по пути: путь части ничем не регламентирован.
// В реальных файлах встречается, например, ppt/slideMasters/theme/theme2.xml —
// тема, лежащая внутри каталога образцов.
type partKind int

const (
	kindOther partKind = iota
	kindSlide
	kindLayout
	kindMaster
	kindNotesSlide
	kindNotesMaster
	kindMedia
)

func kindOf(ct *contentTypes, part string) partKind {
	switch t := ct.typeOf(part); t {
	case ctSlide:
		return kindSlide
	case ctSlideLayout:
		return kindLayout
	case ctSlideMaster:
		return kindMaster
	case ctNotesSlide:
		return kindNotesSlide
	case ctNotesMaster:
		return kindNotesMaster
	default:
		if strings.HasPrefix(t, "image/") || strings.HasPrefix(t, "video/") || strings.HasPrefix(t, "audio/") {
			return kindMedia
		}
		// Часть без внятного типа, но лежащая среди медиа, всё равно медиа:
		// так помечают вложения вроде application/octet-stream.
		if strings.HasPrefix(part, "ppt/media/") {
			return kindMedia
		}
		return kindOther
	}
}

// Части, чья геометрия привязана к размеру слайда.
// Заметки живут на своей странице (p:notesSz) и от размера слайда не зависят.
func (k partKind) scalable() bool {
	return k == kindSlide || k == kindLayout || k == kindMaster
}

// Короткое имя части для сообщений: "ppt/slides/slide7.xml" -> "slide7".
func shortPart(part string) string {
	return strings.TrimSuffix(path.Base(part), path.Ext(part))
}
