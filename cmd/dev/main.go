package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/examples/starter"
)

func main() {
	mux := http.NewServeMux()

	// Starter page (root)
	mux.Handle("/", templ.Handler(starter.Page()))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
