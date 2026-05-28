# AGENTS.md

## Repository

Go library that embeds a [ReDoc](https://github.com/ReDocly/redoc) UI for OpenAPI/Swagger specs. Provides a core `net/http` handler plus middleware adapters for gin, echo, fiber, and iris.

## Multi-module structure

This is a **multi-module repo**. Each framework adapter is a separate Go module with its own `go.mod`:

```
go.mod                  # github.com/mvrilo/go-redoc  (core, go 1.18)
echo/go.mod             # github.com/mvrilo/go-redoc/echo
fiber/go.mod            # github.com/mvrilo/go-redoc/fiber
gin/go.mod              # github.com/mvrilo/go-redoc/gin
iris/go.mod             # github.com/mvrilo/go-redoc/iris
_examples/*/go.mod      # standalone example apps
```

**`go test ./...` from root only tests the core module.** Adapter modules must be tested from their own directories.

The echo, fiber, and gin modules use `replace github.com/mvrilo/go-redoc => ../` for local development. The iris module does **not** — it references the published `v0.1.5`.

## Commands

```sh
make test          # go test -race ./...  (core module only)
make lint          # go fmt + go vet + golangci-lint (no .golangci.yml — uses defaults)
make deps          # installs golangci-lint
make all           # downloads redoc JS from CDN, then lint + test
make assets/redoc.standalone.js   # curl the bundled ReDoc v2.5.1 JS
```

To test an adapter module:

```sh
cd echo && go test -race ./...
cd gin  && go test -race ./...
# etc.
```

## Architecture

- **`redoc.go`** — the entire core: `Redoc` struct, `Body()` (renders HTML via `text/template`), `Handler()` (serves spec + docs as `http.HandlerFunc`)
- **`assets/index.html`** — Go template; loads ReDoc JS from CDN (`cdn.redoc.ly/redoc/v2.5.1`)
- **`assets/redoc.standalone.js`** — embedded via `//go:embed` into the `JavaScript` var but **not used at runtime** (the HTML template uses the CDN). Inflates binary size (~890 KB) for no benefit.
- **Adapter packages** (`echo/`, `fiber/`, `gin/`, `iris/`) — each ~15 lines, wrapping `Handler()` into framework-specific middleware.
- **`Handler()` panics** on setup errors (spec not found, template render failure). It does not return an error.

## Conventions

- Package naming: `echoredoc`, `fiberredoc`, `ginredoc` — except `iris/` which uses package name `iris` (inconsistent).
- Test framework: `stretchr/testify/assert` with `httptest`.
- Test data: `testdata/spec.json` (Swagger 2.0 Petstore).
- Commit style: loose conventional commits (`feat:`, `fix:`, `chore:`, `refactor:`), lowercase after prefix.

## CI

GitHub Actions on push/PR to `master`: lint job + test job, Go matrix `[1.17, 1.21]`. The `go.mod` minimum is 1.18 — the 1.17 matrix entry is stale.
