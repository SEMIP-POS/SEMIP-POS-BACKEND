# Backend Service Project

## Overview
This is a backend service project built with Go (Fiber) that implements product management, transaction handling, and basic reporting features. The project uses a clean architecture pattern with distinct layers for better separation of concerns.

## Tech Stack
- **Backend Framework**: Go-Fiber
- **ORM**: GORM
- **Database**: PostgreSQL
- **Documentation**: Go Swagger
- **Authentication**: JWT

## Project Structure
```
.
├── cmd/                    # Application entry points
│   ├── rootCmd.go         # Root command configurations
│   └── serveHttp.go       # HTTP server initialization
├── config/                 # Configuration management
│   └── config.go          # Configuration structures and loading
├── internal/              # Private application code
│   └── domain/            # Business logic and entities
│       └── healthCheck/   # Health check implementation
├── repository/            # Data access layer
│   └── postgresql/        # PostgreSQL specific implementations
│       └── healthCheck/   # Health check repository
│           ├── check.go
│           ├── type.go
│           └── postgresql.go
├── service/               # Business logic layer
│   └── healthCheck/       # Health check service implementation
│       ├── check.go
│       ├── type.go
│       └── service.go
├── port/                  # Interface adapters
│   └── http/             # HTTP transport layer
│       └── handler/      # HTTP request handlers
│           └── healthCheck/
│               ├── check.go
│               └── type.go
├── .example.config.yaml   # Example configuration file
├── .example.secret.yaml   # Example secrets file
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksum
└── main.go               # Application entry point
```

## Features (MVP)
1. **Product Management**
   - Complete product inventory tracking
   - CRUD operations for product categories
   - CRUD operations for products

2. **Transaction Management**
   - Transaction processing and tracking

3. **Basic Reporting**
   - Essential business reporting features

4. **User Management** (Lower Priority)
   - User authentication and authorization

## Prerequisites
- Docker and Docker Compose
- Go 1.x or higher
- Make

## Getting Started

### Setup
1. Clone the repository
2. Copy the example configuration files:
   ```bash
   cp .example.config.yaml config.yaml
   cp .example.secret.yaml secret.yaml
   ```
3. Update the configuration files with your settings

### Running the Application

1. Start the required services using Docker Compose:
   ```bash
   docker-compose up -d
   ```

2. Run the application:
   ```bash
   make run
   ```

## Project Architecture
This project follows clean architecture principles:
- **Domain Layer**: Contains business logic and entities
- **Repository Layer**: Handles data persistence
- **Service Layer**: Implements business logic
- **Port Layer**: Manages external communications (HTTP handlers)

## Frontend Integration
The backend is designed to work with:
- React/TypeScript CMS with:
  - Tailwind CSS
  - ChakraUI
  - Zustand for state management
- Flutter mobile application using BLOC pattern

## Contributing
1. Follow the existing folder structure
2. Ensure proper error handling
3. Write tests for new features
4. Update documentation as needed