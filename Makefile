# Simple Makefile for a Go project

# Build the application
all: build-deps deps frontend build

build-deps:
	go mod tidy
	go install github.com/a-h/templ/cmd/templ@latest

build: 
	@echo "Building..."
	@go build -o fileserver .

frontend:
	@cd frontend && \
		npm run format && \
		npm run build

frontend-dev:
	@cd frontend && npm format && npm run dev

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests/... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f fileserver
	@rm -rf frontend/dist/*

dev:
	@make watch &
	@make frontend-dev


# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
					go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean frontend
