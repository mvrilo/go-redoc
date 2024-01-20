package redoc

import (
	"bytes"
	"embed"
	"errors"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// ErrSpecNotFound error for when spec file not found
var ErrSpecNotFound = errors.New("spec not found")

// Redoc configuration
type Redoc struct {
	DocsPath    string
	SpecPath    string
	SpecFile    string
	SpecFS      *embed.FS
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

// Handler sets some defaults and returns a HandlerFunc
func (r Redoc) Handler() http.HandlerFunc {
	data, err := r.Body()
	if err != nil {
		panic(err)
	}

	specFile := r.SpecFile
	if specFile == "" {
		panic(ErrSpecNotFound)
	}

	if r.SpecPath == "" {
		r.SpecPath = "/openapi.json"
	}

	var spec []byte
	if r.SpecFS == nil {
		spec, err = os.ReadFile(specFile)
		if err != nil {
			panic(err)
		}
	} else {
		spec, err = r.SpecFS.ReadFile(specFile)
		if err != nil {
			panic(err)
		}
	}

	docsPath := r.DocsPath
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

		if docsPath == "" || docsPath == req.URL.Path {
			header.Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}
	}
}
