REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH)
	go fmt ./...
	go vet ./...
	go test ./...

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)
