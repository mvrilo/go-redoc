package ginredoc_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

const openapimock = `{"openapi":"3.0.2","info":{"title":"Swagger Petstore - OpenAPI 3.0","description":"This is a sample Pet Store Server.","termsOfService":"http://swagger.io/terms/","contact":{"email":"apiteam@swagger.io"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"version":"1.0.5"}}`

func TestNew(t *testing.T) {
	fs := fstest.MapFS{
		"openapi.json": {
			Data: []byte(openapimock),
		},
	}

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFS:      fs,
		SpecFile:    "openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/before", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})
	r.Use(ginredoc.New(doc))
	r.GET("/after", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok after")
	})

	tests := []struct {
		path        string
		status      int
		contentType string
	}{
		{
			path:   "/before",
			status: http.StatusNoContent,
		},
		{
			path:   "/after",
			status: http.StatusOK,
		},
		{
			path:   "/notfound",
			status: http.StatusNotFound,
		},
		{
			path:        "/docs",
			status:      http.StatusOK,
			contentType: "text/html",
		},
		{
			path:        "/openapi.json",
			status:      http.StatusOK,
			contentType: "application/json",
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tt.path, nil)
		r.ServeHTTP(w, req)

		if tt.status != w.Code {
			t.Errorf(fmt.Sprintf("Got status %d, expected %d", w.Code, tt.status))
		}

		contentType := strings.ToLower(w.Result().Header.Get("content-type"))
		if tt.contentType != "" && !strings.Contains(tt.contentType, contentType) {
			t.Errorf(fmt.Sprintf("Got header content-type %s, expected %s", contentType, tt.contentType))
		}
	}
}
