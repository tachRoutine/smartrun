# SmartRun Makefile

.PHONY: build test clean install run dev lint

# Variables
BINARY_NAME=smartrun
BUILD_DIR=dist
CMD_DIR=cmd/smartrun

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Development mode with live reload
dev:
	@echo "Running in development mode..."
	go run ./$(CMD_DIR)

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Help
help:
	@echo "Available commands:"
	@echo "  build    - Build the binary"
	@echo "  test     - Run tests"  
	@echo "  clean    - Clean build artifacts"
	@echo "  install  - Install binary to /usr/local/bin"
	@echo "  run      - Build and run the application"
	@echo "  dev      - Run in development mode"
	@echo "  lint     - Run linter"
	@echo "  fmt      - Format code"
	@echo "  deps     - Download dependencies"
	@echo "  help     - Show this help message"