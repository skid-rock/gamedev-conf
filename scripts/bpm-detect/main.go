package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
)

type options struct {
	min, max  float64
	hint      float64
	drift     bool
	clickFrom float64
	clickLen  float64
}

func main() {
	opts := options{}
	flag.Float64Var(&opts.min, "min", 80, "нижняя граница поиска темпа")
	flag.Float64Var(&opts.max, "max", 200, "верхняя граница поиска темпа")
	flag.Float64Var(&opts.hint, "hint", 0, "опорный темп, если в sidecar его нет или он врёт")
	noDrift := flag.Bool("no-drift", false, "не считать разброс темпа по окнам (быстрее)")
	mdPath := flag.String("md", "", "записать таблицу в markdown-файл")
	jsonPath := flag.String("json", "", "записать результат в json-файл")
	rename := flag.Bool("rename", false, "дописать темп в имя файла (sidecar переименуется следом)")
	sidecar := flag.Bool("sidecar", false, "дописать темп и первую долю в sidecar-json")
	click := flag.Bool("click", false, "собрать отрывок с щелчками по найденной сетке")
	flag.Float64Var(&opts.clickFrom, "click-from", 0, "с какой секунды начинать отрывок с щелчками")
	flag.Float64Var(&opts.clickLen, "click-len", 45, "длина отрывка с щелчками, секунд (0 — весь трек)")
	clickDir := flag.String("click-dir", "", "куда складывать отрывки (по умолчанию рядом с треком)")
	verbose := flag.Bool("v", false, "показать кандидатов и темп по окнам")
	jobs := flag.Int("jobs", runtime.NumCPU(), "сколько файлов считать одновременно")
	flag.Parse()
	opts.drift = !*noDrift

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Укажите аудиофайлы. Пример: bpm-detect --min 150 --max 200 *.m4a")
		flag.Usage()
		os.Exit(2)
	}

	results := runAll(files, opts, *jobs)
	if len(results) == 0 {
		os.Exit(1)
	}

	printTable(results, *verbose)

	if *mdPath != "" {
		if err := writeMarkdown(*mdPath, results); err != nil {
			fmt.Fprintf(os.Stderr, "markdown: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "\nТаблица записана: %s\n", *mdPath)
		}
	}
	if *jsonPath != "" {
		if err := writeJSON(*jsonPath, results); err != nil {
			fmt.Fprintf(os.Stderr, "json: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "JSON записан: %s\n", *jsonPath)
		}
	}

	// Клики и sidecar считаем до переименования: пути в результатах ещё живые.
	if *click {
		fmt.Fprintln(os.Stderr)
		for _, a := range results {
			out, err := writeClickPreview(a, opts.clickFrom, opts.clickLen, *clickDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "клик %s: %v\n", shortName(a.Path), err)
				continue
			}
			fmt.Fprintf(os.Stderr, "клик: %s\n", out)
		}
	}
	if *sidecar {
		fmt.Fprintln(os.Stderr)
		for _, a := range results {
			path, err := updateSidecar(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "sidecar %s: %v\n", shortName(a.Path), err)
				continue
			}
			fmt.Fprintf(os.Stderr, "sidecar: %s\n", path)
		}
	}
	if *rename {
		fmt.Fprintln(os.Stderr)
		for _, a := range results {
			path, err := renameWithBPM(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "переименование %s: %v\n", shortName(a.Path), err)
				continue
			}
			fmt.Fprintf(os.Stderr, "имя: %s\n", shortName(path))
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
					fmt.Fprintf(os.Stderr, "[%d/%d] %s — %s BPM\n", done, len(files), shortName(files[i]), formatBPM(res.Fit.BPM))
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
	env := computeOnset(samples)

	hints, ref, tag, source := collectHints(path, opts.hint, opts.min, opts.max)
	fit, rivals := findTempo(env, hints, opts.min, opts.max)
	if fit.BPM == 0 {
		return analysis{}, fmt.Errorf("темп не определился")
	}

	offset, firstBeat, downbeatRatio := findDownbeat(env, fit)
	res := analysis{
		Path:       path,
		Duration:   durationOf(samples),
		Fit:        fit,
		FirstBeat:  firstBeat,
		BeatOffset: offset,
		Confidence: phaseContrast(env.all, env.fps, fit.BPM, 0, durationOf(samples)),
		Downbeat:   downbeatRatio,
		Source:     source,
		Hint:       ref,
		TagBPM:     tag,
		Rivals:     rivals,
	}
	if opts.drift {
		res.Drift = measureDrift(env, fit, 30, 15)
	}
	return res, nil
}

// Результат в json — для скриптов и для ReaScript.
func writeJSON(path string, results []analysis) error {
	type record struct {
		File       string  `json:"file"`
		BPM        float64 `json:"bpm"`
		FirstBeat  float64 `json:"first_beat"`
		BeatPhase  float64 `json:"beat_phase"`
		Drift      float64 `json:"bpm_drift"`
		Slip       float64 `json:"grid_slip_beats"`
		Confidence float64 `json:"bpm_confidence"`
		Downbeat   float64 `json:"downbeat_ratio"`
		Source     string  `json:"bpm_source"`
		Hint       float64 `json:"hint"`
		TagBPM     float64 `json:"tag_bpm"`
		Duration   float64 `json:"duration"`
		Verdict    string  `json:"verdict"`
	}
	records := make([]record, 0, len(results))
	for _, a := range results {
		records = append(records, record{
			File:       a.Path,
			BPM:        a.Fit.BPM,
			FirstBeat:  a.FirstBeat,
			BeatPhase:  a.Fit.Phase,
			Drift:      a.Drift.Spread,
			Slip:       a.Drift.Slip,
			Confidence: a.Confidence,
			Downbeat:   a.Downbeat,
			Source:     a.Source,
			Hint:       a.Hint,
			TagBPM:     a.TagBPM,
			Duration:   a.Duration,
			Verdict:    verdict(a),
		})
	}
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}
