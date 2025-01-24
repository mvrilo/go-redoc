# go-redoc

[![GoDoc](https://godoc.org/github.com/mvrilo/go-redoc?status.svg)](https://godoc.org/github.com/mvrilo/go-redoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/go-redoc?_=1)](https://goreportcard.com/report/github.com/mvrilo/go-redoc?_=1)

`go-redoc` is an embedded OpenAPI documentation ui for Go using [ReDoc](https://github.com/ReDocly/redoc) and Go's [1.16+'s embed package](https://golang.org/pkg/embed/), with middleware implementations for: `net/http`, `gin`, `fiber`, and `echo`.

The template is based on the ReDoc's [bundle template](https://github.com/ReDocly/redoc/blob/master/cli/template.hbs) with the script already placed in the html instead of depending on a CDN.

This package does not generate openapi spec file. Check [this example](_examples/gen) for using code generation with swag.

## Usage

```go
import "github.com/mvrilo/go-redoc"

...

doc := redoc.Redoc{
    Title:       "Example API",
    Description: "Example API Description",
    SpecFile:    "./openapi.json", // "./openapi.yaml"
    SpecPath:    "/openapi.json",  // "/openapi.yaml"
    DocsPath:    "/docs",
}
```

- `net/http`

```go
import (
	"net/http"
	"github.com/mvrilo/go-redoc"
)

...

http.ListenAndServe(address, doc.Handler())
```

- `gin`

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

...

r := gin.New()
r.Use(ginredoc.New(doc))
```

- `echo`

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

...

r := echo.New()
r.Use(echoredoc.New(doc))
```

- `fiber`

```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/mvrilo/go-redoc"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
)

...

r := fiber.New()
r.Use(fiberredoc.New(doc))
```

See [examples](/_examples)


## Configuration Options


```go
r := redoc.Redoc{
		SpecFile: "testdata/spec.json",
		SpecFS:   &spec,
		SpecPath: "/openapi.json", // "/openapi.yaml"
		Title:    "Test API",
		Description: "Meta Description"
        Options: `{` +
        `"disableSearch": true,` +
        `"theme": { ` +
        `"colors": {"primary":{"main":"#297b21"}},` +
        `"headings": {"font_weight":"800"},` +
        `"sidebar": { "backgroundColor": "#cae6c6" }` +
        `}` +
        `}`,
	}
```

`Title` : The head title of your html page - Shown on search engine.

`Description` : The head meta description of your html page - Shown on search engine.

`Options`: redoc option see [Redoc Configuration Documentation](https://github.com/Redocly/redoc/blob/main/docs/config.md)

`SpecFile`: file path to your openapi/swagger file from your project.

`SpecPath`: url path to call your openapi/swagger file from redoc documentation. Must be aligned with your web server configuration.

`DocsPath` : url path to call your generated API documentation. Must be aligned with your web server configuration.