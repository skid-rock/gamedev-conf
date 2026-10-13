package main

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

// Ключевые части пакета и пространства имён OOXML.
const (
	ctypesPart       = "[Content_Types].xml"
	presentationPart = "ppt/presentation.xml"

	nsRelationships = "http://schemas.openxmlformats.org/package/2006/relationships"
	nsR             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
)

// Типы отношений, которые разбираем осознанно. Остальные копируются как есть.
const (
	relSlide       = nsR + "/slide"
	relSlideLayout = nsR + "/slideLayout"
	relSlideMaster = nsR + "/slideMaster"
	relNotesSlide  = nsR + "/notesSlide"
	relNotesMaster = nsR + "/notesMaster"
	relTheme       = nsR + "/theme"
	relImage       = nsR + "/image"
	relHyperlink   = nsR + "/hyperlink"
)

// Content-type для частей, которые создаём сами.
const (
	ctSlide       = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	ctSlideLayout = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
	ctSlideMaster = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
	ctNotesSlide  = "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml"
	ctNotesMaster = "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml"
	ctTheme       = "application/vnd.openxmlformats-officedocument.theme+xml"
)

// Одно отношение из файла .rels.
type rel struct {
	ID         string
	Type       string
	Target     string
	TargetMode string // "External" у гиперссылок и внешних медиа
}

func (r rel) external() bool { return r.TargetMode == "External" }

// Путь к .rels для заданной части: "ppt/slides/slide1.xml" -> "ppt/slides/_rels/slide1.xml.rels".
func relsPathFor(part string) string {
	return path.Join(dirOf(part), "_rels", path.Base(part)+".rels")
}

// Прочитать отношения части. Отсутствие файла .rels — не ошибка: у части просто нет связей.
func readRels(p *pkg, part string) ([]rel, error) {
	data, ok := p.parts[relsPathFor(part)]
	if !ok {
		return nil, nil
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, fmt.Errorf("%s: %w", relsPathFor(part), err)
	}
	root := doc.SelectElement("Relationships")
	if root == nil {
		return nil, fmt.Errorf("%s: нет корня Relationships", relsPathFor(part))
	}
	var out []rel
	for _, e := range root.SelectElements("Relationship") {
		out = append(out, rel{
			ID:         e.SelectAttrValue("Id", ""),
			Type:       e.SelectAttrValue("Type", ""),
			Target:     e.SelectAttrValue("Target", ""),
			TargetMode: e.SelectAttrValue("TargetMode", ""),
		})
	}
	return out, nil
}

// Записать отношения части, полностью заменив прежний .rels.
func writeRels(p *pkg, part string, rels []rel) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8" standalone="yes"`)
	root := doc.CreateElement("Relationships")
	root.CreateAttr("xmlns", nsRelationships)
	for _, r := range rels {
		e := root.CreateElement("Relationship")
		e.CreateAttr("Id", r.ID)
		e.CreateAttr("Type", r.Type)
		e.CreateAttr("Target", r.Target)
		if r.TargetMode != "" {
			e.CreateAttr("TargetMode", r.TargetMode)
		}
	}
	doc.Indent(2)
	data, err := doc.WriteToBytes()
	if err != nil {
		// etree не выдумывает ошибок на построенном нами дереве; молча терять данные нельзя.
		panic(fmt.Sprintf("сериализация %s: %v", relsPathFor(part), err))
	}
	p.set(relsPathFor(part), data)
}

// Свободный идентификатор отношения, не занятый в списке.
func freeRelID(rels []rel) string {
	max := 0
	for _, r := range rels {
		if n, err := strconv.Atoi(strings.TrimPrefix(r.ID, "rId")); err == nil && n > max {
			max = n
		}
	}
	return "rId" + strconv.Itoa(max+1)
}

// Таблица типов содержимого пакета.
type contentTypes struct {
	defaults  map[string]string // расширение в нижнем регистре -> тип
	overrides map[string]string // "/ppt/slides/slide1.xml" -> тип
}

func readContentTypes(p *pkg) (*contentTypes, error) {
	data, err := p.get(ctypesPart)
	if err != nil {
		return nil, err
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, fmt.Errorf("%s: %w", ctypesPart, err)
	}
	root := doc.SelectElement("Types")
	if root == nil {
		return nil, fmt.Errorf("%s: нет корня Types", ctypesPart)
	}
	ct := &contentTypes{defaults: map[string]string{}, overrides: map[string]string{}}
	for _, e := range root.SelectElements("Default") {
		ct.defaults[strings.ToLower(e.SelectAttrValue("Extension", ""))] = e.SelectAttrValue("ContentType", "")
	}
	for _, e := range root.SelectElements("Override") {
		ct.overrides[e.SelectAttrValue("PartName", "")] = e.SelectAttrValue("ContentType", "")
	}
	return ct, nil
}

func (ct *contentTypes) write(p *pkg) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8" standalone="yes"`)
	root := doc.CreateElement("Types")
	root.CreateAttr("xmlns", "http://schemas.openxmlformats.org/package/2006/content-types")

	exts := make([]string, 0, len(ct.defaults))
	for e := range ct.defaults {
		exts = append(exts, e)
	}
	sort.Strings(exts)
	for _, e := range exts {
		d := root.CreateElement("Default")
		d.CreateAttr("Extension", e)
		d.CreateAttr("ContentType", ct.defaults[e])
	}

	names := make([]string, 0, len(ct.overrides))
	for n := range ct.overrides {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		o := root.CreateElement("Override")
		o.CreateAttr("PartName", n)
		o.CreateAttr("ContentType", ct.overrides[n])
	}
	doc.Indent(2)
	data, err := doc.WriteToBytes()
	if err != nil {
		panic(fmt.Sprintf("сериализация %s: %v", ctypesPart, err))
	}
	p.set(ctypesPart, data)
}

// Тип части в исходном пакете: сначала точное правило, потом правило по расширению.
func (ct *contentTypes) typeOf(part string) string {
	if t, ok := ct.overrides["/"+part]; ok {
		return t
	}
	return ct.defaults[strings.ToLower(strings.TrimPrefix(path.Ext(part), "."))]
}

// Зарегистрировать точное правило для части.
func (ct *contentTypes) setOverride(part, typ string) {
	if typ != "" {
		ct.overrides["/"+part] = typ
	}
}

// Зарегистрировать правило по расширению, если его ещё нет.
func (ct *contentTypes) ensureDefault(ext, typ string) {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	if ext == "" || typ == "" {
		return
	}
	if _, ok := ct.defaults[ext]; !ok {
		ct.defaults[ext] = typ
	}
}

func (ct *contentTypes) dropOverride(part string) {
	delete(ct.overrides, "/"+part)
}
