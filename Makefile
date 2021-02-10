REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH)
	go1.16rc1 fmt ./...
	go1.16rc1 vet ./...
	go1.16rc1 test ./...

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)
