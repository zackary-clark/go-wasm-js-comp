package commands

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	log.Printf("serve from port: %d\nEventually Anyway...\n", Config.port)

	fs := http.FileServer(http.Dir(Config.outputDir))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(fmt.Sprintf(":%d", Config.port), nil)
}
