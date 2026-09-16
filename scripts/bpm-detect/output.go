package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Темп в виде строки: целое, если дробная часть не значима.
func formatBPM(bpm float64) string {
	if math.Abs(bpm-math.Round(bpm)) < 0.005 {
		return fmt.Sprintf("%.0f", bpm)
	}
	return strings.TrimRight(fmt.Sprintf("%.2f", bpm), "0")
}

// Короткое имя трека для таблицы.
func shortName(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// Оценка сетки словами. Решает не разброс темпа по окнам, а накопленный
// уход сетки: полдоли за трек — это уже слышный промах в конце.
func verdict(a analysis) string {
	switch {
	case a.Confidence < 1.6:
		return "сетки не видно"
	case a.Drift.Measured && a.Drift.Slip > 0.5:
		return "уедет, нужна карта"
	case a.Drift.Measured && a.Drift.Slip > 0.2:
		return "поплывёт к концу"
	case a.Confidence < 2.2:
		return "слабовато"
	default:
		return "ровно"
	}
}

// Таблица в стандартный вывод.
func printTable(results []analysis, verbose bool) {
	fmt.Printf("%-24s %8s %9s %8s %7s %7s %6s  %s\n",
		"ТРЕК", "BPM", "1-Я ДОЛЯ", "ДРЕЙФ", "УХОД", "УВЕР", "ТАКТ", "ОЦЕНКА")
	for _, a := range results {
		driftText, slipText := "—", "—"
		if a.Drift.Measured {
			driftText = fmt.Sprintf("±%.2f", a.Drift.Spread/2)
			slipText = fmt.Sprintf("%.2f", a.Drift.Slip)
		}
		fmt.Printf("%-24s %8s %9.3f %8s %7s %7.2f %6.2f  %s\n",
			shortName(a.Path), formatBPM(a.Fit.BPM), a.FirstBeat,
			driftText, slipText, a.Confidence, a.Downbeat, verdict(a))
		if verbose {
			fmt.Printf("    опора: %s", a.Source)
			if a.Hint > 0 {
				fmt.Printf(" (%.0f)", a.Hint)
			}
			if a.TagBPM > 0 {
				fmt.Printf(", тег файла: %.3f", a.TagBPM)
			}
			fmt.Printf(", доля сетки №%d\n", a.BeatOffset+1)
			for _, r := range a.Rivals {
				fmt.Printf("    кандидат %8s  счёт %.4f\n", formatBPM(r.BPM), r.Score)
			}
			if a.Drift.Measured {
				fmt.Printf("    по окнам: %s\n", driftLine(a.Drift))
				fmt.Printf("    сдвиг сетки по окнам, доли: %s\n", slipLine(a.Drift))
			}
		}
	}
}

// Темп по окнам одной строкой; ненадёжные окна в скобках.
func driftLine(d drift) string {
	var parts []string
	for _, w := range d.Windows {
		if w.Reliable {
			parts = append(parts, formatBPM(w.BPM))
		} else {
			parts = append(parts, "("+formatBPM(w.BPM)+")")
		}
	}
	return strings.Join(parts, " ")
}

// Сдвиг сетки по окнам одной строкой: видно, ползёт ли он в одну сторону.
func slipLine(d drift) string {
	var parts []string
	base := 0.0
	for i, w := range d.Windows {
		if i == 0 {
			base = w.Phase
		}
		parts = append(parts, fmt.Sprintf("%+.2f", w.Phase-base))
	}
	return strings.Join(parts, " ")
}

// Таблица в markdown — для заметок проекта.
func writeMarkdown(path string, results []analysis) error {
	var b strings.Builder
	b.WriteString("# Темп треков\n\n")
	fmt.Fprintf(&b, "Измерено %s утилитой `bpm-detect`.\n\n", time.Now().Format("02.01.2006"))
	b.WriteString("| Трек | BPM | 1-я доля, с | Дрейф | Уход, доли | Увер. | Такт | Опора | Тег файла | Оценка |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|---|\n")
	for _, a := range results {
		driftText, slipText := "—", "—"
		if a.Drift.Measured {
			driftText = fmt.Sprintf("±%.2f", a.Drift.Spread/2)
			slipText = fmt.Sprintf("%.2f", a.Drift.Slip)
		}
		hint := "—"
		if a.Hint > 0 {
			hint = fmt.Sprintf("%s (%s)", formatBPM(a.Hint), a.Source)
		}
		tag := "—"
		if a.TagBPM > 0 {
			tag = fmt.Sprintf("%.3f", a.TagBPM)
		}
		fmt.Fprintf(&b, "| %s | %s | %.3f | %s | %s | %.2f | %.2f | %s | %s | %s |\n",
			shortName(a.Path), formatBPM(a.Fit.BPM), a.FirstBeat, driftText, slipText,
			a.Confidence, a.Downbeat, hint, tag, verdict(a))
	}
	b.WriteString("\nУход — на сколько долей разъедется единая сетка от начала трека к концу;")
	b.WriteString(" выше 0.5 доли одной цифрой темпа не обойтись.")
	b.WriteString(" Уверенность — во сколько раз отклик на найденной сетке выше типичного;")
	b.WriteString(" ниже 1.6 доверять нечему. Такт — запас «единицы» над соперницей, от 0 до 1;")
	b.WriteString(" около нуля означает, что сильная доля выбрана наугад.\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

// Хвост вида « - 174.5» в конце имени.
var nameBPMSuffix = regexp.MustCompile(` - \d{2,3}(\.\d+)?$`)

// Дописать темп в имя файла: GridFollow читает его из имени тейка.
// Sidecar переезжает следом, иначе пара распадётся.
func renameWithBPM(a analysis) (string, error) {
	dir := filepath.Dir(a.Path)
	ext := filepath.Ext(a.Path)
	base := strings.TrimSuffix(filepath.Base(a.Path), ext)
	base = nameBPMSuffix.ReplaceAllString(base, "")

	newBase := fmt.Sprintf("%s - %s", base, formatBPM(a.Fit.BPM))
	newPath := filepath.Join(dir, newBase+ext)
	if newPath == a.Path {
		return a.Path, nil
	}
	if _, err := os.Stat(newPath); err == nil {
		return "", fmt.Errorf("файл уже есть: %s", newPath)
	}
	if err := os.Rename(a.Path, newPath); err != nil {
		return "", err
	}
	if old := sidecarPath(a.Path); fileExists(old) {
		_ = os.Rename(old, sidecarPath(newPath))
	}
	return newPath, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Поля, которые утилита пишет в sidecar.
var sidecarKeys = []string{"bpm", "first_beat", "beat_phase", "bpm_drift", "bpm_confidence", "bpm_source", "bpm_measured_at"}

// Дописать результат в sidecar-json, сохранив порядок и состав чужих полей.
// Поэтому правится текст, а не разобранная карта: у карты порядок ключей теряется.
func updateSidecar(a analysis) (string, error) {
	path := sidecarPath(a.Path)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var kept []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		drop := false
		for _, key := range sidecarKeys {
			if strings.HasPrefix(trimmed, `"`+key+`":`) {
				drop = true
				break
			}
		}
		if !drop {
			kept = append(kept, line)
		}
	}

	if len(kept) == 0 || strings.TrimSpace(kept[len(kept)-1]) != "}" {
		return "", fmt.Errorf("неожиданный формат sidecar: %s", path)
	}

	body := kept[:len(kept)-1]
	if len(body) > 0 {
		last := len(body) - 1
		if !strings.HasSuffix(strings.TrimSpace(body[last]), ",") {
			body[last] += ","
		}
	}

	driftText := "null"
	if a.Drift.Measured {
		driftText = fmt.Sprintf("%.3f", a.Drift.Spread)
	}
	added := []string{
		fmt.Sprintf(`  "bpm": %.3f,`, a.Fit.BPM),
		fmt.Sprintf(`  "first_beat": %.3f,`, a.FirstBeat),
		fmt.Sprintf(`  "beat_phase": %.3f,`, a.Fit.Phase),
		fmt.Sprintf(`  "bpm_drift": %s,`, driftText),
		fmt.Sprintf(`  "bpm_confidence": %.2f,`, a.Confidence),
		fmt.Sprintf(`  "bpm_source": %q,`, a.Source),
		fmt.Sprintf(`  "bpm_measured_at": %q`, time.Now().Format(time.RFC3339)),
	}

	out := append(body, added...)
	out = append(out, "}")
	return path, os.WriteFile(path, []byte(strings.Join(out, "\n")+"\n"), 0o644)
}
