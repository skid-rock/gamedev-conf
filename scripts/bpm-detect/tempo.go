package main

import (
	"math"
	"sort"
)

// Кратности, в которых обычно ошибается автоопределение темпа:
// половинный и удвоенный счёт, а также трёхдольное прочтение.
var octaveFactors = []float64{0.5, 2.0 / 3, 0.75, 4.0 / 3, 1.5, 2, 3}

// Подогнанная сетка: темп, время первой доли и счёт совпадения с атаками.
type tempoFit struct {
	BPM   float64
	Phase float64 // время первой доли сетки от начала файла, сек
	Score float64
}

func (f tempoFit) period() float64 { return 60 / f.BPM }

// Итог анализа одного файла.
type analysis struct {
	Path       string
	Duration   float64
	Fit        tempoFit
	FirstBeat  float64 // время первой сильной доли («единицы»), сек
	BeatOffset int     // на сколько долей сильная доля сдвинута относительно Fit.Phase
	Confidence float64 // во сколько раз счёт лучшей фазы выше типичной
	Downbeat   float64 // запас «единицы» над соперницей, 0..1
	Source     string  // откуда взялась опорная метрика
	Hint       float64 // опорный темп из sidecar или флага, 0 если не было
	TagBPM     float64 // темп из тега файла, как есть
	Drift      drift
	Rivals     []tempoFit // ближайшие соперники, для подробного вывода
}

// Разброс темпа по треку: Suno не секвенсор, сетка может плыть.
type drift struct {
	Windows  []driftWindow
	Min, Max float64
	Spread   float64 // Max-Min по надёжным окнам
	Slip     float64 // на сколько долей уедет единая сетка от начала к концу
	Measured bool
}

type driftWindow struct {
	Start    float64
	BPM      float64
	Score    float64
	Phase    float64 // смещение сетки в этом окне при едином темпе, в долях
	Reliable bool
}

// Значение огибающей в произвольный момент времени, с линейной интерполяцией
// между соседними отсчётами.
func sampleEnv(env []float64, fps, t float64) float64 {
	if t < 0 {
		return 0
	}
	pos := t * fps
	i := int(pos)
	if i+1 >= len(env) {
		if i < len(env) {
			return env[i]
		}
		return 0
	}
	frac := pos - float64(i)
	return env[i]*(1-frac) + env[i+1]*frac
}

