package iris

import (
	"github.com/mvrilo/go-redoc"
	"github.com/mvrilo/go-redoc/iris"
	irisdoc "github.com/mvrilo/go-redoc/iris"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := iris.New()
	r.Use(irisdoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
