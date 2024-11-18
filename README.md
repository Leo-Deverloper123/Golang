# Hospital Middleware System

A Go-based middleware system that integrates with Hospital Information Systems (HIS) to provide unified patient information access.

## Features

- Staff authentication and authorization
- Patient information search across multiple hospitals
- Integration with external HIS APIs
- Hospital-specific access control
- Secure JWT-based authentication

## Tech Stack

- Go 1.21
- Gin Web Framework
- PostgreSQL with GORM
- Docker & Docker Compose
- Nginx
- JWT Authentication

## Prerequisites

- Docker
- Docker Compose
- Go 1.21 (for local development)

## Getting Started

1. Clone the repository
2. Create a `.env` file based on the example provided
3. Start the services:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:80`

## API Endpoints

### Public Endpoints

- `POST /staff/create` - Create new staff member
- `POST /staff/login` - Staff login

### Protected Endpoints

- `GET /patient/search` - Search for patients (requires authentication)

## Development

### Running Tests

```bash
go test ./tests/...
```

### Project Structure

```
.
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── nginx.conf
└── internal
    ├── handlers
    │   ├── patient_handler.go
    │   └── staff_handler.go
    ├── middleware
    │   └── auth_middleware.go
    ├── models
    │   ├── hospital.go
    │   ├── patient.go
    │   └── staff.go
    └── services
        ├── auth_service.go
        └── patient_service.go
```

## License

MIT