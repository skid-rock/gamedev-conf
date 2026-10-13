package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Пакет OOXML, целиком прочитанный в память.
// pptx — это обычный zip, где каждая часть адресуется путём вида "ppt/slides/slide1.xml".
// Держим всё в памяти: презентации конференции — десятки мегабайт, это дешевле,
// чем ходить по архиву туда-сюда при разрешении связей.
type pkg struct {
	src   string            // откуда прочитан, только для сообщений об ошибках
	parts map[string][]byte // путь части -> содержимое
	order []string          // порядок частей на входе, чтобы запись была предсказуемой
}

// Прочитать pptx с диска.
func readPkg(p string) (*pkg, error) {
	zr, err := zip.OpenReader(p)
	if err != nil {
		return nil, fmt.Errorf("не открывается как zip: %w", err)
	}
	defer zr.Close()

	pk := &pkg{src: p, parts: map[string][]byte{}}
	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f.Name, err)
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f.Name, err)
		}
		name := path.Clean(f.Name)
		if _, dup := pk.parts[name]; !dup {
			pk.order = append(pk.order, name)
		}
		pk.parts[name] = data
	}
	if _, ok := pk.parts[ctypesPart]; !ok {
		return nil, fmt.Errorf("нет %s — это не pptx", ctypesPart)
	}
	if _, ok := pk.parts[presentationPart]; !ok {
		return nil, fmt.Errorf("нет %s — это не презентация", presentationPart)
	}
	return pk, nil
}

// Глубокая копия пакета: результат склейки строится поверх первого файла,
// а исходник при этом не должен меняться.
func (p *pkg) clone() *pkg {
	c := &pkg{src: p.src, parts: make(map[string][]byte, len(p.parts)), order: append([]string(nil), p.order...)}
	for k, v := range p.parts {
		c.parts[k] = append([]byte(nil), v...)
	}
	return c
}

func (p *pkg) has(name string) bool {
	_, ok := p.parts[name]
	return ok
}

func (p *pkg) get(name string) ([]byte, error) {
	d, ok := p.parts[name]
	if !ok {
		return nil, fmt.Errorf("нет части %s", name)
	}
	return d, nil
}

// Добавить или заменить часть.
func (p *pkg) set(name string, data []byte) {
	if _, ok := p.parts[name]; !ok {
		p.order = append(p.order, name)
	}
	p.parts[name] = data
}

func (p *pkg) del(name string) {
	if _, ok := p.parts[name]; !ok {
		return
	}
	delete(p.parts, name)
	for i, n := range p.order {
		if n == name {
			p.order = append(p.order[:i], p.order[i+1:]...)
			break
		}
	}
}

// Части, лежащие в каталоге dir и подходящие под шаблон "<base><число>.<ext>".
func (p *pkg) numbered(dir, base, ext string) []string {
	re := regexp.MustCompile(`^` + regexp.QuoteMeta(dir+"/"+base) + `\d+\.` + regexp.QuoteMeta(ext) + `$`)
	var out []string
	for name := range p.parts {
		if re.MatchString(name) {
			out = append(out, name)
		}
	}
	sort.Slice(out, func(i, j int) bool { return partIndex(out[i]) < partIndex(out[j]) })
	return out
}

var indexRe = regexp.MustCompile(`(\d+)\.[^.]+$`)

// Числовой суффикс в имени части: "slide12.xml" -> 12. Для сортировки.
func partIndex(name string) int {
	m := indexRe.FindStringSubmatch(name)
	if m == nil {
		return 0
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

// Свободное имя части вида "<dir>/<base><N>.<ext>", начиная с N=1.
func (p *pkg) freeName(dir, base, ext string) string {
	for n := 1; ; n++ {
		name := fmt.Sprintf("%s/%s%d.%s", dir, base, n, ext)
		if !p.has(name) {
			return name
		}
	}
}

// Записать пакет на диск.
// [Content_Types].xml кладём первым: по спецификации порядок не важен,
// но часть инструментов (и старые сборки PowerPoint) ожидают его в начале архива.
func (p *pkg) write(dst string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	names := make([]string, 0, len(p.parts))
	names = append(names, ctypesPart)
	for _, n := range p.order {
		if n != ctypesPart && p.has(n) {
			names = append(names, n)
		}
	}
	// Части, добавленные в обход order (подстраховка).
	seen := map[string]bool{}
	for _, n := range names {
		seen[n] = true
	}
	var extra []string
	for n := range p.parts {
		if !seen[n] {
			extra = append(extra, n)
		}
	}
	sort.Strings(extra)
	names = append(names, extra...)

	for _, name := range names {
		data := p.parts[name]
		w, err := zw.CreateHeader(&zip.FileHeader{Name: name, Method: zip.Deflate})
		if err != nil {
			zw.Close()
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(data)); err != nil {
			zw.Close()
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}

// Суммарный объём частей — для отчёта.
func (p *pkg) size() int64 {
	var n int64
	for _, d := range p.parts {
		n += int64(len(d))
	}
	return n
}

// Каталог части: "ppt/slides/slide1.xml" -> "ppt/slides".
func dirOf(name string) string { return path.Dir(name) }

// Разрешить относительный Target отношения в абсолютный путь части пакета.
// Target в .rels задаётся относительно каталога владельца отношения.
func resolveTarget(owner, target string) string {
	if strings.HasPrefix(target, "/") {
		return path.Clean(strings.TrimPrefix(target, "/"))
	}
	return path.Clean(path.Join(dirOf(owner), target))
}

// Обратная операция: путь части -> Target относительно владельца.
func relTarget(owner, target string) string {
	ownerDir := dirOf(owner)
	if ownerDir == "." {
		return target
	}
	ownerSegs := strings.Split(ownerDir, "/")
	targetSegs := strings.Split(target, "/")
	i := 0
	for i < len(ownerSegs) && i < len(targetSegs)-1 && ownerSegs[i] == targetSegs[i] {
		i++
	}
	var b strings.Builder
	for range ownerSegs[i:] {
		b.WriteString("../")
	}
	b.WriteString(strings.Join(targetSegs[i:], "/"))
	return b.String()
}
