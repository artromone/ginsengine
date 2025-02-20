BINARY=ginsengine
ASSETS=assets

build:
	go build -o $(BINARY) ./cmd/$(BINARY)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY).exe ./cmd/$(BINARY)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux ./cmd/$(BINARY)

run:
	go run ./cmd/$(BINARY)

assets:
	cp -r $(ASSETS) $(BINARY)_assets

clean:
	rm -f $(BINARY) $(BINARY).exe $(BINARY)-linux

.PHONY: build build-windows build-linux run clean
