package main

import (
	"math"
	"math/cmplx"
	"slices"
)

const (
	// Окно 0.74 с даёт шаг 1.35 Гц: этого хватает, чтобы различить полутоны
	// даже в басу, где между A1 и A#1 всего 3.3 Гц.
	frameSize = 8192
	hopSize   = 1024

	// Диапазон, в котором ищутся ноты, и граница баса с серединой, Гц.
	lowFreq   = 30.0
	highFreq  = 2500.0
	bassSplit = 200.0

	// Пик считается нотой, если он во столько раз выше медианы соседних бинов.
	// Так отсекаются шум щёток и шейкера: у них спектр ровный, без выступов.
	prominenceHalf  = 10
	prominenceRatio = 4.0
)

// Спектральный пик: частота с уточнением между бинами и вес (амплитуда).
type peak struct {
	Freq   float64
	Weight float64
}

// Пики всех кадров трека.
func framePeaks(samples []float32) [][]peak {
	window := hann(frameSize)
	buf := make([]complex128, frameSize)
	lo := int(math.Ceil(lowFreq * frameSize / analysisRate))
	hi := int(math.Floor(highFreq * frameSize / analysisRate))
	mag := make([]float64, hi+prominenceHalf+2)

	var frames [][]peak
	for off := 0; off+frameSize <= len(samples); off += hopSize {
		for i := range buf {
			buf[i] = complex(float64(samples[off+i])*window[i], 0)
		}
		fft(buf)
		for i := range mag {
			mag[i] = cmplx.Abs(buf[i])
		}
		frames = append(frames, pickPeaks(mag, lo, hi))
	}
	return frames
}

// Время середины кадра, секунды.
func frameTime(index int) float64 {
	return (float64(index*hopSize) + frameSize/2) / analysisRate
}

// Выбрать выступающие пики кадра и уточнить их частоту параболой по логарифму амплитуды.
func pickPeaks(mag []float64, lo, hi int) []peak {
	var ps []peak
	local := make([]float64, 0, 2*prominenceHalf+1)
	for i := lo; i < hi; i++ {
		if mag[i] <= 0 || mag[i] <= mag[i-1] || mag[i] < mag[i+1] {
			continue
		}

		local = local[:0]
		for k := max(0, i-prominenceHalf); k <= i+prominenceHalf; k++ {
			local = append(local, mag[k])
		}
		slices.Sort(local)
		if mag[i] < prominenceRatio*local[len(local)/2] {
			continue
		}

		a := math.Log(mag[i-1] + 1e-12)
		b := math.Log(mag[i] + 1e-12)
		c := math.Log(mag[i+1] + 1e-12)
		shift := 0.0
		if den := a - 2*b + c; den != 0 {
			shift = 0.5 * (a - c) / den
		}
		ps = append(ps, peak{
			Freq:   (float64(i) + shift) * analysisRate / frameSize,
			Weight: mag[i],
		})
	}
	suppressOvertones(ps)
	return ps
}

// Приглушить пики, похожие на обертоны более низкой ноты того же кадра.
// Третий и шестой обертоны дают квинту, пятый — большую терцию, седьмой — почти
// малую септиму. Без этого бас на A подсвечивает C# и G и сдвигает лад
// к мажору или миксолидийскому. Чётные степени двойки высоту класса не меняют.
func suppressOvertones(ps []peak) {
	for i := range ps {
		for j := 0; j < i; j++ {
			base := ps[j]
			if base.Freq > ps[i].Freq/2.9 {
				break
			}
			// Обертон, втрое громче основного тона, — скорее самостоятельная нота.
			if base.Weight*3 < ps[i].Weight {
				continue
			}
			ratio := ps[i].Freq / base.Freq
			h := math.Round(ratio)
			if h != 3 && h != 5 && h != 6 && h != 7 {
				continue
			}
			if math.Abs(1200*math.Log2(ratio/h)) < 20 {
				ps[i].Weight *= 0.25
				break
			}
		}
	}
}

// Отклонение строя от A440 в полутонах: круговое среднее дробной части высоты пиков.
// Suno нередко строит на 10–30 центов выше или ниже, и без поправки
// ноты на границе полутона уходят в соседний класс.
func estimateTuning(frames [][]peak) float64 {
	var x, y float64
	for _, ps := range frames {
		for _, p := range ps {
			if p.Freq < 80 || p.Freq > 2000 {
				continue
			}
			m := 69 + 12*math.Log2(p.Freq/440)
			angle := 2 * math.Pi * (m - math.Round(m))
			x += p.Weight * math.Cos(angle)
			y += p.Weight * math.Sin(angle)
		}
	}
	return math.Atan2(y, x) / (2 * math.Pi)
}

func hann(n int) []float64 {
	w := make([]float64, n)
	for i := range w {
		w[i] = 0.5 - 0.5*math.Cos(2*math.Pi*float64(i)/float64(n))
	}
	return w
}

// Быстрое преобразование Фурье на месте, длина — степень двойки.
func fft(a []complex128) {
	n := len(a)
	for i, j := 1, 0; i < n; i++ {
		bit := n >> 1
		for ; j&bit != 0; bit >>= 1 {
			j ^= bit
		}
		j ^= bit
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	for size := 2; size <= n; size <<= 1 {
		step := cmplx.Exp(complex(0, -2*math.Pi/float64(size)))
		half := size / 2
		for start := 0; start < n; start += size {
			w := complex(1, 0)
			for k := 0; k < half; k++ {
				u, v := a[start+k], a[start+k+half]*w
				a[start+k], a[start+k+half] = u+v, u-v
				w *= step
			}
		}
	}
}
