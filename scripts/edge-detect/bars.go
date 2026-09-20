package main

import (
	"math"
	"slices"
)

// Пороги в децибелах относительно тела трека (медианы тактов между интро и аутро).
const (
	// Слышно почти как в теле трека.
	presentDB = -8.0
	// Ниже этого партии, считай, нет.
	absentDB = -15.0
	// Вход партии: такт не тише этого, и следующий тоже.
	entryDB = -6.0
	// Последний звучащий такт: не тише максимума трека на столько.
	soundDB = -40.0
	// Конец: последний такт тише тела на столько — стоп; громче stopLoudDB — звук до конца.
	stopQuietDB = -20.0
	stopLoudDB  = -6.0
	// Середина тише этого — о нотах судить не по чему.
	midQuietDB = -12.0
	// Ровность середины ниже этой доли от тела — там ноты, а не удары и шум.
	// Порог подобран на четырёх дублях A10: интро с нотами дают 0.07–0.43.
	tonalRatio = 0.5
)

// Такт: уровни полос в дБ и ровность середины.
type barStat struct {
	Index int
	Start float64
	Level [bandCount]float64
	Flat  float64
}

// Собрать кадры в такты по сетке. Неполный последний такт отбрасывается.
func groupBars(frames []frame, bpm, start float64) []barStat {
	barLen := 240 / bpm
	var peak float64
	for _, f := range frames {
		peak = math.Max(peak, f.Energy[bandAll])
	}

	type acc struct {
		sum   [bandCount]float64
		n     int
		flats []float64
	}
	var accs []acc
	for _, f := range frames {
		if f.Time < start {
			continue
		}
		i := int((f.Time - start) / barLen)
		for len(accs) <= i {
			accs = append(accs, acc{})
		}
		a := &accs[i]
		for b := range f.Energy {
			a.sum[b] += f.Energy[b]
		}
		a.n++
		// Ровность тихих кадров ничего не говорит: там шум кодека.
		if !math.IsNaN(f.Flatness) && f.Energy[bandAll] > peak*1e-5 {
			a.flats = append(a.flats, f.Flatness)
		}
	}

	full := barLen * analysisRate / hopSize
	var bars []barStat
	for i, a := range accs {
		if a.n == 0 || (i == len(accs)-1 && float64(a.n) < full*0.9) {
			continue
		}
		s := barStat{Index: i, Start: start + float64(i)*barLen, Flat: math.NaN()}
		for b := range s.Level {
			s.Level[b] = toDB(a.sum[b] / float64(a.n))
		}
		if len(a.flats) > 0 {
			s.Flat = median(a.flats)
		}
		bars = append(bars, s)
	}
	return bars
}

func toDB(power float64) float64 {
	return 10 * math.Log10(power+1e-20)
}

func median(xs []float64) float64 {
	s := slices.Clone(xs)
	slices.Sort(s)
	return s[len(s)/2]
}

// Сводка по отрезку тактов относительно тела трека.
type span struct {
	From, To int // номера тактов с единицы, включительно
	Rel      [bandCount]float64
	// Ровность середины, делённая на ровность тела: меньше единицы — больше нот.
	FlatRatio float64
	// Сколько тактов отрезка с партией в полосе и сколько с нотами в середине.
	// Медиана врёт, когда половина отрезка с барабанами, а половина без, — счёт нет.
	Present [bandCount]int
	Tonal   int
	N       int
}

// Разбор краёв одного трека.
type edges struct {
	Bars []barStat
	Body [bandCount]float64
	Flat float64

	Intro, Outro, End span
	HasOutro, HasEnd  bool
	// Последний звучащий такт, с единицы.
	LastBar int
	// Первый такт после начала, где вошли саб и хэты; 0 — не вошли.
	SubEntry, HighEntry int
}

func findEdges(bars []barStat, p plan) edges {
	e := edges{Bars: bars}
	if len(bars) == 0 {
		return e
	}

	var peak float64 = math.Inf(-1)
	for _, b := range bars {
		peak = math.Max(peak, b.Level[bandAll])
	}
	last := len(bars) - 1
	for last > 0 && bars[last].Level[bandAll] < peak+soundDB {
		last--
	}
	e.LastBar = last + 1

	// Конец и аутро отсчитываются от последнего звучащего такта, интро — от начала.
	endFrom := last - p.End + 1
	outroFrom := endFrom - p.Outro
	bodyFrom, bodyTo := p.Intro, outroFrom-1
	if bodyTo-bodyFrom < 3 {
		bodyFrom, bodyTo = 0, last
	}

	body := bars[max(0, bodyFrom) : bodyTo+1]
	for b := range e.Body {
		vals := make([]float64, len(body))
		for i, s := range body {
			vals[i] = s.Level[b]
		}
		e.Body[b] = median(vals)
	}
	e.Flat = medianFlat(body)

	e.Intro = e.spanOf(0, p.Intro-1)
	if p.Outro > 0 && outroFrom >= 0 {
		e.Outro, e.HasOutro = e.spanOf(outroFrom, endFrom-1), true
	}
	if p.End > 0 && endFrom >= 0 {
		e.End, e.HasEnd = e.spanOf(endFrom, last), true
	}

	e.SubEntry = e.entry(bandSub)
	e.HighEntry = e.entry(bandHigh)
	return e
}

// Сводка по тактам from..to (с нуля, включительно): медианы уровней и ровность.
func (e edges) spanOf(from, to int) span {
	from = max(0, from)
	to = min(len(e.Bars)-1, to)
	s := span{From: from + 1, To: to + 1, FlatRatio: math.NaN()}
	if from > to {
		return s
	}
	part := e.Bars[from : to+1]
	for b := range s.Rel {
		vals := make([]float64, len(part))
		for i, x := range part {
			vals[i] = x.Level[b] - e.Body[b]
		}
		s.Rel[b] = median(vals)
	}
	if f := medianFlat(part); !math.IsNaN(f) && e.Flat > 0 {
		s.FlatRatio = f / e.Flat
	}
	s.N = len(part)
	for _, x := range part {
		for b := range s.Present {
			if x.Level[b]-e.Body[b] >= presentDB {
				s.Present[b]++
			}
		}
		if x.tonal(e) {
			s.Tonal++
		}
	}
	return s
}

func medianFlat(bars []barStat) float64 {
	var vals []float64
	for _, b := range bars {
		if !math.IsNaN(b.Flat) {
			vals = append(vals, b.Flat)
		}
	}
	if len(vals) == 0 {
		return math.NaN()
	}
	return median(vals)
}

// Первый такт, с которого полоса звучит почти как в теле два такта подряд.
func (e edges) entry(b band) int {
	for i := 0; i+1 < len(e.Bars); i++ {
		if e.Bars[i].Level[b]-e.Body[b] >= entryDB && e.Bars[i+1].Level[b]-e.Body[b] >= entryDB {
			return i + 1
		}
	}
	return 0
}

// Последний такт относительно тела: по нему видно, остановился ли трек.
func (e edges) lastRel() float64 {
	return e.Bars[e.LastBar-1].Level[bandAll] - e.Body[bandAll]
}

// Ноты в середине такта: середина слышна и спектр в ней с выступающими пиками.
func (b barStat) tonal(e edges) bool {
	return b.Level[bandMid]-e.Body[bandMid] > midQuietDB && !math.IsNaN(b.Flat) && e.Flat > 0 && b.Flat/e.Flat < tonalRatio
}
