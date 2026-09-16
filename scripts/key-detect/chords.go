package main

import (
	"sort"
)

// Качество аккорда.
type chordQuality int

const (
	majorChord chordQuality = iota
	minorChord
	// Квинта без терции: мажор это или минор, по звуку не сказать.
	powerChord
)

// Порог доли терции в середине: тише — терции в аккорде нет.
const thirdThreshold = 0.04

// Бас на корне или на другой ноте аккорда, если его доля в басу не меньше этой.
const inversionShare = 0.5

// Догадка об аккорде: трезвучие или квинта, с басом.
type chordGuess struct {
	Root    int
	Quality chordQuality
	// Нота баса, если это не корень, иначе −1.
	Bass int
	// Доля энергии середины внутри аккорда, 0..1.
	Fit   float64
	Empty bool
}

// Название без баса: «Am», «E5».
func (g chordGuess) Name() string {
	if g.Empty {
		return "—"
	}
	name := noteNames[g.Root]
	switch g.Quality {
	case minorChord:
		name += "m"
	case powerChord:
		name += "5"
	}
	return name
}

// Название с басом: «Am/E».
func (g chordGuess) String() string {
	if g.Empty || g.Bass < 0 {
		return g.Name()
	}
	return g.Name() + "/" + noteNames[g.Bass]
}

// Подобрать трезвучие по середине, с доводом от баса. Нота баса на корне весит
// больше, чем на терции или квинте: так Am отличается от C, у которых две общие
// ноты, но Am/E (бас на квинте) не превращается в E.
func guessChord(bass, mid chroma) chordGuess {
	b := bass.normalized()
	m := mid.normalized()
	if m.sum() == 0 && b.sum() == 0 {
		return chordGuess{Empty: true}
	}

	best := chordGuess{}
	bestScore := -1.0
	for root := 0; root < 12; root++ {
		fifth := (root + 7) % 12
		for _, q := range []chordQuality{majorChord, minorChord} {
			third := (root + thirdStep(q)) % 12
			fit := m[root] + m[third] + m[fifth]
			score := fit + 0.5*b[root] + 0.25*(b[third]+b[fifth])
			if score > bestScore {
				bestScore = score
				best = chordGuess{Root: root, Quality: q, Fit: fit}
			}
		}
	}

	// Ни малой, ни большой терции не слышно ни сверху, ни в басу — это квинта.
	minor3, major3 := (best.Root+3)%12, (best.Root+4)%12
	if m[minor3] < thirdThreshold && m[major3] < thirdThreshold &&
		b[minor3] < thirdThreshold && b[major3] < thirdThreshold {
		best.Quality = powerChord
		best.Fit = m[best.Root] + m[(best.Root+7)%12]
	}

	best.Bass = -1
	if b.sum() > 0 {
		low := b.top(1)[0]
		if low != best.Root && b[low] >= inversionShare && best.hasTone(low) {
			best.Bass = low
		}
	}
	return best
}

func thirdStep(q chordQuality) int {
	if q == minorChord {
		return 3
	}
	return 4
}

// Входит ли нота в аккорд.
func (g chordGuess) hasTone(pc int) bool {
	step := (pc - g.Root + 12) % 12
	switch {
	case step == 0 || step == 7:
		return true
	case g.Quality == powerChord:
		return false
	default:
		return step == thirdStep(g.Quality)
	}
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
	step := float64(barsPerRow) * barSeconds(bpm)

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

// Длина четырёхдольного такта в секундах.
func barSeconds(bpm float64) float64 {
	return 4 * 60 / bpm
}

// Сколько раз встретился каждый аккорд, по убыванию. Обращения считаются
// вместе с основным видом: сводка нужна, чтобы увидеть вращение гармонии.
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
		name := r.Chord.Name()
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
