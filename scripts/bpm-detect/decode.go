package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os/exec"
)

// Частота дискретизации, в которую сводится звук перед анализом.
// Для огибающей атак хватает с запасом: интересны удары, а не тембр.
const analysisRate = 11025

// Декодировать файл в моно-поток float32 внешним ffmpeg.
// Своего декодера Opus/AAC в проекте нет, а ffmpeg уже нужен и для клик-трека.
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
		bits := binary.LittleEndian.Uint32(raw[i*4:])
		v := math.Float32frombits(bits)
		if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
			v = 0
		}
		samples[i] = v
	}
	if len(samples) < analysisRate {
		return nil, fmt.Errorf("слишком короткий или пустой звук: %d отсчётов", len(samples))
	}
	return samples, nil
}

// Длительность декодированного потока в секундах.
func durationOf(samples []float32) float64 {
	return float64(len(samples)) / analysisRate
}
