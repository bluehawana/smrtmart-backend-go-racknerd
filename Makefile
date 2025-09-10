# SmrtMart Backend Makefile

.PHONY: help build run test clean docker-build docker-run docker-stop migrate-up migrate-down

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the Go application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  migrate-up   - Run database migrations"
	@echo "  migrate-down - Rollback database migrations"
	@echo "  deps         - Download dependencies"
	@echo "  lint         - Run linter"
	@echo "  swagger      - Generate Swagger documentation"

# Build the application
build:
	@echo "Building SmrtMart API..."
	go build -o bin/smrtmart-api ./cmd/server

# Run the application locally
run:
	@echo "Starting SmrtMart API..."
	go run ./cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf uploads/
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

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger docs..."
	swag init -g cmd/server/main.go -o docs

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t smrtmart-api .

docker-run:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down

docker-logs:
	@echo "Showing Docker logs..."
	docker-compose logs -f

# Database migration commands
migrate-up:
	@echo "Running database migrations..."
	migrate -path migrations -database "postgres://postgres:postgres123@localhost:5432/smrtmart_db?sslmode=disable" up

migrate-down:
	@echo "Rolling back database migrations..."
	migrate -path migrations -database "postgres://postgres:postgres123@localhost:5432/smrtmart_db?sslmode=disable" down

migrate-create:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $$name

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	cp .env.example .env
	@echo "Please edit .env file with your configuration"
	@echo "Then run: make docker-run"

# Production deployment
deploy:
	@echo "Deploying to production..."
	docker build -t smrtmart-api:latest .
	# Add your deployment commands here

# Database seed (for development)
seed:
	@echo "Seeding database with sample data..."
	# Add seed commands here