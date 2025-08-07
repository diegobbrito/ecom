GO = go
PROJECT_DIR = $(shell pwd)


build:
	@echo "(1/2) Building the server..."
	@$(GO) build -o $(PROJECT_DIR)/bin/ecom $(PROJECT_DIR)/cmd/main.go
	@echo "(2/2) Server built into: $(PROJECT_DIR)/bin/ecom"

test:
	@go test -v ./...

run: build
	@echo "(1/1) Running the server..."
	@PORT=$(PORT) $(PROJECT_DIR)/bin/ecom

clean:
	@echo "(1/2)Cleaning up build artifacts..."
	@rm -rf $(PROJECT_DIR)/bin
	@echo "(2/2) $(PROJECT_DIR)/bin Cleanup complete."

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

.DEFAULT_GOAL := run