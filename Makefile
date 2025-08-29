GO = go
PROJECT_DIR = $(shell pwd)
BIN_DIR = $(PROJECT_DIR)/bin
BIN = $(BIN_DIR)/ecom
MAIN = $(PROJECT_DIR)/cmd/main.go

.PHONY: build test run clean migration migrate-up migrate-down

build:
    @echo "(1/2) Building the server..."
    @mkdir -p $(BIN_DIR)
    @$(GO) build -o $(BIN) $(MAIN)
    @echo "(2/2) Server built into: $(BIN)"

test:
    @$(GO) test -v ./...

run: build
    @echo "(1/1) Running the server..."
    @PORT=$(PORT) $(BIN)

clean:
    @echo "(1/2) Cleaning up build artifacts..."
    @rm -f $(BIN)
    @echo "(2/2) $(BIN) cleanup complete."

migration:
    @migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
    @$(GO) run cmd/migrate/main.go up

migrate-down:
    @$(GO) run cmd/migrate/main.go down

.DEFAULT_GOAL := run