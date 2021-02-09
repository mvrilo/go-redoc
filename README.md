# go-redoc

[![GoDoc](https://godoc.org/github.com/mvrilo/go-redoc?status.svg)](https://godoc.org/github.com/mvrilo/go-redoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/go-redoc?_=1)](https://goreportcard.com/report/github.com/mvrilo/go-redoc?_=1)

`go-redoc` is an embedded OpenAPI documentation ui for Go using [ReDoc](https://github.com/Redocly/redoc), with middleware implementations for: `net/http`, `gin` and `echo`.

## Usage

```go
import "github.com/mvrilo/go-redoc"

...

doc := redoc.New(redoc.Config{
    Title:       "Example API",
    Description: "Example API Description",
    SpecFile:    "./openapi.json",
    SpecPath:    "/openapi.json",
    DocsPath:    "/docs",
})
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

See [examples](/_examples)

## Related projects

- https://github.com/go-openapi/runtime/blob/master/middleware/redoc.go
- https://github.com/holdatech/go-redoc-middleware
