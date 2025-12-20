package commands

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func Generate() {
	clearOutputDir()
	createOutputDir()

	buildIndexPage()
	copyAssets()
}

type IndexData struct {
	Title         string
	StyleFileName string
}

func buildIndexPage() {
	tmpl := parseIndexTemplate()
	out, err := os.Create(filepath.Join(Config.outputDir, "index.html"))
	if err != nil {
		log.Fatal("failed to create indexFile")
	}
	defer out.Close()
	data := IndexData{
		Title:         "wasm calculator",
		StyleFileName: filepath.Join(Config.assetDir, "index.css"),
	}
	tmpl.Execute(out, data)
}

func createOutputDir() {
	if err := os.MkdirAll(Config.outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory")
	}
}

func clearOutputDir() {
	log.Printf("Clearing /%s ...", Config.outputDir)
	err := os.RemoveAll(Config.outputDir)
	if err != nil {
		log.Fatal("failed to clear output directory")
	}
}

func parseIndexTemplate() *template.Template {
	tmpl, err := template.ParseGlob("templates/index.html")
	if err != nil {
		log.Fatal("failed to parse template")
	}
	return tmpl
}

func copyAssets() {
	os.CopyFS(filepath.Join(Config.outputDir, Config.assetDir), os.DirFS(Config.assetDir))
}
