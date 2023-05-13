# Variables
BIN_NAME := myapp
BUILD_DIR := build
PLATFORMS := windows linux darwin
BUILD_FLAGS := -ldflags="-s -w"  # to create smaller binaries

# Targets
.PHONY: all clean $(PLATFORMS)

all: $(PLATFORMS)

$(PLATFORMS):
	@echo "Building $(BIN_NAME) for $@"
	@mkdir -p $(BUILD_DIR)
	@GOOS=$@ GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BIN_NAME)_$@
	@echo "Build complete: $(BUILD_DIR)/$(BIN_NAME)_$@"

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete"

