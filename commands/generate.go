package commands

import (
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Generate() {
	clearOutputDir()
	createOutputDir()

	generateIndexPage()
	copyAssets()
	bundleJS()
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
	data := IndexData{
		FaviconFileName: filepath.Join(Config.assetDir, "gopher_favicon.svg"),
		Title:           "go wasm",
		ScriptFileName:  "bundle.js",
		StyleFileName:   "bundle.css",
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
	log.Print("Bundling JS ...")
	buildCmd := exec.Command("npm", "run", "build")
	buildCmd.Dir = "ts"
	err := buildCmd.Run()
	if err != nil {
		log.Fatal("failed to bundle JS", err)
	}
	srcJSName := filepath.Join("bundle.js")
	destJSName := filepath.Join(Config.outputDir, "bundle.js")
	srcStyleName := filepath.Join("bundle.css")
	destStyleName := filepath.Join(Config.outputDir, "bundle.css")
	copyFile(srcJSName, destJSName)
	copyFile(srcJSName+".map", destJSName+".map")
	copyFile(srcStyleName, destStyleName)
}

func copyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal("failed to open srcFile", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		log.Fatal("failed to creat destFile", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		log.Fatal("failed to copy file", err)
	}

	err = destFile.Sync()
	if err != nil {
		log.Fatal("failed to sync destFile", err)
	}
}
