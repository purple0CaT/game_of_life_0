# Variables
VERSION ?= v0.1        # Default version if not provided
MAIN_FILE = main.go   # Default main file if not provided

# Directories
SOURCE_DIR := $(VERSION)
BIN_DIR := bin

# Build command
build:
	@echo "Building $(VERSION)..."
	@cd $(SOURCE_DIR) && go build -o ../$(BIN_DIR)/$(VERSION) $(MAIN_FILE)


# Run command
run:
	@echo "Running $(VERSION)..."
	@cd $(BIN_DIR) && ./$(VERSION)

# Build and run command
build_and_run: build run

# Clean command
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)

.PHONY: build run build_and_run clean
