package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
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

	r := gin.New()
	r.Use(ginredoc.New(doc))

	println("Documentation served at", addr+"docs")
	panic(r.Run(":8000"))
}
