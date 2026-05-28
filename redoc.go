package redoc

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	SpecDir     string
	SpecFS      *embed.FS
	Title       string
	Description string
	Options     map[string]any
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

	var optionsString = "{}"

	var optionsByte, errM = json.Marshal(r.Options)
	if errM == nil {
		optionsString = string(optionsByte)
	} else {
		log.Printf("Invalid json options provided, using default options instead.")
	}

	if err = tpl.Execute(buf, map[string]string{
		"body":        JavaScript,
		"title":       r.Title,
		"url":         r.SpecPath,
		"description": r.Description,
		"options":     optionsString,
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

	if r.SpecDir == "" {
		r.SpecDir = "components"
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
			return
		}

		// load spec files
		ext := filepath.Ext(req.URL.Path)
		if ext == ".yaml" || ext == ".json" {
			header.Set("Content-Type", "application/json")
			p := filepath.Join(r.SpecDir, filepath.FromSlash(req.URL.Path))
			subSpec, err := os.ReadFile(p)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("file not found."))
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(subSpec)
			return
		}

	}
}
