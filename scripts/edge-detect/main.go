package main

import (
	"flag"
	"fmt"
	"os"
)

type options struct {
	bpm               float64
	start             float64
	intro, outro, end int
}

// Результат разбора одного файла.
type result struct {
	Path       string
	BPM, Start float64
	GridSource string
	Plan       plan
	Edges      edges
}

func main() {
	opts := options{}
	verbose := flag.Bool("v", false, "показать потактовую таблицу краёв")
	around := flag.Int("around", 4, "сколько тактов тела показывать рядом с краями в -v")
	flag.Float64Var(&opts.bpm, "bpm", 0, "темп сетки (по умолчанию из sidecar)")
	flag.Float64Var(&opts.start, "start", -1, "секунда первой доли (по умолчанию из sidecar, иначе 0)")
	flag.IntVar(&opts.intro, "intro", -1, "длина интро в тактах (по умолчанию из Lyrics, иначе 8)")
	flag.IntVar(&opts.outro, "outro", -1, "длина аутро в тактах (по умолчанию из Lyrics, иначе 8)")
	flag.IntVar(&opts.end, "end", -1, "длина конца в тактах (по умолчанию из Lyrics, иначе 2)")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Укажите аудиофайлы. Пример: edge-detect -v *.m4a")
		flag.Usage()
		os.Exit(2)
	}

	var results []result
	for i, path := range files {
		r, err := analyze(path, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[%d/%d] %s — ошибка: %v\n", i+1, len(files), shortName(path), err)
			continue
		}
		fmt.Fprintf(os.Stderr, "[%d/%d] %s\n", i+1, len(files), shortName(path))
		results = append(results, r)
	}
	if len(results) == 0 {
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr)

	for _, r := range results {
		printTrack(r)
		if *verbose {
			printBars(r, *around)
		}
	}
	printSummary(results)
}

// Полный разбор одного файла.
func analyze(path string, opts options) (result, error) {
	meta := readSidecar(path)
	r := result{Path: path, Plan: makePlan(meta.sections(), opts)}
	r.BPM, r.Start, r.GridSource = resolveGrid(meta, opts)
	if r.BPM <= 0 {
		return result{}, fmt.Errorf("темп неизвестен: нет ни sidecar с bpm, ни BPM в промпте — задайте --bpm")
	}

	samples, err := decodeMono(path)
	if err != nil {
		return result{}, err
	}
	bars := groupBars(analyzeFrames(samples), r.BPM, r.Start)
	if len(bars) < r.Plan.Intro+r.Plan.Outro+r.Plan.End+4 {
		return result{}, fmt.Errorf("трек короче плана краёв: %d тактов", len(bars))
	}
	r.Edges = findEdges(bars, r.Plan)
	return r, nil
}

// Сетка тактов: флаги сильнее sidecar, измерение bpm-detect сильнее промпта.
func resolveGrid(meta sidecar, opts options) (bpm, start float64, source string) {
	switch {
	case opts.bpm > 0:
		bpm, source = opts.bpm, "флаг --bpm"
	case meta.BPM > 0:
		bpm, source = meta.BPM, "bpm-detect (sidecar)"
	default:
		bpm, source = meta.declaredBPM(), "темп из промпта, первая доля не измерена"
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
