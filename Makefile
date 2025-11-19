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

.PHONY: migrate-up migrate-down migrate-create migrate-force

DB_URL=sqlite3://rt-forum.db
MIGRATION_PATH=migrations

migrate-up: ## Apply all up migrations
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down: ## Apply all down migrations
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) down 1

migrate-create: ## Create a new migration file. Usage: make migrate-create NAME=<name>
	@if [ -z "$(name)" ]; then echo "Error: Please provide a name for the migration (e.g., make migrate-create name=add_users)"; exit 1; fi
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migrate-force: ## Force set the migration version. Usage: make migrate-force VERSION=<version>
	@if [ -z "$(version)" ]; then echo "Error: Please provide a version number (e.g., make migrate-force version=1); exit 1; fi
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) force $(version)
