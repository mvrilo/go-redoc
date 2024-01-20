.PHONY: all lint test deps

REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH) lint test

lint:
	go fmt ./...
	go vet ./...
	golangci-lint run ./...

test:
	go test -race ./...

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)
