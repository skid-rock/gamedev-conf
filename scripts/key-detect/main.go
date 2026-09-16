package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
)

type options struct {
	bars  int
	bpm   float64
	start float64
}

// Результат разбора одного файла.
type analysis struct {
	Path     string
	Duration float64
	// Отклонение строя от A440 в полутонах.
	Tuning float64

	Bass, Mid, Total chroma
	Ranked           []keyCandidate

	Declared    keyCandidate
	HasDeclared bool

	// Сетка для таблицы по тактам и откуда она взята.
	BPM, Start float64
	GridSource string
	Rows       []barRow
}

func (a analysis) best() keyCandidate {
	return a.Ranked[0]
}

func main() {
	opts := options{}
	verbose := flag.Bool("v", false, "показать профили нот, опознавательные ступени и аккорды")
	flag.IntVar(&opts.bars, "bars", 0, "таблица по тактам: сколько тактов в строке (0 — не показывать)")
	flag.Float64Var(&opts.bpm, "bpm", 0, "темп для разбивки на такты (по умолчанию из sidecar)")
	flag.Float64Var(&opts.start, "start", -1, "секунда первой доли (по умолчанию из sidecar, иначе 0)")
	top := flag.Int("top", 5, "сколько кандидатов показывать в подробном выводе")
	jsonPath := flag.String("json", "", "записать результат в json-файл")
	jobs := flag.Int("jobs", runtime.NumCPU(), "сколько файлов считать одновременно")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Укажите аудиофайлы. Пример: key-detect -v *.m4a")
		flag.Usage()
		os.Exit(2)
	}

	results := runAll(files, opts, *jobs)
	if len(results) == 0 {
		os.Exit(1)
	}

	printTable(results)
	if *verbose {
		for _, a := range results {
			printDetails(a, *top)
		}
	}
	if opts.bars > 0 {
		for _, a := range results {
			printBars(a, opts.bars)
		}
	}

	if *jsonPath != "" {
		if err := writeJSON(*jsonPath, results, *top); err != nil {
			fmt.Fprintf(os.Stderr, "json: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "\nJSON записан: %s\n", *jsonPath)
		}
	}
}

// Посчитать все файлы, сохранив порядок аргументов.
func runAll(files []string, opts options, jobs int) []analysis {
	if jobs < 1 {
		jobs = 1
	}
	type item struct {
		index int
		res   analysis
	}

	in := make(chan int)
	out := make(chan item)
	var wg sync.WaitGroup

	var done int
	var mu sync.Mutex

	for w := 0; w < jobs; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range in {
				res, err := analyze(files[i], opts)
				mu.Lock()
				done++
				if err != nil {
					fmt.Fprintf(os.Stderr, "[%d/%d] %s — ошибка: %v\n", done, len(files), shortName(files[i]), err)
				} else {
					fmt.Fprintf(os.Stderr, "[%d/%d] %s — %s\n", done, len(files), shortName(files[i]), res.best().Name())
				}
				mu.Unlock()
				if err == nil {
					out <- item{index: i, res: res}
				}
			}
		}()
	}

	go func() {
		for i := range files {
			in <- i
		}
		close(in)
		wg.Wait()
		close(out)
	}()

	var items []item
	for it := range out {
		items = append(items, it)
	}
	sort.Slice(items, func(a, b int) bool { return items[a].index < items[b].index })

	results := make([]analysis, 0, len(items))
	for _, it := range items {
		results = append(results, it.res)
	}
	fmt.Fprintln(os.Stderr)
	return results
}

// Полный разбор одного файла.
func analyze(path string, opts options) (analysis, error) {
	samples, err := decodeMono(path)
	if err != nil {
		return analysis{}, err
	}
	peaks := framePeaks(samples)
	tune := estimateTuning(peaks)
	frames := foldFrames(peaks, tune)
	bass, mid, total := summarize(frames)
	if total.sum() == 0 {
		return analysis{}, fmt.Errorf("нот не нашлось: трек тихий или состоит из шума")
	}

	meta := readSidecar(path)
	res := analysis{
		Path:     path,
		Duration: durationOf(samples),
		Tuning:   tune,
		Bass:     bass,
		Mid:      mid,
		Total:    total,
		Ranked:   rankKeys(total),
	}
	res.Declared, res.HasDeclared = meta.declaredKey()
	res.BPM, res.Start, res.GridSource = resolveGrid(meta, opts)
	res.Rows = barRows(frames, res.BPM, res.Start, 1, res.Duration)
	return res, nil
}

// Сетка тактов: флаги сильнее sidecar, измерение bpm-detect сильнее промпта.
func resolveGrid(meta sidecar, opts options) (bpm, start float64, source string) {
	switch {
	case opts.bpm > 0:
		bpm, source = opts.bpm, "флаг --bpm"
	case meta.BPM > 0:
		bpm, source = meta.BPM, "bpm-detect (sidecar)"
	default:
		bpm, source = meta.declaredBPM(), "промпт, первая доля не измерена"
	}

	switch {
	case opts.start >= 0:
		start = opts.start
		source += ", первая доля из --start"
	case meta.HasFirst:
		start = meta.FirstBeat
	}
	return bpm, start, source
}
