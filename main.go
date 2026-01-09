package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(Page()))

	http.ListenAndServe(":8080", nil)

}
