package commands

import (
	"flag"
	"fmt"
	"os"
)

var configFlagSet = flag.NewFlagSet("Config Flags", flag.ExitOnError)
var assetDir = configFlagSet.String("assetDir", "assets", "Asset Directory")
var javascriptDir = configFlagSet.String("javascriptDir", "ts", "Javascript Directory")
var outputDir = configFlagSet.String("outputDir", "out", "Output Directory for static build")
var port = configFlagSet.Int("port", 8080, "Server Port")
var stylesheetDir = configFlagSet.String("stylesheetDir", "stylesheets", "Stylesheet Directory")
var templateDir = configFlagSet.String("templateDir", "templates", "HTML Template Directory")
var verbose = configFlagSet.Bool("verbose", false, "Verbose Output?")

var Config = struct {
	assetDir      string
	javascriptDir string
	outputDir     string
	port          int
	stylesheetDir string
	templateDir   string
	verbose       bool
}{
	assetDir:      *assetDir,
	javascriptDir: *javascriptDir,
	outputDir:     *outputDir,
	port:          *port,
	stylesheetDir: *stylesheetDir,
	templateDir:   *templateDir,
	verbose:       *verbose,
}

// TODO: use flag sets, and allow toml/env configuration
func init() {
	args := os.Args
	if len(args) > 1 {
		args = os.Args[2:]
	}
	configFlagSet.Parse(args)

	if *verbose {
		fmt.Printf("\nConfig Flags:\n\t%+v\n\n", Config)
	}
}
