package main

import (
	"math"
	"sort"
)

var noteNames = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

// Названия ступеней от тоники по полутонам.
var degreeNames = [12]string{"1", "b2", "2", "b3", "3", "4", "#4", "5", "b6", "6", "b7", "7"}

// Хрома: энергия по двенадцати классам высоты, индекс 0 — C.
type chroma [12]float64

func (c chroma) sum() float64 {
	var s float64
	for _, v := range c {
		s += v
	}
	return s
}

// Доли вместо абсолютной энергии; пустая хрома остаётся пустой.
func (c chroma) normalized() chroma {
	s := c.sum()
	if s == 0 {
		return c
	}
	for i := range c {
		c[i] /= s
	}
	return c
}

func (c *chroma) add(o chroma, weight float64) {
	for i := range c {
		c[i] += o[i] * weight
	}
}

// Тональная ясность 0..1: единица минус нормированная энтропия.
// Ровная хрома (шум, ударные) даёт ноль, одна нота — единицу, трезвучие — около 0.56.
func (c chroma) clarity() float64 {
	n := c.normalized()
	if n.sum() == 0 {
		return 0
	}
	var h float64
	for _, v := range n {
		if v > 0 {
			h -= v * math.Log(v)
		}
	}
	return 1 - h/math.Log(12)
}

// Индексы самых громких классов по убыванию.
func (c chroma) top(n int) []int {
	idx := make([]int, 12)
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool { return c[idx[a]] > c[idx[b]] })
	return idx[:n]
}

// Разложить пики полосы lo..hi по классам высоты с поправкой на строй.
// Пики дальше 30 центов от полутона пропускаются: это скорее шум или глиссандо.
func foldPeaks(ps []peak, lo, hi, tune float64) chroma {
	var c chroma
	for _, p := range ps {
		if p.Freq < lo || p.Freq >= hi {
			continue
		}
		m := 69 + 12*math.Log2(p.Freq/440) - tune
		r := math.Round(m)
		if math.Abs(m-r) > 0.3 {
			continue
		}
		pc := (int(r)%12 + 12) % 12
		c[pc] += p.Weight
	}
	return c
}

// Хрома одного кадра, отдельно для баса и середины.
type frameChroma struct {
	Time float64
	Bass chroma
	Mid  chroma
	// Вес кадра в общем профиле: тональная ясность; у тихих кадров — ноль.
	Weight float64
}

// Порог тишины относительно самого громкого кадра, по сумме амплитуд пиков (−34 дБ).
const silenceRatio = 0.02

func foldFrames(frames [][]peak, tune float64) []frameChroma {
	out := make([]frameChroma, len(frames))
	var loudest float64
	for i, ps := range frames {
		out[i] = frameChroma{
			Time: frameTime(i),
			Bass: foldPeaks(ps, lowFreq, bassSplit, tune),
			Mid:  foldPeaks(ps, bassSplit, highFreq, tune),
		}
		loudest = max(loudest, out[i].Bass.sum()+out[i].Mid.sum())
	}
	for i := range out {
		f := &out[i]
		if f.Bass.sum()+f.Mid.sum() < loudest*silenceRatio {
			continue
		}
		var both chroma
		both.add(f.Bass.normalized(), 1)
		both.add(f.Mid.normalized(), 1)
		f.Weight = both.clarity()
	}
	return out
}

// Профили трека: бас, середина и их равная смесь.
// Каждый кадр входит долями, а не энергией: иначе громкие куски
// перевешивают тихие, а тональность от громкости не зависит.
func summarize(frames []frameChroma) (bass, mid, total chroma) {
	for _, f := range frames {
		if f.Weight == 0 {
			continue
		}
		bass.add(f.Bass.normalized(), f.Weight)
		mid.add(f.Mid.normalized(), f.Weight)
	}
	bass = bass.normalized()
	mid = mid.normalized()
	total.add(bass, 1)
	total.add(mid, 1)
	return bass, mid, total.normalized()
}
