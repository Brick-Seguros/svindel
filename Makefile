APP_NAME = svindel

# Default target
.DEFAULT_GOAL := help

# Run all tests
test: ## Run all tests
	go test -v ./...

# Run all tests with race detector
test-race: ## Run all tests with race detector
	go test -v -race ./...

# Run all tests with coverage
test-cover: ## Run all tests with coverage report
	go test -cover ./...

# Generate coverage HTML report
cover-html: ## Generate coverage HTML report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts and test cache
clean: ## Clean build artifacts and test cache
	go clean -testcache
	rm -f coverage.out

# Run the app (if you have a main.go in cmd/api or similar)
run: ## Run the application
	go run .

# Install dependencies (optional if you vend)
deps: ## Download dependencies
	go mod tidy

# Format the code
fmt: ## Format the code
	go fmt ./...

# Show help (auto-generated)
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'
