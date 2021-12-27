package main

import (
	"net/http"

	"github.com/mvrilo/go-redoc"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	// mux := http.ServeMux{}
	// handler := doc.Handler()
	// mux.Handle("/docs", handler)
	// mux.Handle("/openapi.json", handler)

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(http.ListenAndServe(":8000", doc.Handler()))
}
