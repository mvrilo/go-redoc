package redoc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mvrilo/go-redoc/statik"
	"github.com/rakyll/statik/fs"
)

// ErrSpecNotFound error for when spec file not found
var ErrSpecNotFound = errors.New("spec not found")

// Config is Redoc configuration for the js and html
type Config struct {
	DocsPath    string
	SpecPath    string
	SpecFile    string
	Title       string
	Description string
}

// Redoc contains configuration and the filesystem for the assets
type Redoc struct {
	fs     http.FileSystem
	Config *Config
}

// New takes a Config and initializes a Redoc
func New(config Config) *Redoc {
	filesystem, err := fs.New()
	if err != nil {
		panic(err)
	}

	return &Redoc{
		fs:     filesystem,
		Config: &config,
	}
}

func (r *Redoc) open(file string) ([]byte, error) {
	f, err := r.fs.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// Body returns the final html with the js in the body
func (r *Redoc) Body() ([]byte, error) {
	html, err := r.open("/index.html")
	if err != nil {
		return nil, err
	}

	js, err := r.open("/redoc.standalone.js")
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("redoc").Parse(string(html))
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(buf, map[string]string{
		"body":        string(js),
		"title":       r.Config.Title,
		"url":         r.Config.SpecPath,
		"description": r.Config.Description,
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Handler sets some defaults and returns a HandlerFunc
func (r *Redoc) Handler() http.HandlerFunc {
	data, err := r.Body()
	if err != nil {
		panic(err)
	}

	specFile := r.Config.SpecFile
	if specFile == "" {
		panic(ErrSpecNotFound)
	}

	specPath := r.Config.SpecPath
	if specPath == "" {
		specPath = "./openapi.json"
	}

	docsPath := r.Config.DocsPath
	if docsPath == "" {
		docsPath = "/"
	}

	spec, err := ioutil.ReadFile(specFile)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		method := strings.ToLower(req.Method)

		if method != "get" && method != "head" {
			return
		}

		switch req.URL.Path {
		case docsPath:
			w.WriteHeader(200)
			w.Header().Set("content-type", "text/html")
			w.Write(data)
		case specPath:
			w.WriteHeader(200)
			w.Header().Set("content-type", "application/json")
			w.Write(spec)
		default:
		}
	}
}
