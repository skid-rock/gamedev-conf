package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Короткое имя трека для таблицы.
func shortName(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// Сверка найденного с заявленным в промпте.
func verdict(a analysis) string {
	if !a.HasDeclared {
		return "в промпте нет"
	}
	best := a.best()
	switch {
	case best.sameKey(a.Declared):
		return "совпало"
	case best.Tonic == a.Declared.Tonic:
		return "тоника та же, лад другой"
	case best.samePitchSet(a.Declared):
		return "ноты те же, тоника другая"
	default:
		return "не совпало"
	}
}

// Отрыв лучшего кандидата от второго: чем меньше, тем спорнее ответ.
func margin(a analysis) float64 {
	return a.Ranked[0].Score - a.Ranked[1].Score
}

// Таблица в стандартный вывод.
func printTable(results []analysis) {
	fmt.Printf("%-24s %-18s %-18s %6s %6s %8s %5s %6s  %s\n",
		"ТРЕК", "ЗАЯВЛЕНО", "НАЙДЕНО", "СЧЁТ", "ОТРЫВ", "ЗВУКОРЯД", "БАС", "СТРОЙ", "ОЦЕНКА")
	for _, a := range results {
		declared := "—"
		if a.HasDeclared {
			declared = a.Declared.Name()
		}
		best := a.best()
		fmt.Printf("%-24s %-18s %-18s %6.3f %6.3f %7.0f%% %5s %+5.0fц  %s\n",
			shortName(a.Path),
			declared,
			best.Name(),
			best.Score,
			margin(a),
			scaleShare(a.Total, best)*100,
			noteNames[a.Bass.top(1)[0]],
			a.Tuning*100,
			verdict(a),
		)
	}
}

// Подробный разбор одного трека.
func printDetails(a analysis, top int) {
	best := a.best()
	fmt.Printf("\n== %s ==\n", shortName(a.Path))
	fmt.Printf("Длительность %.1f с, строй %+.0f центов от A440\n", a.Duration, a.Tuning*100)

	fmt.Println("\nКандидаты:")
	for i, k := range a.Ranked[:min(top, len(a.Ranked))] {
		fmt.Printf("  %d. %-18s %.3f  звукоряд %2.0f%%\n", i+1, k.Name(), k.Score, scaleShare(a.Total, k)*100)
	}
	if a.HasDeclared {
		rank, k := rankOf(a.Ranked, a.Declared)
		fmt.Printf("  Заявлено: %s — место %d, счёт %.3f, звукоряд %2.0f%%\n",
			k.Name(), rank, k.Score, scaleShare(a.Total, k)*100)
	}

	printEvidence(a, best.Tonic)
	if a.HasDeclared && a.Declared.Tonic != best.Tonic {
		printEvidence(a, a.Declared.Tonic)
	}

	printProfile(a, best.Tonic)

	if len(a.Rows) > 0 {
		var parts []string
		for _, c := range chordSummary(a.Rows) {
			parts = append(parts, fmt.Sprintf("%s ×%d", c.Name, c.Count))
		}
		fmt.Printf("\nАккорды по тактам (%.2f BPM, первая доля %.2f с; %s):\n  %s\n",
			a.BPM, a.Start, a.GridSource, strings.Join(parts, ", "))
	}
}

// Опознавательные ступени: какой из двух соседних вариантов звучит.
func printEvidence(a analysis, tonic int) {
	fmt.Printf("\nОпознавательные ступени от %s (общий профиль):\n", noteNames[tonic])
	for _, p := range degreeEvidence(a.Total, tonic) {
		fmt.Printf("  %-3s %-3s %5.1f%%   %-3s %-3s %5.1f%%   %s\n",
			degreeNames[p.Low], noteNames[(tonic+p.Low)%12], p.LowShare*100,
			degreeNames[p.High], noteNames[(tonic+p.High)%12], p.HighShare*100,
			leaning(p))
	}
}

// Какая из двух ступеней перевешивает и насколько.
func leaning(p degreePair) string {
	lo, hi := p.LowShare, p.HighShare
	switch {
	case lo+hi < 0.02:
		return "обеих почти нет"
	case lo > hi*3:
		return "→ " + degreeNames[p.Low]
	case hi > lo*3:
		return "→ " + degreeNames[p.High]
	default:
		return "обе звучат"
	}
}

// Профиль нот от тоники: бас и середина рядом.
func printProfile(a analysis, tonic int) {
	fmt.Printf("\nПрофиль нот от %s:\n", noteNames[tonic])
	fmt.Printf("  %-4s %-3s %-26s %s\n", "ст.", "нота", "бас", "середина")
	maxBass, maxMid := maxOf(a.Bass), maxOf(a.Mid)
	for step := 0; step < 12; step++ {
		pc := (tonic + step) % 12
		fmt.Printf("  %-4s %-3s %5.1f%% %-19s %5.1f%% %s\n",
			degreeNames[step], noteNames[pc],
			a.Bass[pc]*100, bar(a.Bass[pc], maxBass, 18),
			a.Mid[pc]*100, bar(a.Mid[pc], maxMid, 18))
	}
}

// Таблица по тактам.
func printBars(a analysis, barsPerRow int) {
	fmt.Printf("\n== %s: по %d такт(а) ==\n", shortName(a.Path), barsPerRow)
	if a.BPM <= 0 {
		fmt.Println("  Темп неизвестен: запустите bpm-detect --sidecar или передайте --bpm.")
		return
	}
	fmt.Printf("  %.2f BPM, первая доля %.2f с (%s)\n", a.BPM, a.Start, a.GridSource)

	rows := a.Rows
	if barsPerRow != 1 {
		// Строки по одному такту уже посчитаны; для других размеров режем заново
		// не кадры, а готовые строки: такт — наименьшая единица таблицы.
		rows = mergeRows(a.Rows, barsPerRow)
	}

	tonic := a.best().Tonic
	fmt.Printf("  %-5s %6s  %-6s %-22s %-34s\n", "ТАКТ", "ВРЕМЯ", "АККОРД", "БАС", "СЕРЕДИНА")
	for _, r := range rows {
		fmt.Printf("  %-5d %6.1f  %-6s %-22s %-34s\n",
			r.Bar, r.Start, r.Chord.String(), topNotes(r.Bass, 3, tonic), topNotes(r.Mid, 5, tonic))
	}
}

// Склеить потактовые строки в строки по n тактов.
func mergeRows(rows []barRow, n int) []barRow {
	var out []barRow
	for i := 0; i < len(rows); i += n {
		merged := barRow{Bar: rows[i].Bar, Start: rows[i].Start}
		for _, r := range rows[i:min(i+n, len(rows))] {
			merged.Bass.add(r.Bass, 1)
			merged.Mid.add(r.Mid, 1)
		}
		merged.Chord = guessChord(merged.Bass, merged.Mid)
		out = append(out, merged)
	}
	return out
}

// Самые громкие ноты куска: «A 41 · E 22 · G 12», тихий кусок — прочерк.
func topNotes(c chroma, n, tonic int) string {
	if c.sum() == 0 {
		return "—"
	}
	norm := c.normalized()
	var parts []string
	for _, pc := range norm.top(n) {
		if norm[pc] < 0.03 {
			break
		}
		parts = append(parts, fmt.Sprintf("%s %.0f", noteNames[pc], norm[pc]*100))
	}
	return strings.Join(parts, " · ")
}

func maxOf(c chroma) float64 {
	var m float64
	for _, v := range c {
		m = max(m, v)
	}
	return m
}

func bar(v, maxV float64, width int) string {
	if maxV <= 0 {
		return ""
	}
	return strings.Repeat("█", int(v/maxV*float64(width)+0.5))
}

// Результат в json — для скриптов и сравнения прогонов.
func writeJSON(path string, results []analysis, top int) error {
	type candidate struct {
		Key        string  `json:"key"`
		Score      float64 `json:"score"`
		ScaleShare float64 `json:"scale_share"`
	}
	type record struct {
		File         string             `json:"file"`
		Declared     string             `json:"declared_key,omitempty"`
		DeclaredRank int                `json:"declared_rank,omitempty"`
		Key          string             `json:"key"`
		Score        float64            `json:"key_score"`
		Margin       float64            `json:"key_margin"`
		ScaleShare   float64            `json:"scale_share"`
		BassRoot     string             `json:"bass_root"`
		TuningCents  float64            `json:"tuning_cents"`
		Verdict      string             `json:"verdict"`
		Candidates   []candidate        `json:"candidates"`
		Bass         map[string]float64 `json:"bass_chroma"`
		Mid          map[string]float64 `json:"mid_chroma"`
		Chords       []string           `json:"chords_by_bar,omitempty"`
	}

	toMap := func(c chroma) map[string]float64 {
		m := make(map[string]float64, 12)
		for i, v := range c {
			m[noteNames[i]] = round3(v)
		}
		return m
	}

	records := make([]record, 0, len(results))
	for _, a := range results {
		best := a.best()
		r := record{
			File:        a.Path,
			Key:         best.Name(),
			Score:       round3(best.Score),
			Margin:      round3(margin(a)),
			ScaleShare:  round3(scaleShare(a.Total, best)),
			BassRoot:    noteNames[a.Bass.top(1)[0]],
			TuningCents: float64(int(a.Tuning*1000)) / 10,
			Verdict:     verdict(a),
			Bass:        toMap(a.Bass),
			Mid:         toMap(a.Mid),
		}
		if a.HasDeclared {
			rank, _ := rankOf(a.Ranked, a.Declared)
			r.Declared = a.Declared.Name()
			r.DeclaredRank = rank
		}
		for _, k := range a.Ranked[:min(top, len(a.Ranked))] {
			r.Candidates = append(r.Candidates, candidate{k.Name(), round3(k.Score), round3(scaleShare(a.Total, k))})
		}
		for _, row := range a.Rows {
			r.Chords = append(r.Chords, row.Chord.String())
		}
		records = append(records, r)
	}

	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func round3(v float64) float64 {
	return float64(int(v*1000+0.5*sign(v))) / 1000
}

func sign(v float64) float64 {
	if v < 0 {
		return -1
	}
	return 1
}
