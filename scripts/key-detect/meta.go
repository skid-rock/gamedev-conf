package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Темп, названный в промпте: «… 88 BPM».
var bpmInText = regexp.MustCompile(`(?i)(\d{2,3}(?:\.\d+)?)\s*bpm`)

// Поля sidecar-json, которые нужны утилите.
// bpm и first_beat появляются, если по треку уже прошёл bpm-detect --sidecar.
type sidecar struct {
	Style     string  `json:"style"`
	Prompt    string  `json:"prompt"`
	BPM       float64 `json:"bpm"`
	FirstBeat float64 `json:"first_beat"`
	HasFirst  bool    `json:"-"`
}

// Путь к sidecar-json рядом с аудиофайлом.
func sidecarPath(audioPath string) string {
	ext := filepath.Ext(audioPath)
	return strings.TrimSuffix(audioPath, ext) + ".json"
}

// Прочитать sidecar; если его нет или он битый — вернуть пустой.
func readSidecar(audioPath string) sidecar {
	data, err := os.ReadFile(sidecarPath(audioPath))
	if err != nil {
		return sidecar{}
	}
	var s sidecar
	if json.Unmarshal(data, &s) != nil {
		return sidecar{}
	}
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) == nil {
		_, s.HasFirst = raw["first_beat"]
	}
	return s
}

// Тональность из промпта: сначала style, потом разметка prompt.
func (s sidecar) declaredKey() (keyCandidate, bool) {
	for _, text := range []string{s.Style, s.Prompt} {
		if k, ok := parseKey(text); ok {
			return k, true
		}
	}
	return keyCandidate{}, false
}

// Темп из промпта — запасной вариант, когда bpm-detect ещё не запускали.
func (s sidecar) declaredBPM() float64 {
	for _, text := range []string{s.Style, s.Prompt} {
		if m := bpmInText.FindStringSubmatch(text); m != nil {
			if v, err := strconv.ParseFloat(m[1], 64); err == nil {
				return v
			}
		}
	}
	return 0
}
