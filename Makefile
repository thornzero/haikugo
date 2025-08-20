.PHONY: build test clean install dev-deps coverage lint help all
.DEFAULT_GOAL := help

BINARY_NAME := haikuctl
CMD_DIR := ./cmd/haikuctl
BUILD_DIR := ./build
COVERAGE_DIR := ./coverage

# Build the binary
build: ## Build the haikuctl binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "✓ Built $(BUILD_DIR)/$(BINARY_NAME)"

# Run all tests
test: ## Run all unit tests
	@echo "Running tests..."
	@go test ./...
	@echo "✓ All tests passed"

# Run tests with verbose output
test-verbose: ## Run tests with verbose output
	@echo "Running tests (verbose)..."
	@go test -v ./...

# Generate test coverage report
coverage: ## Generate test coverage report
	@echo "Generating coverage report..."
	@mkdir -p $(COVERAGE_DIR)
	@go test -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	@go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "✓ Coverage report generated: $(COVERAGE_DIR)/coverage.html"

# Run linting checks
lint: ## Run linting checks
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found, running basic checks..."; \
		go vet ./...; \
		go fmt ./...; \
	fi
	@echo "✓ Linting complete"

# Clean build artifacts
clean: ## Clean build artifacts and caches
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(COVERAGE_DIR)
	@rm -f $(BINARY_NAME)
	@go clean -cache
	@echo "✓ Cleaned build artifacts"

# Install the binary to GOPATH/bin
install: ## Install haikuctl to GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@go install $(CMD_DIR)
	@echo "✓ Installed $(BINARY_NAME) to GOPATH/bin"

# Install development dependencies
dev-deps: ## Install development dependencies
	@echo "Installing development dependencies..."
	@go mod download
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@echo "✓ Development dependencies installed"

# Run the development version
dev-run: build ## Build and run with sample haiku
	@echo "Running development version..."
	@echo -e "an old silent pond\na frog jumps into the pond\nsplash silence again" | $(BUILD_DIR)/$(BINARY_NAME)

# Format all Go code
fmt: ## Format all Go code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

# Tidy up module dependencies
tidy: ## Tidy up module dependencies
	@echo "Tidying modules..."
	@go mod tidy
	@echo "✓ Modules tidied"

# Full development setup
setup: dev-deps tidy ## Full development environment setup
	@echo "✓ Development environment ready"

# Run all checks (test, lint, build)
all: clean fmt tidy test lint build ## Run all checks and build
	@echo "✓ All checks passed and binary built"

# Release build (with optimizations)
release: ## Build optimized release binary
	@echo "Building release version..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "✓ Release binary built: $(BUILD_DIR)/$(BINARY_NAME)"

# Show help
help: ## Show this help message
	@echo "haikugo - Haiku analysis and validation tool"
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*##"; printf ""} /^[a-zA-Z_-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
