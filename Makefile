.PHONY: build run test clean docker-build docker-run migrate-up migrate-down

# Build the application
build:
	go build -o bin/main cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t backend-api .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Stop Docker Compose
docker-stop:
	docker-compose down

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate Swagger docs
swagger:
	swag init -g cmd/api/main.go

# Database migrations (placeholder)
migrate-up:
	@echo "Run database migrations up"

migrate-down:
	@echo "Run database migrations down"