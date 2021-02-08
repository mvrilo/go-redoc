# go-redoc

`go-redoc` is an embeded OpenAPI/Swagger documentation for Go programs using [ReDoc](https://github.com/Redocly/redoc).

## Usage

```go
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
		SpecURL:     addr + "openapi.json",
	})

	r := gin.New()
	r.StaticFile("/openapi.json", "./openapi.json")
	r.Use(ginredoc.New(doc))

	println("Server started at", addr)
	panic(r.Run(":8000"))
}
```
