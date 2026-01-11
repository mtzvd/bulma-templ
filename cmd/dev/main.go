package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/examples/kitchensink"
	"github.com/mtzvd/bulma-templ/examples/starter"
)

func main() {
	mux := http.NewServeMux()

	// Starter page (root)
	mux.Handle("/", templ.Handler(starter.Page()))

	// Kitchen Sink page
	mux.Handle("/kitchensink", templ.Handler(kitchensink.Page()))

	log.Println("Listening on http://localhost:8080")
	log.Println("  - Starter:      http://localhost:8080/")
	log.Println("  - Kitchen Sink: http://localhost:8080/kitchensink")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
