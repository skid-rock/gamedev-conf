package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Темп, названный в промпте: «instrumental 172 BPM two-step drive…».
var bpmInText = regexp.MustCompile(`(?i)(\d{2,3}(?:\.\d+)?)\s*bpm`)

// Путь к sidecar-json рядом с аудиофайлом.
func sidecarPath(audioPath string) string {
	ext := filepath.Ext(audioPath)
	return strings.TrimSuffix(audioPath, ext) + ".json"
}

// Заявленный темп из sidecar: сначала смотрим style, потом prompt.
// Это не измерение, а намерение — Suno держит его лишь примерно,
// зато метрика («считаем по 172, а не по 86») из него видна точно.
func declaredBPM(audioPath string) float64 {
	data, err := os.ReadFile(sidecarPath(audioPath))
	if err != nil {
		return 0
	}
	var doc map[string]any
	if json.Unmarshal(data, &doc) != nil {
		return 0
	}
	for _, key := range []string{"style", "prompt", "display_tags", "genre"} {
		text, _ := doc[key].(string)
		if m := bpmInText.FindStringSubmatch(text); m != nil {
			if v, err := strconv.ParseFloat(m[1], 64); err == nil {
				return v
			}
		}
	}
	return 0
}

// Темп из тега самого файла. Suno пишет туда результат своего анализатора:
// дробная часть выглядит осмысленно, а вот метрика чаще всего чужая —
// встречаются и две трети реального темпа, и половина.
func taggedBPM(audioPath string) float64 {
	if _, err := exec.LookPath("ffprobe"); err != nil {
		return 0
	}
	out, err := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format_tags=BPM",
		"-of", "default=nw=1:nk=1",
		audioPath,
	).Output()
	if err != nil {
		return 0
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		return 0
	}
	return v
}

// Опорные темпы для поиска, в порядке доверия к метрике.
// Первый элемент задаёт ту самую метрику, к которой притягивается ответ.
func collectHints(audioPath string, flagHint, min, max float64) ([]float64, float64, float64, string) {
	declared := declaredBPM(audioPath)
	tag := taggedBPM(audioPath)

	var hints []float64
	source := ""

	switch {
	case flagHint > 0:
		hints = append(hints, flagHint)
		source = "флаг --hint"
	case declared > 0:
		hints = append(hints, declared)
		source = "промпт (sidecar)"
	}

	// Заявленный темп полезен как центр поиска, даже если метрику задал флаг.
	if declared > 0 && (len(hints) == 0 || hints[0] != declared) {
		hints = append(hints, declared)
	}

	// Тег и его кратности — кандидаты по значению, а не по метрике.
	if tag > 0 {
		for _, f := range append([]float64{1}, octaveFactors...) {
			v := tag * f
			if v >= min && v <= max {
				hints = append(hints, v)
			}
		}
		if source == "" {
			source = "тег файла"
		}
	}

	if source == "" {
		source = "только звук"
	}
	ref := 0.0
	if len(hints) > 0 {
		ref = hints[0]
	}
	return hints, ref, tag, source
}
