package main

import (
	"math"
	"sort"
)

// Шаг огибающей в отсчётах исходного потока: 64 при 11025 Гц дают ~172 значения
// в секунду, то есть разрешение по времени около 5.8 мс.
const onsetHop = 64

// Окно, по которому считается громкость, — четыре шага, около 23 мс.
// Короче нельзя: период кика на 50 Гц равен 20 мс, и в коротком окне
// вместо атаки меряется сама волна.
const onsetWindowBlocks = 4

// Границы полос. Разделение грубое, однополюсными фильтрами: важно не
// качество среза, а то, что кик, снэйр и хэты дают вклад по отдельности
// и тихая полоса не тонет в громкой.
var bandCutoffs = []float64{150, 400, 1200, 3500}

// Огибающая атак: функция новизны, по которой ищется сетка долей.
// Низ и верх хранятся отдельно — дальше по ним отличается сильная доля
// такта от слабой.
type onsetEnvelope struct {
	fps  float64   // значений в секунду
	all  []float64 // суммарная функция новизны
	low  []float64 // вклад низа (кик, бас)
	high []float64 // вклад верха (снэйр, хэты)
}

// Позиция отсчёта огибающей во времени.
func (e onsetEnvelope) timeAt(index int) float64 {
	return float64(index) / e.fps
}

// Построить огибающую атак.
func computeOnset(samples []float32) onsetEnvelope {
	bands := len(bandCutoffs) + 1
	frames := len(samples) / onsetHop

	// Коэффициенты однополюсных фильтров нижних частот.
	alpha := make([]float64, len(bandCutoffs))
	for i, fc := range bandCutoffs {
		alpha[i] = 1 - math.Exp(-2*math.Pi*fc/analysisRate)
	}
	state := make([]float64, len(bandCutoffs))

	// Сумма квадратов по каждой полосе внутри блока в onsetHop отсчётов.
	blockSum := make([][]float64, bands)
	for b := range blockSum {
		blockSum[b] = make([]float64, frames)
	}

	for f := 0; f < frames; f++ {
		base := f * onsetHop
		for i := 0; i < onsetHop; i++ {
			x := float64(samples[base+i])
			rest := x
			for b := 0; b < len(alpha); b++ {
				state[b] += alpha[b] * (rest - state[b])
				band := state[b]
				blockSum[b][f] += band * band
				rest -= band
			}
			blockSum[bands-1][f] += rest * rest
		}
	}

	// Громкость в окне: блоки складываются внахлёст.
	rms := make([][]float64, bands)
	for b := 0; b < bands; b++ {
		rms[b] = make([]float64, frames)
		for f := 0; f < frames; f++ {
			var sum float64
			count := 0
			for k := 0; k < onsetWindowBlocks; k++ {
				idx := f - k
				if idx < 0 {
					break
				}
				sum += blockSum[b][idx]
				count++
			}
			rms[b][f] = math.Sqrt(sum / float64(count*onsetHop))
		}
	}

	novelty := make([][]float64, bands)
	for b := 0; b < bands; b++ {
		novelty[b] = noveltyFrom(rms[b])
	}

	all := make([]float64, frames)
	low := make([]float64, frames)
	high := make([]float64, frames)
	for f := 0; f < frames; f++ {
		for b := 0; b < bands; b++ {
			all[f] += novelty[b][f]
			if b < 2 {
				low[f] += novelty[b][f]
			} else {
				high[f] += novelty[b][f]
			}
		}
	}

	return onsetEnvelope{
		fps:  float64(analysisRate) / float64(onsetHop),
		all:  all,
		low:  low,
		high: high,
	}
}

// Из огибающей громкости сделать функцию новизны: сжать динамику,
// взять положительный прирост, вычесть локальный фон.
// Локальный фон важен: без него тихий брейкдаун не даёт вклада вовсе,
// а громкий дроп перевешивает весь остальной трек.
func noveltyFrom(rms []float64) []float64 {
	scale := percentile(rms, 0.95)
	if scale <= 0 {
		return make([]float64, len(rms))
	}

	compressed := make([]float64, len(rms))
	for i, v := range rms {
		compressed[i] = math.Log1p(10 * v / scale)
	}

	// Прирост берётся через шаг окна, а не через соседний отсчёт: окно
	// перекрывается, и у соседей разница почти всегда вырождается в шум.
	diff := make([]float64, len(rms))
	for i := onsetWindowBlocks; i < len(compressed); i++ {
		d := compressed[i] - compressed[i-onsetWindowBlocks]
		if d > 0 {
			diff[i] = d
		}
	}

	framesPerSecond := float64(analysisRate) / float64(onsetHop)
	background := movingAverage(diff, int(1.0*framesPerSecond))

	out := make([]float64, len(diff))
	for i := range diff {
		v := diff[i] - background[i]
		if v > 0 {
			out[i] = v
		}
	}
	return out
}

// Скользящее среднее по окну, центрированное, с усечением на краях.
func movingAverage(values []float64, window int) []float64 {
	if window < 3 {
		window = 3
	}
	prefix := make([]float64, len(values)+1)
	for i, v := range values {
		prefix[i+1] = prefix[i] + v
	}

	half := window / 2
	out := make([]float64, len(values))
	for i := range values {
		lo := i - half
		if lo < 0 {
			lo = 0
		}
		hi := i + half + 1
		if hi > len(values) {
			hi = len(values)
		}
		out[i] = (prefix[hi] - prefix[lo]) / float64(hi-lo)
	}
	return out
}

// Значение заданного квантиля.
func percentile(values []float64, q float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	idx := int(q * float64(len(sorted)-1))
	return sorted[idx]
}
