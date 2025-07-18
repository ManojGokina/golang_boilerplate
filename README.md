# Backend API Boilerplate

A scalable, production-ready Go backend API boilerplate with clean architecture principles.

## Features

- **Clean Architecture** - Separation of concerns with layered architecture
- **Domain-Driven Design** - Business logic encapsulation
- **RESTful API** - Standard HTTP methods and status codes
- **Database Integration** - MongoDB with native driver
- **Caching** - Redis support (optional)
- **Middleware** - CORS, Rate limiting, Logging, Recovery
- **Configuration** - Environment-based configuration
- **Logging** - Structured logging with Logrus
- **Containerization** - Docker and Docker Compose
- **API Documentation** - Swagger/OpenAPI support
- **Testing** - Unit test structure
- **Graceful Shutdown** - Proper server shutdown handling

## Project Structure

```
├── cmd/
│   └── api/                 # Application entrypoint
├── internal/
│   ├── config/             # Configuration management
│   ├── domain/             # Domain models and interfaces
│   ├── handler/            # HTTP handlers
│   ├── middleware/         # HTTP middleware
│   ├── repository/         # Data access layer
│   ├── service/            # Business logic layer
│   └── database/           # Database connection
├── pkg/
│   ├── logger/             # Logging utilities
│   ├── response/           # API response utilities
│   └── validator/          # Validation utilities
├── migrations/             # Database migrations
└── docs/                   # API documentation
```

## Getting Started

### Prerequisites

- Go 1.21+
- MongoDB
- Redis (optional)
- Docker & Docker Compose (optional)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd backend
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Update environment variables in `.env`

4. Install dependencies:
```bash
make deps
```

### Running the Application

#### Local Development

```bash
# Run with Go
make run

# Or build and run
make build
./bin/main
```

#### With Docker Compose

```bash
# Start all services (app, database, redis)
make docker-run

# Stop services
make docker-stop
```

### API Endpoints

#### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
- `GET /api/v1/users` - List users (paginated)

#### Health Check
- `GET /health` - Health check endpoint

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| ENVIRONMENT | Application environment | development |
| PORT | Server port | 8080 |
| MONGO_URL | MongoDB connection string | required |
| MONGO_DB | MongoDB database name | backend |
| REDIS_URL | Redis connection string | optional |
| JWT_SECRET | JWT signing secret | required |
| LOG_LEVEL | Logging level | info |
| RATE_LIMIT_RPS | Rate limit requests per second | 100 |

### Development Commands

```bash
# Format code
make fmt

# Run tests
make test

# Lint code
make lint

# Generate Swagger docs
make swagger

# Clean build artifacts
make clean
```

### Architecture Principles

1. **Dependency Inversion** - High-level modules don't depend on low-level modules
2. **Interface Segregation** - Clients shouldn't depend on interfaces they don't use
3. **Single Responsibility** - Each module has one reason to change
4. **Repository Pattern** - Abstract data access layer
5. **Service Layer** - Business logic encapsulation

### Adding New Features

1. Define domain models in `internal/domain/`
2. Create repository interface and implementation
3. Implement service layer with business logic
4. Add HTTP handlers for API endpoints
5. Register routes in main application

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...
```

### Deployment

The application is containerized and can be deployed using:

- Docker containers
- Kubernetes
- Cloud platforms (AWS, GCP, Azure)
- Traditional servers

### Contributing

1. Follow Go conventions and best practices
2. Write tests for new features
3. Update documentation
4. Use meaningful commit messages

## License

This project is licensed under the MIT License.