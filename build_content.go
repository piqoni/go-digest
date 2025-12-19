package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Load markdown files
	files, err := filepath.Glob("pre-docs/*.md")
	if err != nil {
		return err
	}

	digests := make(map[string]string, len(files))
	for _, path := range files {
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		html := markdown.ToHTML(b, nil, nil)
		date := strings.TrimSuffix(filepath.Base(path), ".md")
		digests[date] = string(html)
	}

	// JSON encode digests
	data, err := json.MarshalIndent(digests, "    ", "  ")
	if err != nil {
		return err
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		return err
	}

	out, err := os.Create("pre-docs/index.html")
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, map[string]string{
		"DigestsJSON": string(data),
	})
}
