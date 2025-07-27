# Makefile for Portfolio Backend

.PHONY: help build run test clean docker-build docker-run docker-stop deps lint format init-atlas

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application in development mode"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download dependencies"
	@echo "  lint         - Run linter"
	@echo "  format       - Format code"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  init-atlas   - Initialize MongoDB Atlas database"

# Build the application
build:
	@echo "Building portfolio backend..."
	go build -o portfolio-backend cmd/main.go

# Run the application
run:
	@echo "Running portfolio backend..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f portfolio-backend
	go clean

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
format:
	@echo "Formatting code..."
	go fmt ./...
	go vet ./...

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t portfolio-backend .

# Run with Docker Compose
docker-run:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

# Stop Docker containers
docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down

# Initialize MongoDB Atlas database
init-atlas:
	@echo "Initializing MongoDB Atlas database..."
	@if [ ! -d "node_modules" ]; then \
		echo "Installing Node.js dependencies..."; \
		npm install; \
	fi
	@echo "Running database initialization script..."
	npm run init-db

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Setup development environment
setup: install-tools deps
	@echo "Development environment setup complete!"

# Database operations
db-reset:
	@echo "Resetting database..."
	docker-compose down -v
	docker-compose up -d mongodb
	@echo "Waiting for MongoDB to start..."
	sleep 10
	docker-compose up -d portfolio-backend

# Production build
prod-build:
	@echo "Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o portfolio-backend cmd/main.go

# Generate API documentation
docs:
	@echo "Generating API documentation..."
	@echo "API documentation is available in README.md"

# Security check
security:
	@echo "Running security checks..."
	go list -json -deps ./... | nancy sleuth

# Performance test
bench:
	@echo "Running benchmarks..."
	go test -bench=. ./... 