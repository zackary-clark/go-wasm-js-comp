package commands

import (
	"flag"
	"fmt"
	"os"
)

var configFlagSet = flag.NewFlagSet("Config Flags", flag.ExitOnError)
var assetDir = configFlagSet.String("assetDir", "assets", "Asset Directory")
var outputDir = configFlagSet.String("outputDir", "out", "Output Directory for static build")
var port = configFlagSet.Int("port", 8080, "Server Port")
var templateDir = configFlagSet.String("templateDir", "templates", "HTML Template Directory")
var verbose = configFlagSet.Bool("verbose", false, "Verbose Output?")

var Config = struct {
	assetDir    string
	outputDir   string
	port        int
	templateDir string
	verbose     bool
}{
	assetDir:    *assetDir,
	outputDir:   *outputDir,
	port:        *port,
	templateDir: *templateDir,
	verbose:     *verbose,
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
