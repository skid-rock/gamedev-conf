// pptx-merge склеивает несколько презентаций PowerPoint в одну.
//
// Слайды переносятся вместе со своим оформлением: макетом, образцом, темой,
// картинками и заметками докладчика. Единый шаблон к ним не применяется —
// каждый доклад остаётся таким, каким его сделал автор.
//
// Размер слайда в pptx один на всю колоду. Если источники свёрстаны под разные
// размеры, флаг -scale пересчитывает геометрию под целевой; без него слайды
// переносятся как есть и отрисуются не в масштабе.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	out := flag.String("o", "", "файл результата")
	dry := flag.Bool("n", false, "пробный прогон: показать план и предупреждения, ничего не записывать")
	scale := flag.Bool("scale", false, "пересчитать геометрию слайдов под целевой размер")
	sizeFlag := flag.String("size", "12192000x6858000", "целевой размер слайда в EMU (ШИРИНАxВЫСОТА) либо first — взять у первого файла")
	force := flag.Bool("f", false, "перезаписать файл результата, если он уже существует")
	flag.Usage = usage
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		usage()
		os.Exit(2)
	}
	if !*dry && *out == "" {
		fail("укажите файл результата флагом -o (или -n для пробного прогона)")
	}
	if !*dry && !*force {
		if _, err := os.Stat(*out); err == nil {
			fail("%s уже существует — укажите другое имя или добавьте -f для перезаписи", *out)
		}
	}

	srcs, err := readSources(files)
	if err != nil {
		fail("%v", err)
	}

	target, err := parseTargetSize(*sizeFlag, srcs[0].size)
	if err != nil {
		fail("%v", err)
	}

	printPlan(srcs, target, *scale)
	printPreflight(srcs, target, *scale)

	m, err := newMerger(srcs[0].pk, target, *scale)
	if err != nil {
		fail("%s: %v", srcs[0].name, err)
	}
	m.indexBaseMedia()
	if err := m.rescaleBase(srcs[0].size); err != nil {
		fail("%s: %v", srcs[0].name, err)
	}
	for _, s := range srcs[1:] {
		fmt.Fprintf(os.Stderr, "переношу %s (%d %s)\n", s.name, s.slides, plural(s.slides, "слайд", "слайда", "слайдов"))
		if err := m.mergeFile(s.pk, s.size); err != nil {
			fail("%s: %v", s.name, err)
		}
	}
	if err := m.finish(); err != nil {
		fail("сборка: %v", err)
	}
	fmt.Fprintln(os.Stderr)

	if !*dry {
		if err := m.out.write(*out); err != nil {
			fail("запись %s: %v", *out, err)
		}
	}
	printResult(m, srcs[0].slides, *out, *dry)
}

// Прочитать все входные файлы и узнать размер слайда каждого.
func readSources(files []string) ([]source, error) {
	var srcs []source
	for _, f := range files {
		pk, err := readPkg(f)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", shortName(f), err)
		}
		size, err := readSlideSize(pk)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", shortName(f), err)
		}
		slides, err := slideOrder(pk)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", shortName(f), err)
		}
		srcs = append(srcs, source{pk: pk, size: size, slides: len(slides), name: shortName(f)})
	}
	return srcs, nil
}

// Целевой размер: явные числа или размер первого файла.
func parseTargetSize(flagValue string, first slideSize) (slideSize, error) {
	if flagValue == "first" {
		return first, nil
	}
	parts := strings.SplitN(strings.ToLower(flagValue), "x", 2)
	if len(parts) != 2 {
		return slideSize{}, fmt.Errorf("-size: ожидается ШИРИНАxВЫСОТА в EMU либо first, получено %q", flagValue)
	}
	cx, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	cy, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
	if err1 != nil || err2 != nil || cx <= 0 || cy <= 0 {
		return slideSize{}, fmt.Errorf("-size: не разбирается %q", flagValue)
	}
	return slideSize{cx, cy}, nil
}

func usage() {
	fmt.Fprint(os.Stderr, `pptx-merge — склеить презентации PowerPoint в одну.

Использование:
  pptx-merge -o итог.pptx первая.pptx вторая.pptx ...
  pptx-merge -n первая.pptx вторая.pptx        пробный прогон без записи

Порядок слайдов в результате повторяет порядок файлов в командной строке.
Оформление каждого доклада сохраняется: свои макеты, образцы и темы.

Флаги:
`)
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, `
Размеры слайда (EMU, 914400 на дюйм):
  12192000x6858000   16:9, стандартный в PowerPoint   13.33×7.5″
  9144000x5143500    16:9, прежний размер             10×5.63″
  18288000x10287000  16:9, удвоенный                  20×11.25″
`)
}

func fail(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "ошибка: "+format+"\n", a...)
	os.Exit(1)
}
