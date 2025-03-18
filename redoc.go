package redoc

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"text/template"

	_ "embed"
)

// ErrSpecNotFound error for when spec file not found
var ErrSpecNotFound = errors.New("spec not found")

// Redoc configuration
type Redoc struct {
	DocsPath string
	SpecPath string

	SpecFile string
	SpecFS   fs.FS

	Title       string
	Description string
}

// HTML represents the redoc index.html page
//
//go:embed assets/index.html
var HTML string

// JavaScript represents the redoc standalone javascript
//
//go:embed assets/redoc.standalone.js
var JavaScript string

// Body returns the final html with the js in the body
func (r Redoc) Body() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("redoc").Parse(HTML)
	if err != nil {
		return nil, err
	}

	if err = tpl.Execute(buf, map[string]string{
		"body":        JavaScript,
		"title":       r.Title,
		"url":         r.SpecPath,
		"description": r.Description,
	}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r Redoc) getDocAndSpec() ([]byte, []byte) {
	if r.SpecFile == "" {
		panic(ErrSpecNotFound)
	}

	if r.SpecPath == "" {
		r.SpecPath = "/openapi.json"
	}

	var file fs.File
	var err error

	if r.SpecFS != nil {
		file, err = r.SpecFS.Open(r.SpecFile)
	} else {
		file, err = os.Open(r.SpecFile)
	}
	if err != nil {
		panic(err)
	}

	spec, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data, err := r.Body()
	if err != nil {
		panic(err)
	}

	return data, spec
}

// Handler sets some defaults and returns a HandlerFunc
func (r Redoc) Handler() http.HandlerFunc {
	doc, spec := r.getDocAndSpec()

	return func(w http.ResponseWriter, req *http.Request) {
		method := strings.ToLower(req.Method)
		if method != "get" && method != "head" {
			return
		}

		header := w.Header()
		if strings.HasSuffix(req.URL.Path, r.SpecPath) {
			header.Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(spec)
			return
		}

		if r.DocsPath == "" || r.DocsPath == req.URL.Path {
			header.Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(doc)
		}
	}
}
