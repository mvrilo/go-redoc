# go-redoc

`go-redoc` is an embedded OpenAPI/Swagger documentation ui for Go using [ReDoc](https://github.com/Redocly/redoc).

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

## With Gin

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

## Usage with Echo

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
