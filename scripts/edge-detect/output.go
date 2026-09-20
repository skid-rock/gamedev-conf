package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
)

// Короткое имя трека для таблицы.
func shortName(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// Короткий итог по отрезку: в скольких тактах есть саб, хэты и ноты в середине.
func spanVerdict(s span) string {
	return fmt.Sprintf("саб %d/%d · хэты %d/%d · ноты %d/%d",
		s.Present[bandSub], s.N, s.Present[bandHigh], s.N, s.Tonal, s.N)
}

// Остановился ли трек в конце.
func stopVerdict(e edges) string {
	rel := e.lastRel()
	var text string
	switch {
	case rel <= stopQuietDB:
		text = "стоп"
	case rel >= stopLoudDB:
		text = "звук до конца"
	default:
		text = "затухание"
	}
	if e.HasEnd && e.End.FlatRatio < tonalRatio/2 && e.End.Rel[bandMid] > soundDB {
		text += ", хвост с нотами"
	}
	return text
}

// Цифры отрезка: уровни полос и ровность.
func spanNumbers(s span) string {
	ratio := "—"
	if !math.IsNaN(s.FlatRatio) {
		ratio = fmt.Sprintf("%.2f", s.FlatRatio)
	}
	return fmt.Sprintf("саб %+4.0f  низ-сер %+4.0f  сер %+4.0f  верх %+4.0f дБ, ровность %s",
		s.Rel[bandSub], s.Rel[bandLowMid], s.Rel[bandMid], s.Rel[bandHigh], ratio)
}

func entryText(bar, planned int) string {
	if bar == 0 {
		return "не вошёл"
	}
	if bar == planned {
		return fmt.Sprintf("такт %d", bar)
	}
	return fmt.Sprintf("такт %d (план %d)", bar, planned)
}

// Разбор одного трека блоком.
func printTrack(r result) {
	e := r.Edges
	total := "—"
	if r.Plan.Total > 0 {
		total = fmt.Sprint(r.Plan.Total)
	}
	fmt.Printf("== %s ==\n", shortName(r.Path))
	fmt.Printf("  %.2f BPM, первая доля %.3f с (%s); тактов %d, по плану %s; длины краёв — %s\n",
		r.BPM, r.Start, r.GridSource, e.LastBar, total, r.Plan.Source)
	fmt.Printf("  Интро  %3d–%-3d  %s\n", e.Intro.From, e.Intro.To, spanNumbers(e.Intro))
	fmt.Printf("                  → %s\n", spanVerdict(e.Intro))
	fmt.Printf("  Входят          саб: %s, хэты: %s\n", entryText(e.SubEntry, r.Plan.Intro+1), entryText(e.HighEntry, r.Plan.Intro+1))
	if e.HasOutro {
		fmt.Printf("  Аутро  %3d–%-3d  %s\n", e.Outro.From, e.Outro.To, spanNumbers(e.Outro))
		fmt.Printf("                  → %s\n", spanVerdict(e.Outro))
	}
	if e.HasEnd {
		fmt.Printf("  Конец  %3d–%-3d  %s\n", e.End.From, e.End.To, spanNumbers(e.End))
	}
	fmt.Printf("                  последний такт %+.0f дБ от тела → %s\n\n", e.lastRel(), stopVerdict(e))
}

// Сводная таблица по всем трекам.
func printSummary(results []result) {
	fmt.Printf("%-22s %7s %6s  %-30s %-6s %-30s %s\n", "ТРЕК", "BPM", "ТАКТОВ", "ИНТРО", "САБ С", "АУТРО", "КОНЕЦ")
	for _, r := range results {
		e := r.Edges
		outro := "—"
		if e.HasOutro {
			outro = spanVerdict(e.Outro)
		}
		sub := "—"
		if e.SubEntry > 0 {
			sub = fmt.Sprint(e.SubEntry)
		}
		fmt.Printf("%-22s %7.2f %6d  %-30s %-6s %-30s %s\n",
			shortName(r.Path), r.BPM, e.LastBar, spanVerdict(e.Intro), sub, outro, stopVerdict(e))
	}
}

// Потактовая таблица краёв: начало и конец трека.
func printBars(r result, around int) {
	e := r.Edges
	fmt.Printf("== %s: такты, дБ от тела ==\n", shortName(r.Path))
	fmt.Printf("  %5s %7s %6s %8s %6s %6s %6s  %s\n", "ТАКТ", "ВРЕМЯ", "САБ", "НИЗ-СЕР", "СЕР", "ВЕРХ", "ВСЁ", "РОВНОСТЬ")
	show := func(from, to int) {
		for i := max(0, from); i <= min(e.LastBar-1, to); i++ {
			b := e.Bars[i]
			flat := "—"
			if !math.IsNaN(b.Flat) && e.Flat > 0 {
				flat = fmt.Sprintf("%.2f", b.Flat/e.Flat)
				if b.tonal(e) {
					flat += " ноты"
				}
			}
			fmt.Printf("  %5d %7.1f %+6.0f %+8.0f %+6.0f %+6.0f %+6.0f  %s\n", i+1, b.Start,
				b.Level[bandSub]-e.Body[bandSub], b.Level[bandLowMid]-e.Body[bandLowMid],
				b.Level[bandMid]-e.Body[bandMid], b.Level[bandHigh]-e.Body[bandHigh],
				b.Level[bandAll]-e.Body[bandAll], flat)
		}
	}
	show(0, r.Plan.Intro+around-1)
	fmt.Println("  …")
	show(e.LastBar-r.Plan.End-r.Plan.Outro-around, e.LastBar-1)
	fmt.Println()
}
