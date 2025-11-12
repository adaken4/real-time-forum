.PHONY: run test lint fmt build clean help

run: ## Run the application
	go run ./cmd/server/main.go

test: ## Run tests
	go test ./...

fmt: ## Format code with gofmt
	go fmt ./...

lint: fmt ## Run formatting and linting
	golangci-lint run

build: ## Build the binary for the current OS/ARCH into bin/
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/server ./cmd/server/main.go

clean: ## Clean built binaries and artifacts
	rm -rf bin/*
	rm -rf coverage.out

help: ## Show usage and available commands
	@echo "Usage: make <target>"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'
	@echo ""

