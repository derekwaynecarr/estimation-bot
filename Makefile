.PHONY: build run clean test

# Default target
all: build

# Build the binary
build:
	go build -o estimation-bot .

# Run with example (you'll need to provide actual URLs and API key)
run:
	@echo "Example usage:"
	@echo "./estimation-bot -urls='https://github.com/owner/repo/blob/main/design.md' -api-key='your-api-key'"
	@echo "Or set GEMINI_API_KEY environment variable and run:"
	@echo "./estimation-bot -urls='https://github.com/owner/repo/blob/main/design.md'"

# Clean build artifacts
clean:
	rm -f estimation-bot

# Test the application
test:
	go test ./...

# Download dependencies
deps:
	go mod tidy
	go mod download

# Install the binary to $GOPATH/bin
install:
	go install .

# Show help
help:
	@echo "Available targets:"
	@echo "  build   - Build the estimation-bot binary"
	@echo "  run     - Show example usage"
	@echo "  clean   - Remove build artifacts"
	@echo "  test    - Run tests"
	@echo "  deps    - Download and tidy dependencies"
	@echo "  install - Install binary to GOPATH/bin"
	@echo "  help    - Show this help message" 