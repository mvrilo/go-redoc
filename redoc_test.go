package redoc_test

import (
	"embed"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mvrilo/go-redoc"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/spec.json
var spec embed.FS

func TestRedoc(t *testing.T) {
	r := redoc.Redoc{
		SpecFile: "testdata/spec.json",
		SpecFS:   &spec,
		SpecPath: "/openapi.json", // "/openapi.yaml"
		Title:    "Test API",
	}

	t.Run("Body", func(t *testing.T) {
		body, err := r.Body()
		assert.NoError(t, err)
		assert.Contains(t, string(body), r.Title)
	})

	t.Run("Handler", func(t *testing.T) {
		handler := r.Handler()

		t.Run("Spec", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(body), `"swagger":"2.0"`)
		})

		t.Run("Docs", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))

			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(body), r.Title)
		})
	})
}
