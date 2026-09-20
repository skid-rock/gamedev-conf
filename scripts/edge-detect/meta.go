package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Темп, названный в промпте: «… 107 BPM».
var bpmInText = regexp.MustCompile(`(?i)(\d{2,3}(?:\.\d+)?)\s*bpm`)

// Метка секции в Lyrics: «[Intro: garage drums and soft hats, 8 bars]».
var sectionTag = regexp.MustCompile(`\[([^\]]+)\]`)
var barsInTag = regexp.MustCompile(`(?i)(\d+)\s*bars?\b`)

// Поля sidecar-json, которые нужны утилите.
// bpm и first_beat появляются, если по треку уже прошёл bpm-detect --sidecar.
type sidecar struct {
	Style     string  `json:"style"`
	Prompt    string  `json:"prompt"`
	BPM       float64 `json:"bpm"`
	FirstBeat float64 `json:"first_beat"`
	HasFirst  bool    `json:"-"`
}

// Секция из разметки Lyrics.
type section struct {
	Name string
	Bars int
}

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

// Секции с числом тактов из разметки Lyrics. Метки без тактов пропускаются.
func (s sidecar) sections() []section {
	var out []section
	for _, m := range sectionTag.FindAllStringSubmatch(s.Prompt, -1) {
		body := m[1]
		bm := barsInTag.FindStringSubmatch(body)
		if bm == nil {
			continue
		}
		n, err := strconv.Atoi(bm[1])
		if err != nil || n <= 0 {
			continue
		}
		name := body
		if i := strings.IndexAny(body, ":,|"); i >= 0 {
			name = body[:i]
		}
		out = append(out, section{Name: strings.TrimSpace(name), Bars: n})
	}
	return out
}

// План краёв: длины интро, аутро и конца в тактах и общая длина по разметке.
type plan struct {
	Intro, Outro, End int
	Total             int
	// Откуда взяты длины: разметка Lyrics, флаги или значения по умолчанию.
	Source string
}

// Длины по умолчанию — как в файлах событий: интро 8, аутро 8, конец 2.
const defaultIntro, defaultOutro, defaultEnd = 8, 8, 2

func makePlan(secs []section, opts options) plan {
	p := plan{Intro: defaultIntro, Outro: defaultOutro, End: defaultEnd, Source: "по умолчанию"}
	if len(secs) > 0 {
		p.Source = "Lyrics"
		for _, s := range secs {
			p.Total += s.Bars
		}
		if hasPrefixFold(secs[0].Name, "intro") {
			p.Intro = secs[0].Bars
		}
		last := len(secs) - 1
		if hasPrefixFold(secs[last].Name, "end") {
			p.End = secs[last].Bars
			if last > 0 && hasPrefixFold(secs[last-1].Name, "outro") {
				p.Outro = secs[last-1].Bars
			}
		} else if hasPrefixFold(secs[last].Name, "outro") {
			// Секции конца нет: аутро само заканчивает трек.
			p.Outro = secs[last].Bars
			p.End = 0
		}
	}
	if opts.intro >= 0 {
		p.Intro, p.Source = opts.intro, "флаги"
	}
	if opts.outro >= 0 {
		p.Outro, p.Source = opts.outro, "флаги"
	}
	if opts.end >= 0 {
		p.End, p.Source = opts.end, "флаги"
	}
	return p
}

func hasPrefixFold(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}
