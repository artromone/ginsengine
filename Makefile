BINARY=ginsengine
ASSETS=assets
BUILD_DIR=./build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY) ./cmd/$(BINARY)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-windows.exe ./cmd/$(BINARY)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-linux ./cmd/$(BINARY)

run: build
	./$(BUILD_DIR)/$(BINARY)

assets: build
	cp -r $(ASSETS) $(BUILD_DIR)/$(BINARY)_assets

clean:
	rm -rf $(BUILD_DIR)

.PHONY: build build-windows build-linux run clean assets
