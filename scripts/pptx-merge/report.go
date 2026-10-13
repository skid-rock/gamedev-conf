package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

// Один входной файл и всё, что о нём известно до склейки.
type source struct {
	pk     *pkg
	size   slideSize
	slides int
	name   string
}

// Обрезать имя до ширины колонки, считая руны, а не байты: имена файлов русские.
func fit(s string, w int) string {
	if utf8.RuneCountInString(s) <= w {
		return s + strings.Repeat(" ", w-utf8.RuneCountInString(s))
	}
	r := []rune(s)
	return string(r[:w-1]) + "…"
}

// План: что и с каким коэффициентом поедет в результат.
func printPlan(srcs []source, target slideSize, scale bool) {
	fmt.Printf("Целевой размер слайда: %s EMU (%s)\n\n", target, target.inches())

	const w = 44
	fmt.Printf("%-*s %7s  %-19s %s\n", w, "Файл", "Слайдов", "Размер слайда", "Пересчёт")
	fmt.Println(strings.Repeat("─", w+7+21+10))

	total := 0
	for _, s := range srcs {
		total += s.slides
		k := 1.0
		if s.size.cx > 0 {
			k = float64(target.cx) / float64(s.size.cx)
		}
		note := "не нужен"
		switch {
		case math.Abs(k-1) < 1e-9:
			note = "совпадает"
		case scale:
			note = fmt.Sprintf("×%g", round4(k))
		default:
			note = fmt.Sprintf("×%g — НЕ будет применён", round4(k))
		}
		fmt.Printf("%-*s %7d  %-19s %s\n", w, fit(s.name, w), s.slides, s.size.inches(), note)
	}
	fmt.Println(strings.Repeat("─", w+7+21+10))
	fmt.Printf("%-*s %7d\n\n", w, "Итого", total)
}

// Предупредить о том, что видно до сборки: разные пропорции, выключенный пересчёт.
func printPreflight(srcs []source, target slideSize, scale bool) {
	var mismatched int
	for _, s := range srcs {
		if s.size.cx != target.cx || s.size.cy != target.cy {
			mismatched += s.slides
		}
		if s.size.cx > 0 && math.Abs(s.size.ratio()-target.ratio()) > 0.005 {
			warnf("%s: пропорция %.3f против целевой %.3f — содержимое будет обрезано или растянуто",
				s.name, s.size.ratio(), target.ratio())
		}
	}
	if mismatched == 0 {
		return
	}
	if scale {
		warnf("геометрия %d %s будет пересчитана: координаты, кегль текста, отступы, таблицы, тени",
			mismatched, plural(mismatched, "слайда", "слайдов", "слайдов"))
		warnf("пересчёт не затрагивает внутреннюю разметку SmartArt и диаграмм, автоподбор размера текста может сработать иначе — результат нужно просмотреть")
	} else {
		warnf("слайдов другого размера: %d — пересчёт выключен, они отрисуются не в масштабе и уедут в левый верхний угол", mismatched)
		warnf("включить пересчёт: флаг -scale")
	}
}

// Итоги сборки.
func printResult(m *merger, baseSlides int, out string, dry bool) {
	fmt.Println("Результат:")
	fmt.Printf("  слайдов          %d (из первого файла %d, перенесено %d)\n", baseSlides+m.st.slides, baseSlides, m.st.slides)
	fmt.Printf("  макетов          %d перенесено\n", m.st.layouts)
	fmt.Printf("  образцов         %d перенесено\n", m.st.masters)
	fmt.Printf("  заметок          %d перенесено\n", m.st.notes)
	fmt.Printf("  медиафайлов      %d скопировано, %d переиспользовано (одинаковые не дублируются)\n", m.st.mediaCopied, m.st.mediaReused)
	if m.st.scaledAttrs > 0 {
		fmt.Printf("  пересчитано      %d размерных значений\n", m.st.scaledAttrs)
	}
	fmt.Printf("  объём            %.1f МБ\n", float64(m.out.size())/(1<<20))

	for _, w := range m.st.warnings {
		warnf("%s", w)
	}

	fmt.Println()
	if dry {
		fmt.Println("Пробный прогон: файл не записан. Уберите -n, чтобы собрать презентацию.")
		return
	}
	fmt.Printf("Записано: %s\n", out)
}

func warnf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "  ! "+format+"\n", a...)
}

func shortName(p string) string { return path.Base(p) }

// Округлить коэффициент до четырёх знаков, чтобы ×1.3333333 не пугал в отчёте.
func round4(k float64) float64 { return math.Round(k*10000) / 10000 }

// Согласование существительного с числом: 1 слайд, 2 слайда, 5 слайдов.
func plural(n int, one, few, many string) string {
	if m := n % 100; m >= 11 && m <= 14 {
		return many
	}
	switch n % 10 {
	case 1:
		return one
	case 2, 3, 4:
		return few
	}
	return many
}
