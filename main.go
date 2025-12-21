package main

import (
	"log"
	"os"

	"github.com/zackary-clark/go-wasm-js-comp/commands"
)

func main() {
	cmd := "serve"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "generate":
		commands.Generate()
	case "serve":
		commands.Serve()
	default:
		log.Fatal("Invalid command!")
	}
}
