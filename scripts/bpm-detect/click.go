package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Частота клик-дорожки: совпадает с обычной для mp3, лишней пересборки нет.
const clickRate = 44100

// Собрать отрывок с щелчками по найденной сетке.
// Цифре можно верить или не верить, а вот попадание щелчка в удар слышно сразу.
func writeClickPreview(a analysis, from, length float64, outDir string) (string, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return "", fmt.Errorf("не найден ffmpeg: %w", err)
	}

	wav, err := renderClickTrack(a)
	if err != nil {
		return "", err
	}
	tmp, err := os.CreateTemp("", "bpm-click-*.wav")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write(wav); err != nil {
		tmp.Close()
		return "", err
	}
	tmp.Close()

	base := strings.TrimSuffix(filepath.Base(a.Path), filepath.Ext(a.Path))
	if outDir == "" {
		outDir = filepath.Dir(a.Path)
	}
	out := filepath.Join(outDir, fmt.Sprintf("%s - click %s.mp3", base, formatBPM(a.Fit.BPM)))

	args := []string{
		"-v", "error", "-y",
		"-i", a.Path,
		"-i", tmp.Name(),
		"-filter_complex", "[0:a]volume=0.55[m];[1:a]volume=1.0[c];[m][c]amix=inputs=2:duration=shortest:normalize=0[a]",
		"-map", "[a]",
	}
	if from > 0 {
		args = append(args, "-ss", fmt.Sprintf("%.3f", from))
	}
	if length > 0 {
		args = append(args, "-t", fmt.Sprintf("%.3f", length))
	}
	args = append(args, "-c:a", "libmp3lame", "-q:a", "4", out)

	var errBuf bytes.Buffer
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffmpeg: %w: %s", err, errBuf.String())
	}
	return out, nil
}

// Клик-дорожка на всю длину трека: «единица» звонче и выше остальных долей.
func renderClickTrack(a analysis) ([]byte, error) {
	total := int(a.Duration*clickRate) + clickRate
	samples := make([]int16, total)

	period := a.Fit.period()
	beat := 0
	for t := a.FirstBeat; t < a.Duration; t += period {
		strong := beat%4 == 0
		freq, amp := 900.0, 0.35
		if strong {
			freq, amp = 1600, 0.6
		}
		addClick(samples, t, freq, amp)
		beat++
	}

	// Щелчки до первой сильной доли: сетка начинается не с нуля, но начало
	// трека тоже надо слышать размеченным.
	beat = 0
	for t := a.FirstBeat - period; t >= 0; t -= period {
		beat++
		strong := beat%4 == 0
		freq, amp := 900.0, 0.35
		if strong {
			freq, amp = 1600, 0.6
		}
		addClick(samples, t, freq, amp)
	}

	return encodeWAV(samples, clickRate)
}

// Короткий затухающий щелчок в заданной точке.
func addClick(samples []int16, at, freq, amp float64) {
	const durationSec = 0.03
	start := int(at * clickRate)
	count := int(durationSec * clickRate)
	for i := 0; i < count; i++ {
		idx := start + i
		if idx < 0 || idx >= len(samples) {
			continue
		}
		x := float64(i) / clickRate
		env := math.Exp(-x / 0.006)
		v := amp * env * math.Sin(2*math.Pi*freq*x)
		sum := float64(samples[idx])/32767 + v
		sum = math.Max(-1, math.Min(1, sum))
		samples[idx] = int16(sum * 32767)
	}
}

// WAV PCM 16 бит моно.
func encodeWAV(samples []int16, rate int) ([]byte, error) {
	var buf bytes.Buffer
	dataLen := len(samples) * 2

	buf.WriteString("RIFF")
	binary.Write(&buf, binary.LittleEndian, uint32(36+dataLen))
	buf.WriteString("WAVE")
	buf.WriteString("fmt ")
	binary.Write(&buf, binary.LittleEndian, uint32(16))
	binary.Write(&buf, binary.LittleEndian, uint16(1))
	binary.Write(&buf, binary.LittleEndian, uint16(1))
	binary.Write(&buf, binary.LittleEndian, uint32(rate))
	binary.Write(&buf, binary.LittleEndian, uint32(rate*2))
	binary.Write(&buf, binary.LittleEndian, uint16(2))
	binary.Write(&buf, binary.LittleEndian, uint16(16))
	buf.WriteString("data")
	binary.Write(&buf, binary.LittleEndian, uint32(dataLen))
	if err := binary.Write(&buf, binary.LittleEndian, samples); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
