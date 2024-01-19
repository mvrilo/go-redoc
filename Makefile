REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH) lint test

lint:
	go fmt ./...
	go vet ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

test:
	go test -race ./...

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)
