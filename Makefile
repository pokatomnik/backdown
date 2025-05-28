BUILD_DIR := build

.PHONY: build clean windows linux darwin

build:
	@case "$(filter-out $@,$(MAKECMDGOALS))" in \
		windows) $(MAKE) build-windows ;; \
		linux) $(MAKE) build-linux ;; \
		darwin) $(MAKE) build-darwin ;; \
		*) echo "Usage: make build [windows|linux|darwin]" ;; \
	esac

windows linux darwin:
	@:

build-windows:
	@echo "Собираем для Windows x86_64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/backdown.exe .

build-linux:
	@echo "Собираем для Linux x86_64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/backdown .

build-darwin:
	@echo "Собираем для MacOS arm64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/backdown .

clean:
	@echo "Чистим build/..."
	@rm -rf $(BUILD_DIR)
