# Makefile for the Anagram Finder project

# Binary output
BINARY=bin/anagram

.PHONY: all build clean test lint

all: build

# Build the project
build:
	@mkdir -p bin
	go build -o $(BINARY) ./cmd/anagram

# Run tests
test:
	go test -v ./...

# Run linters
lint:
	go vet ./...
	golangci-lint run ./...

# Clean build artifacts
clean:
	@rm -rf $(BINARY)
	@echo "Cleaned up build artifacts"
