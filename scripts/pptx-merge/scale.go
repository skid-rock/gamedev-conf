package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/beevik/etree"
)

// Атрибуты, которые пересчитываются при смене размера слайда.
// Ключ — локальное имя элемента (без префикса пространства имён).
//
// Сюда попадают только абсолютные величины: EMU (914400 на дюйм) и сотые доли пункта
// у кегля и межстрочных интервалов. Всё, что задано в процентах или в собственных
// координатах фигуры, не трогаем — оно и так пересчитывается от размера родителя:
//   - a:srcRect, a:fillRect, a:tileRect — тысячные процента;
//   - a:prstGeom/a:avLst/a:gd, a:custGeom/a:path — координаты внутри фигуры;
//   - a:spcPct, a:lnSpc в процентах, a:lumMod и прочие модификаторы цвета.
var scaledAttrs = map[string][]string{
	"off":        {"x", "y"},
	"chOff":      {"x", "y"},
	"ext":        {"cx", "cy"},
	"chExt":      {"cx", "cy"},
	"ln":         {"w"},
	"rPr":        {"sz", "spc"},
	"defRPr":     {"sz", "spc"},
	"endParaRPr": {"sz", "spc"},
	"pPr":        {"marL", "marR", "indent", "defTabSz"},
	"bodyPr":     {"lIns", "tIns", "rIns", "bIns"},
	"spcPts":     {"val"},
	"gridCol":    {"w"},
	"tr":         {"h"},
	"tcPr":       {"marL", "marR", "marT", "marB"},
	"outerShdw":  {"blurRad", "dist"},
	"innerShdw":  {"blurRad", "dist"},
	"reflection": {"blurRad", "dist"},
	"glow":       {"rad"},
	"softEdge":   {"rad"},
	"tab":        {"pos"},
}

func init() {
	// Уровни списка в lstStyle несут те же отступы, что и a:pPr.
	for i := 1; i <= 9; i++ {
		scaledAttrs[fmt.Sprintf("lvl%dpPr", i)] = scaledAttrs["pPr"]
	}
}

// Пересчитать геометрию XML-части под коэффициент k.
// Возвращает новое содержимое и число поправленных атрибутов.
func scaleXML(data []byte, k float64) ([]byte, int, error) {
	if k == 1 {
		return data, 0, nil
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, 0, err
	}
	n := 0
	for _, e := range doc.ChildElements() {
		n += scaleElement(e, k)
	}
	out, err := doc.WriteToBytes()
	if err != nil {
		return nil, 0, err
	}
	return out, n, nil
}

// Обойти элемент и его потомков, умножив известные абсолютные величины на k.
func scaleElement(e *etree.Element, k float64) int {
	n := 0
	if attrs, ok := scaledAttrs[e.Tag]; ok {
		for _, name := range attrs {
			a := e.SelectAttr(name)
			if a == nil {
				continue
			}
			v, err := strconv.ParseInt(a.Value, 10, 64)
			if err != nil {
				// Величина задана не числом (например, "auto") — оставляем как есть.
				continue
			}
			a.Value = strconv.FormatInt(int64(math.Round(float64(v)*k)), 10)
			n++
		}
	}
	for _, c := range e.ChildElements() {
		n += scaleElement(c, k)
	}
	return n
}

// Есть ли в части SmartArt или диаграмма: их внутренняя геометрия живёт
// в отдельных частях и нашим обходом не покрывается — о таких слайдах предупреждаем.
func hasUnscalableGraphics(data []byte) bool {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return false
	}
	found := false
	var walk func(*etree.Element)
	walk = func(e *etree.Element) {
		if found {
			return
		}
		if e.Tag == "graphicData" {
			uri := e.SelectAttrValue("uri", "")
			if uri == "http://schemas.openxmlformats.org/drawingml/2006/diagram" ||
				uri == "http://schemas.openxmlformats.org/drawingml/2006/chart" {
				found = true
				return
			}
		}
		for _, c := range e.ChildElements() {
			walk(c)
		}
	}
	for _, e := range doc.ChildElements() {
		walk(e)
	}
	return found
}