// Счёт гребёнки: средний отклик огибающей в точках сетки долей на отрезке
// [from, to]. Именно среднее, а не сумма: иначе быстрый темп выигрывал бы
// только за счёт большего числа долей.
func combScore(env []float64, fps, bpm, phase, from, to float64) float64 {
	period := 60 / bpm
	if period <= 0 {
		return 0
	}
	start := phase
	if start < from {
		k := math.Ceil((from - phase) / period)
		start = phase + k*period
	}
	var sum float64
	var count int
	for t := start; t <= to; t += period {
		sum += sampleEnv(env, fps, t)
		count++
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

// Лучшая фаза для заданного темпа: сначала перебор с шагом в один отсчёт
// огибающей, потом уточнение вокруг найденного.
func bestPhase(env []float64, fps, bpm, from, to float64) (float64, float64) {
	period := 60 / bpm
	coarse := 1 / fps

	best, bestScore := 0.0, -1.0
	for p := 0.0; p < period; p += coarse {
		if s := combScore(env, fps, bpm, p, from, to); s > bestScore {
			best, bestScore = p, s
		}
	}

	fine := coarse / 8
	for p := best - coarse; p <= best+coarse; p += fine {
		if p < 0 {
			continue
		}
		if s := combScore(env, fps, bpm, p, from, to); s > bestScore {
			best, bestScore = p, s
		}
	}
	return math.Mod(best+period, period), bestScore
}

// Лучшая фаза вблизи ожидаемой. Нужна при отслеживании сетки по окнам:
// при полном переборе окно норовит перескочить на оффбит, и ровный трек
// выглядит как уехавший на полдоли.
func bestPhaseNear(env []float64, fps, bpm, center, span, from, to float64) (float64, float64) {
	step := 1 / (fps * 4)
	best, bestScore := center, -1.0
	for p := center - span; p <= center+span; p += step {
		if p < 0 {
			continue
		}
		if s := combScore(env, fps, bpm, p, from, to); s > bestScore {
			best, bestScore = p, s
		}
	}
	return best, bestScore
}

// Грубый проход по всему диапазону темпов.
func scanRange(env []float64, fps, min, max, step, from, to float64) []tempoFit {
	var out []tempoFit
	for bpm := min; bpm <= max; bpm += step {
		phase, score := bestPhase(env, fps, bpm, from, to)
		out = append(out, tempoFit{BPM: bpm, Phase: phase, Score: score})
	}
	return out
}

// Локальные максимумы грубого прохода — кандидаты на уточнение.
func peaksOf(fits []tempoFit, limit int) []tempoFit {
	var peaks []tempoFit
	for i := 1; i < len(fits)-1; i++ {
		if fits[i].Score >= fits[i-1].Score && fits[i].Score > fits[i+1].Score {
			peaks = append(peaks, fits[i])
		}
	}
	sort.Slice(peaks, func(a, b int) bool { return peaks[a].Score > peaks[b].Score })
	if len(peaks) > limit {
		peaks = peaks[:limit]
	}
	return peaks
}

// Уточнение темпа вокруг центра с мелким шагом.
func refine(env []float64, fps, center, span, step, from, to float64) tempoFit {
	best := tempoFit{Score: -1}
	for bpm := center - span; bpm <= center+span; bpm += step {
		if bpm <= 0 {
			continue
		}
		phase, score := bestPhase(env, fps, bpm, from, to)
		if score > best.Score {
			best = tempoFit{BPM: bpm, Phase: phase, Score: score}
		}
	}
	return best
}

// Насколько лучшая фаза выделяется на фоне остальных при том же темпе.
// Единица означает, что сетки в звуке нет вовсе.
func phaseContrast(env []float64, fps, bpm, from, to float64) float64 {
	period := 60 / bpm
	step := 1 / fps
	var scores []float64
	for p := 0.0; p < period; p += step {
		scores = append(scores, combScore(env, fps, bpm, p, from, to))
	}
	if len(scores) == 0 {
		return 1
	}
	sort.Float64s(scores)
	median := scores[len(scores)/2]
	if median <= 0 {
		return 1
	}
	return scores[len(scores)-1] / median
}

// Найти темп и фазу. Подсказки задают опорную метрику: сам по себе счёт
// гребёнки не отличает 86.5 от 173, а вот «в промпте просили 172» отличает.
func findTempo(env onsetEnvelope, hints []float64, min, max float64) (tempoFit, []tempoFit) {
	from, to := 0.0, float64(len(env.all))/env.fps

	coarse := scanRange(env.all, env.fps, min, max, 0.2, from, to)
	centers := peaksOf(coarse, 4)
	for _, h := range hints {
		if h >= min && h <= max {
			centers = append(centers, tempoFit{BPM: h})
		}
	}

	var fits []tempoFit
	seen := map[int]bool{}
	for _, c := range centers {
		key := int(c.BPM * 2)
		if seen[key] {
			continue
		}
		seen[key] = true
		fits = append(fits, refine(env.all, env.fps, c.BPM, 1.5, 0.01, from, to))
	}
	sort.Slice(fits, func(a, b int) bool { return fits[a].Score > fits[b].Score })
	if len(fits) == 0 {
		return tempoFit{}, nil
	}

	best := fits[0]

	// Проверка кратности: если рядом с опорной метрикой стоит вариант почти
	// с тем же счётом, выбираем его. Так 115.33 из тега не побеждает 173.
	if len(hints) > 0 {
		ref := hints[0]
		candidates := []tempoFit{best}
		for _, f := range octaveFactors {
			bpm := best.BPM * f
			if bpm < min || bpm > max {
				continue
			}
			candidates = append(candidates, refine(env.all, env.fps, bpm, 0.6, 0.01, from, to))
		}
		chosen := best
		for _, c := range candidates {
			if c.Score < best.Score*0.85 {
				continue
			}
			if math.Abs(c.BPM-ref) < math.Abs(chosen.BPM-ref) {
				chosen = c
			}
		}
		if chosen.BPM != best.BPM {
			fits = append([]tempoFit{chosen}, fits...)
			best = chosen
		}
	}

	rivals := fits
	if len(rivals) > 4 {
		rivals = rivals[:4]
	}
	return best, rivals
}

// Сильная доля такта. В two-step кик стоит на «единице», а снэйр на «тройке»,
// поэтому смещение выбирается по весу низа, а не по общей громкости.
func findDownbeat(env onsetEnvelope, fit tempoFit) (int, float64, float64) {
	period := fit.period()
	to := float64(len(env.all)) / env.fps

	scores := make([]float64, 4)
	for s := 0; s < 4; s++ {
		phase := fit.Phase + float64(s)*period
		low := combScore(env.low, env.fps, fit.BPM/4, phase, 0, to)
		high := combScore(env.high, env.fps, fit.BPM/4, phase, 0, to)
		// На «единице» бьёт кик, на «тройке» в two-step — снэйр.
		// Поэтому верх идёт в минус, а не в плюс.
		scores[s] = low - 0.5*high
	}

	bestIdx := 0
	for s := 1; s < 4; s++ {
		if scores[s] > scores[bestIdx] {
			bestIdx = s
		}
	}

	// Счёт почти равный — решает начало музыки: трек обычно начинается
	// с «единицы», а не с середины такта.
	start := musicStart(env)
	for s := 0; s < 4; s++ {
		if s == bestIdx || scores[s] < scores[bestIdx]*0.9 {
			continue
		}
		current := firstAfter(fit.Phase+float64(bestIdx)*period, 4*period, start)
		rival := firstAfter(fit.Phase+float64(s)*period, 4*period, start)
		if rival < current {
			bestIdx = s
		}
	}

	second := math.Inf(-1)
	for s := 0; s < 4; s++ {
		if s != bestIdx && scores[s] > second {
			second = scores[s]
		}
	}

	// Запас «единицы» над ближайшей соперницей, в долях размаха счётов:
	// отношение тут не годится, счёт бывает и отрицательным.
	spread := scores[bestIdx] - minOf(scores)
	margin := 0.0
	if spread > 0 {
		margin = (scores[bestIdx] - second) / spread
	}

	return bestIdx, fit.Phase + float64(bestIdx)*period, margin
}

// Темп по окнам: если Suno уводит сетку, одной цифрой дело не обойтись.
func measureDrift(env onsetEnvelope, fit tempoFit, windowSec, hopSec float64) drift {
	total := float64(len(env.all)) / env.fps
	if total < windowSec*1.5 {
		return drift{}
	}

	var windows []driftWindow
	var scores []float64
	for start := 0.0; start+windowSec <= total; start += hopSec {
		w := refine(env.all, env.fps, fit.BPM, 3, 0.02, start, start+windowSec)
		windows = append(windows, driftWindow{Start: start, BPM: w.BPM, Score: w.Score})
		scores = append(scores, w.Score)
	}
	if len(windows) == 0 {
		return drift{}
	}

	// Окно без ударных (интро, брейкдаун) даёт слабый счёт — его темп в разброс
	// не берём, иначе тишина выглядит как скачок темпа.
	sorted := append([]float64(nil), scores...)
	sort.Float64s(sorted)
	median := sorted[len(sorted)/2]
	threshold := median * 0.5

	res := drift{Windows: windows, Min: math.Inf(1), Max: math.Inf(-1)}
	for i := range res.Windows {
		if res.Windows[i].Score >= threshold {
			res.Windows[i].Reliable = true
			res.Measured = true
			res.Min = math.Min(res.Min, res.Windows[i].BPM)
			res.Max = math.Max(res.Max, res.Windows[i].BPM)
		}
	}
	if !res.Measured {
		return drift{Windows: windows}
	}
	res.Spread = res.Max - res.Min
	res.Slip = measureSlip(env, fit, &res, windowSec)
	return res
}

// Насколько уедет единая сетка. Локальный темп по окнам можно списать на
// погрешность, а вот накопленный сдвиг — это ровно то, что слышно в миксе:
// щелчок в конце трека перестаёт попадать в удар.
func measureSlip(env onsetEnvelope, fit tempoFit, d *drift, windowSec float64) float64 {
	period := fit.period()
	var previous float64
	var first, last float64
	var haveFirst bool

	for i := range d.Windows {
		w := &d.Windows[i]
		var phase float64
		if haveFirst {
			// Сетка отслеживается от предыдущего окна и не дальше четверти
			// доли: настоящий уход накапливается постепенно, а перескок
			// на оффбит — это сразу полдоли.
			phase, _ = bestPhaseNear(env.all, env.fps, fit.BPM, previous, period*0.25, w.Start, w.Start+windowSec)
		} else {
			phase, _ = bestPhase(env.all, env.fps, fit.BPM, w.Start, w.Start+windowSec)
		}
		w.Phase = phase / period

		if !w.Reliable {
			continue
		}
		if !haveFirst {
			first, haveFirst = phase, true
		}
		last = phase
		previous = phase
	}
	if !haveFirst {
		return 0
	}
	return math.Abs(last-first) / period
}

// Момент, с которого в треке началась музыка: первый заметный всплеск атак.
func musicStart(env onsetEnvelope) float64 {
	threshold := percentile(env.all, 0.9) * 0.3
	for i, v := range env.all {
		if v >= threshold {
			return env.timeAt(i)
		}
	}
	return 0
}

// Первая точка сетки с заданным шагом, попадающая не раньше указанного момента.
func firstAfter(phase, step, from float64) float64 {
	t := math.Mod(phase, step)
	for t < from {
		t += step
	}
	return t
}

// Наименьшее из значений.
func minOf(values []float64) float64 {
	out := math.Inf(1)
	for _, v := range values {
		if v < out {
			out = v
		}
	}
	return out
}
