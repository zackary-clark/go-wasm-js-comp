package commands

import (
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Generate() {
	clearOutputDir()
	createOutputDir()

	bundleJS()
	copyAssets()
	generateIndexPage()
}

type IndexData struct {
	FaviconFileName string
	Title           string
	ScriptFileName  string
	StyleFileName   string
}

func generateIndexPage() {
	tmpl := parseIndexTemplate()
	out, err := os.Create(filepath.Join(Config.outputDir, "index.html"))
	if err != nil {
		log.Fatal("failed to create indexFile")
	}
	defer out.Close()
	dir, err := os.ReadDir(Config.outputDir)
	if err != nil {
		log.Fatal("failed to read outputDir")
	}
	scriptFileName := ""
	styleFileName := ""
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".css") {
			styleFileName = file.Name()
		}
		if strings.HasSuffix(file.Name(), ".js") {
			scriptFileName = file.Name()
		}
	}
	data := IndexData{
		FaviconFileName: filepath.Join(Config.assetDir, "gopher_favicon.svg"),
		Title:           "go wasm",
		ScriptFileName:  scriptFileName,
		StyleFileName:   styleFileName,
	}
	err = tmpl.Execute(out, data)
	if err != nil {
		log.Fatal("failed to execute tmpl", err)
	}
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

func bundleJS() {
	buildCmd := exec.Command("npm", "run", "build")
	buildCmd.Dir = "ts"
	buildCmd.Env = append(buildCmd.Environ(), "BUILD_OUTPUT_DIR="+Config.outputDir)
	err := buildCmd.Run()
	if err != nil {
		log.Fatal("failed to bundle JS", err)
	}
}
