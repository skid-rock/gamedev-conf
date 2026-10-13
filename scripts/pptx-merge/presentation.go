package main

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

// Размер слайда в EMU (914400 единиц на дюйм).
type slideSize struct{ cx, cy int64 }

func (s slideSize) String() string { return fmt.Sprintf("%dx%d", s.cx, s.cy) }

// В дюймах — для читаемого отчёта.
func (s slideSize) inches() string {
	return fmt.Sprintf("%.2f×%.2f″", float64(s.cx)/914400, float64(s.cy)/914400)
}

func (s slideSize) ratio() float64 {
	if s.cy == 0 {
		return 0
	}
	return float64(s.cx) / float64(s.cy)
}

// Разобрать presentation.xml в дерево.
func presentationDoc(p *pkg) (*etree.Document, *etree.Element, error) {
	data, err := p.get(presentationPart)
	if err != nil {
		return nil, nil, err
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", presentationPart, err)
	}
	root := doc.Root()
	if root == nil {
		return nil, nil, fmt.Errorf("%s: пустой документ", presentationPart)
	}
	return doc, root, nil
}

func savePresentation(p *pkg, doc *etree.Document) error {
	data, err := doc.WriteToBytes()
	if err != nil {
		return err
	}
	p.set(presentationPart, data)
	return nil
}

// Размер слайда презентации.
func readSlideSize(p *pkg) (slideSize, error) {
	_, root, err := presentationDoc(p)
	if err != nil {
		return slideSize{}, err
	}
	e := root.SelectElement("sldSz")
	if e == nil {
		return slideSize{}, fmt.Errorf("%s: нет p:sldSz", presentationPart)
	}
	cx, err1 := strconv.ParseInt(e.SelectAttrValue("cx", ""), 10, 64)
	cy, err2 := strconv.ParseInt(e.SelectAttrValue("cy", ""), 10, 64)
	if err1 != nil || err2 != nil || cx <= 0 || cy <= 0 {
		return slideSize{}, fmt.Errorf("%s: не читается p:sldSz", presentationPart)
	}
	return slideSize{cx, cy}, nil
}

// Проставить размер слайда. Атрибут type убираем: он описывает предустановку
// вроде screen16x9 и после пересчёта размеров может ей противоречить.
func setSlideSize(p *pkg, s slideSize) error {
	doc, root, err := presentationDoc(p)
	if err != nil {
		return err
	}
	e := root.SelectElement("sldSz")
	if e == nil {
		return fmt.Errorf("%s: нет p:sldSz", presentationPart)
	}
	e.CreateAttr("cx", strconv.FormatInt(s.cx, 10))
	e.CreateAttr("cy", strconv.FormatInt(s.cy, 10))
	e.RemoveAttr("type")
	return savePresentation(p, doc)
}

// Пути слайдов в порядке показа, а не в порядке номеров файлов:
// презентация задаёт порядок списком p:sldIdLst.
func slideOrder(p *pkg) ([]string, error) {
	_, root, err := presentationDoc(p)
	if err != nil {
		return nil, err
	}
	rels, err := readRels(p, presentationPart)
	if err != nil {
		return nil, err
	}
	byID := map[string]rel{}
	for _, r := range rels {
		byID[r.ID] = r
	}

	lst := root.SelectElement("sldIdLst")
	if lst == nil {
		return nil, nil
	}
	var out []string
	for _, e := range lst.SelectElements("sldId") {
		id := e.SelectAttrValue("r:id", "")
		r, ok := byID[id]
		if !ok || r.external() {
			continue
		}
		out = append(out, resolveTarget(presentationPart, r.Target))
	}
	return out, nil
}

// Свободный идентификатор для sldId/sldMasterId.
// PowerPoint берёт их из общего пула и требует значения не меньше 256.
func freeElementID(root *etree.Element) int {
	max := 255
	for _, listName := range []string{"sldMasterIdLst", "notesMasterIdLst", "handoutMasterIdLst", "sldIdLst"} {
		lst := root.SelectElement(listName)
		if lst == nil {
			continue
		}
		for _, e := range lst.ChildElements() {
			if n, err := strconv.Atoi(e.SelectAttrValue("id", "")); err == nil && n > max {
				max = n
			}
		}
	}
	return max + 1
}

// Порядок дочерних элементов в CT_Presentation фиксирован схемой,
// поэтому новый список вставляем на положенное ему место.
var presentationChildOrder = []string{
	"sldMasterIdLst", "notesMasterIdLst", "handoutMasterIdLst", "sldIdLst",
	"sldSz", "notesSz", "smartTags", "embeddedFontLst", "custShowLst",
	"photoAlbum", "custDataLst", "kinsoku", "defaultTextStyle", "modifyVerifier", "extLst",
}

// Найти список, создав его при необходимости в правильной позиции.
func ensureList(root *etree.Element, name string) *etree.Element {
	if e := root.SelectElement(name); e != nil {
		return e
	}
	pos := -1
	for i, n := range presentationChildOrder {
		if n == name {
			pos = i
			break
		}
	}
	e := etree.NewElement("p:" + name)
	if pos < 0 {
		root.AddChild(e)
		return e
	}
	// Встать перед первым из элементов, которые по схеме идут следом.
	later := map[string]bool{}
	for _, n := range presentationChildOrder[pos+1:] {
		later[n] = true
	}
	for i, c := range root.ChildElements() {
		if later[c.Tag] {
			root.InsertChildAt(i, e)
			return e
		}
	}
	root.AddChild(e)
	return e
}

// Зарегистрировать часть в presentation.xml: добавить отношение и запись в список.
// listName — sldIdLst / sldMasterIdLst / notesMasterIdLst, itemName — sldId / sldMasterId / notesMasterId.
func registerInPresentation(p *pkg, part, relType, listName, itemName string) error {
	doc, root, err := presentationDoc(p)
	if err != nil {
		return err
	}
	rels, err := readRels(p, presentationPart)
	if err != nil {
		return err
	}
	id := freeRelID(rels)
	rels = append(rels, rel{ID: id, Type: relType, Target: relTarget(presentationPart, part)})
	writeRels(p, presentationPart, rels)

	lst := ensureList(root, listName)
	e := lst.CreateElement("p:" + itemName)
	if itemName != "notesMasterId" {
		// У notesMasterId нет атрибута id — только ссылка на часть.
		e.CreateAttr("id", strconv.Itoa(freeElementID(root)))
	}
	e.CreateAttr("r:id", id)
	return savePresentation(p, doc)
}

// Путь части мастера заметок, если он зарегистрирован в презентации.
func notesMasterOf(p *pkg) (string, error) {
	_, root, err := presentationDoc(p)
	if err != nil {
		return "", err
	}
	lst := root.SelectElement("notesMasterIdLst")
	if lst == nil {
		return "", nil
	}
	e := lst.SelectElement("notesMasterId")
	if e == nil {
		return "", nil
	}
	rels, err := readRels(p, presentationPart)
	if err != nil {
		return "", err
	}
	want := e.SelectAttrValue("r:id", "")
	for _, r := range rels {
		if r.ID == want && !r.external() {
			return resolveTarget(presentationPart, r.Target), nil
		}
	}
	return "", nil
}
