package main

import (
	"math"
	"math/cmplx"
)

const (
	// Окно 93 мс с шагом 46 мс: на такт при 107 BPM приходится около 48 кадров.
	frameSize = 2048
	hopSize   = 1024
)

// Полосы, по которым видно, кто играет:
// саб — бочка и саб-бас, низ-середина — тело малого и бас-стабы,
// середина — аккорды, клавиши, стабы, верх — хэты и шейкер.
type band int

const (
	bandSub band = iota
	bandLowMid
	bandMid
	bandHigh
	bandAll
	bandCount
)

var bandNames = [bandCount]string{"саб", "низ-сер", "сер", "верх", "всё"}

// Границы полос в герцах.
var bandEdges = [bandCount][2]float64{
	{30, 90},
	{150, 500},
	{500, 3000},
	{6000, 11000},
	{30, 11000},
}

// Диапазон, в котором считается ровность спектра середины.
const flatLow, flatHigh = 200.0, 4000.0

// Один кадр: энергия по полосам и ровность середины.
type frame struct {
	Time   float64
	Energy [bandCount]float64
	// Спектральная ровность середины: геометрическое среднее мощности, делённое на арифметическое.
	// Около нуля — выступающие пики, то есть ноты; ближе к единице — шум и удары.
	Flatness float64
}

// Посчитать кадры всего трека.
func analyzeFrames(samples []float32) []frame {
	window := hann(frameSize)
	buf := make([]complex128, frameSize)
	power := make([]float64, frameSize/2+1)

	var lo, hi [bandCount]int
	for b := range bandEdges {
		lo[b] = binOf(bandEdges[b][0])
		hi[b] = binOf(bandEdges[b][1])
	}
	flo, fhi := binOf(flatLow), binOf(flatHigh)

	var frames []frame
	for off := 0; off+frameSize <= len(samples); off += hopSize {
		for i := range buf {
			buf[i] = complex(float64(samples[off+i])*window[i], 0)
		}
		fft(buf)
		for i := range power {
			a := cmplx.Abs(buf[i])
			power[i] = a * a
		}

		f := frame{Time: (float64(off) + frameSize/2) / analysisRate}
		for b := range bandEdges {
			for i := lo[b]; i <= hi[b]; i++ {
				f.Energy[b] += power[i]
			}
		}
		f.Flatness = flatness(power[flo : fhi+1])
		frames = append(frames, f)
	}
	return frames
}

func binOf(freq float64) int {
	return int(math.Round(freq * frameSize / analysisRate))
}

// Ровность спектра; для тишины — NaN, чтобы её не учитывать.
func flatness(p []float64) float64 {
	var logSum, sum float64
	for _, v := range p {
		logSum += math.Log(v + 1e-20)
		sum += v
	}
	mean := sum / float64(len(p))
	if mean < 1e-12 {
		return math.NaN()
	}
	return math.Exp(logSum/float64(len(p))) / mean
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
