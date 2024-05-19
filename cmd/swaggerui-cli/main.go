package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/flowchartsman/swaggerui"
)

func main() {
	bind := ""
	filename := ""
	prefix := ""

	flag.StringVar(&bind, "bind", "localhost:8080", "Address to bind to. Default localhost:8080")
	flag.StringVar(&prefix, "prefix", "/", "Path prefix . Defaults to \"/\"")
	flag.StringVar(&filename, "filename", "", "Swagger file to load")
	flag.Parse()

	if len(filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if len(prefix) == 0 {
		prefix = "/"
	}

	log.Printf("Using prefix %s", prefix)

	spec, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	log.Printf("Loaded %s", filename)

	// strip the prefix for the handler func
	stripped := strings.TrimRight(prefix, "/")

	http.Handle(prefix, http.StripPrefix(stripped, swaggerui.Handler(spec)))
	log.Printf("Listening on %s", bind)

	log.Fatal(http.ListenAndServe(bind, nil))
}
