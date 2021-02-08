package redoc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mvrilo/go-redoc/statik"
	"github.com/rakyll/statik/fs"
)

const indexTemplate = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{ .title }}</title>
  <meta name="description" content="{{ .description }}">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
  body {
	margin: 0;
	padding: 0;
  }
  </style>
</head>
<body>
  <div id="main"></div>
  <script>{{ .body }}</script>
  <script>Redoc.init("{{ .url }}", {}, document.getElementById("main"))</script>
</body>
</html>`

type Config struct {
	Path        string
	SpecURL     string
	Title       string
	Description string
}

type Redoc struct {
	fs     http.FileSystem
	Config *Config
}

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

func (r *Redoc) Body() ([]byte, error) {
	body, err := r.fs.Open("/redoc.standalone.js")
	if err != nil {
		return nil, err
	}
	defer body.Close()

	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("redoc").Parse(indexTemplate)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(buf, map[string]string{
		"body":        string(data),
		"title":       r.Config.Title,
		"url":         r.Config.SpecURL,
		"description": r.Config.Description,
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r *Redoc) Handler() http.HandlerFunc {
	data, err := r.Body()
	if err != nil {
		panic(err)
	}

	path := r.Config.Path
	if path == "" {
		path = "/"
	}

	return func(w http.ResponseWriter, req *http.Request) {
		method := strings.ToLower(req.Method)
		if req.URL.Path == path && (method == "get" || method == "head") {
			w.WriteHeader(200)
			w.Write(data)
		}
	}
}
