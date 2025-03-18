.PHONY: all lint golangci-lint test deps examples middlewares

REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH) lint test examples middlewares

lint:
	go fmt ./...
	go vet ./...

golangci-lint:
	golangci-lint run ./...

test:
	go test -cover -race ./...

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)

examples:
	for f in $(shell ls _examples); do ( \
		cd _examples/$$f && \
		go mod tidy && \
		go vet ./... && \
		go build -o /dev/null \
		); done

middlewares:
	for f in gin echo fiber; do ( \
		cd $$f && \
		go mod tidy && \
		go vet ./... && \
		go test -cover -race \
		); done
