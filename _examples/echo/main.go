package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func main() {
	addr := "http://127.0.0.1:8000/"

	doc := redoc.New(redoc.Config{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	})

	r := echo.New()
	r.Use(echoredoc.New(doc))

	println("Documentation served at", addr+"docs")
	panic(r.Start(":8000"))
}
