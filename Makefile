REDOC_PATH=assets/redoc.standalone.js
REDOC_URL=https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js

all: $(REDOC_PATH) statik

$(REDOC_PATH):
	curl -sL -o $(REDOC_PATH) $(REDOC_URL)

statik:
	go run github.com/rakyll/statik -src=./assets

clear:
	rm -rf statik
