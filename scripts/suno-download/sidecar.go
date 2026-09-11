package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type sidecarJSON struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	DisplayName  string `json:"display_name"`
	Handle       string `json:"handle"`
	Genre        string `json:"genre"`
	DisplayTags  string `json:"display_tags"`
	Prompt       string `json:"prompt"`
	Style        string `json:"style"`
	NegativeTags string `json:"negative_tags"`
	Model        string `json:"model"`
	CreatedAt    string `json:"created_at"`
}

func sidecarPath(m4aPath string) string {
	ext := filepath.Ext(m4aPath)
	return strings.TrimSuffix(m4aPath, ext) + ".json"
}

func writeSidecar(m4aPath string, clip clipJSON) (string, error) {
	path := sidecarPath(m4aPath)
	doc := sidecarJSON{
		ID:           clip.ID,
		Title:        clip.Title,
		Author:       clip.DisplayName,
		DisplayName:  clip.DisplayName,
		Handle:       clip.Handle,
		Genre:        clip.DisplayTags,
		DisplayTags:  clip.DisplayTags,
		Prompt:       clip.Metadata.Prompt,
		Style:        clip.Metadata.Tags,
		NegativeTags: clip.Metadata.NegativeTags,
		Model:        clip.modelLabel(),
		CreatedAt:    clip.CreatedAt,
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return "", err
	}
	return path, nil
}
