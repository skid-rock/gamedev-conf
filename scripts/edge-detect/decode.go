package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os/exec"
)

// Частота дискретизации для анализа. Хэты и шейкер живут выше 6 кГц,
// поэтому предел нужен выше, чем у key-detect: при 22050 Гц он 11 кГц.
const analysisRate = 22050

// Декодировать файл в моно-поток float32 внешним ffmpeg.
func decodeMono(path string) ([]float32, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return nil, fmt.Errorf("не найден ffmpeg: %w", err)
	}

	cmd := exec.Command("ffmpeg",
		"-v", "error",
		"-i", path,
		"-ac", "1",
		"-ar", fmt.Sprint(analysisRate),
		"-f", "f32le",
		"-",
	)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffmpeg: %w: %s", err, errBuf.String())
	}

	raw := out.Bytes()
	samples := make([]float32, len(raw)/4)
	for i := range samples {
		v := math.Float32frombits(binary.LittleEndian.Uint32(raw[i*4:]))
		if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
			v = 0
		}
		samples[i] = v
	}
	if len(samples) < frameSize {
		return nil, fmt.Errorf("слишком короткий или пустой звук: %d отсчётов", len(samples))
	}
	return samples, nil
}
