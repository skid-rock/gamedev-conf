package main

import (
	"sort"
)

// Догадка об аккорде: мажорное или минорное трезвучие.
type chordGuess struct {
	Root  int
	Minor bool
	// Доля энергии середины внутри трезвучия, 0..1.
	Fit   float64
	Empty bool
}

func (g chordGuess) String() string {
	if g.Empty {
		return "—"
	}
	name := noteNames[g.Root]
	if g.Minor {
		name += "m"
	}
	return name
}

// Подобрать трезвучие по середине, с доводом от баса: нота баса на корне
// отличает Am от C, у которых две общие ноты.
func guessChord(bass, mid chroma) chordGuess {
	b := bass.normalized()
	m := mid.normalized()
	if m.sum() == 0 && b.sum() == 0 {
		return chordGuess{Empty: true}
	}

	best := chordGuess{}
	bestScore := -1.0
	for root := 0; root < 12; root++ {
		for _, minor := range []bool{false, true} {
			third := 4
			if minor {
				third = 3
			}
			fit := m[root] + m[(root+third)%12] + m[(root+7)%12]
			score := fit + 0.5*b[root]
			if score > bestScore {
				bestScore = score
				best = chordGuess{Root: root, Minor: minor, Fit: fit}
			}
		}
	}
	return best
}

// Строка таблицы по тактам.
type barRow struct {
	// Номер первого такта строки, с единицы.
	Bar   int
	Start float64
	Bass  chroma
	Mid   chroma
	Chord chordGuess
}

// Разрезать трек на куски по barsPerRow тактов, начиная с первой доли.
// Размер такта считается четырёхдольным.
func barRows(frames []frameChroma, bpm, start float64, barsPerRow int, duration float64) []barRow {
	if bpm <= 0 || barsPerRow <= 0 {
		return nil
	}
	step := float64(barsPerRow) * 4 * 60 / bpm

	var rows []barRow
	for k := 0; ; k++ {
		t0 := start + float64(k)*step
		if t0+step/2 > duration {
			break
		}
		row := barRow{Bar: k*barsPerRow + 1, Start: t0}
		for _, f := range frames {
			if f.Weight == 0 || f.Time < t0 || f.Time >= t0+step {
				continue
			}
			row.Bass.add(f.Bass, 1)
			row.Mid.add(f.Mid, 1)
		}
		row.Chord = guessChord(row.Bass, row.Mid)
		rows = append(rows, row)
	}
	return rows
}

// Сколько раз встретился каждый аккорд, по убыванию.
type chordCount struct {
	Name  string
	Count int
}

func chordSummary(rows []barRow) []chordCount {
	counts := map[string]int{}
	var order []string
	for _, r := range rows {
		if r.Chord.Empty {
			continue
		}
		name := r.Chord.String()
		if counts[name] == 0 {
			order = append(order, name)
		}
		counts[name]++
	}
	out := make([]chordCount, 0, len(order))
	for _, name := range order {
		out = append(out, chordCount{name, counts[name]})
	}
	sort.SliceStable(out, func(a, b int) bool { return out[a].Count > out[b].Count })
	return out
}
