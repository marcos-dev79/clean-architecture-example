# Trafilea Prepare - Clean Architecture

This application follows Clean Architecture principles, ensuring separation of concerns and dependency inversion.

## Architecture Layers

### 1. Domain Layer (`domain/`)
The innermost layer containing business entities and interfaces. This layer has no dependencies on other layers.

- **`domain/entity/`**: Core business entities (e.g., `Subscription`)
- **`domain/repository/`**: Repository interfaces (contracts for data access)
- **`domain/usecase/`**: Use case interfaces (contracts for business logic)

### 2. Use Case Layer (`usecase/`)
Contains the application's business logic implementations. Depends only on the domain layer.

- **`usecase/subscription/`**: Implementation of subscription use cases

### 3. Infrastructure Layer (`infrastructure/`)
Implements the interfaces defined in the domain layer. Handles external concerns like databases, file systems, etc.

- **`infrastructure/repository/`**: Repository implementations
- **`infrastructure/db/`**: Database access code

### 4. Interface Layer (`interfaces/`)
Handles external interfaces like HTTP, gRPC, CLI, etc. Depends on use case layer.

- **`interfaces/http/handler/`**: HTTP request handlers

### 5. Entry Point (`cmd/`)
Application entry point that wires all dependencies together.

- **`cmd/main.go`**: Dependency injection and server setup

## Dependency Rule

The dependency rule states that:
- **Inner layers** (Domain) have no dependencies on outer layers
- **Outer layers** depend on inner layers through interfaces
- Dependencies point **inward** toward the domain

```
┌─────────────────────────────────────┐
│   cmd/ (Entry Point)                │
│   └─> interfaces/                   │
│       └─> usecase/                  │
│           └─> domain/               │
│   infrastructure/ ──────────────────┘
│       └─> domain/ (implements interfaces)
└─────────────────────────────────────┘
```

## Running the Application

```bash
go run cmd/main.go
```

The server will start on port 8080.

## API Endpoints

- `GET /subscriptions?user_id=<id>` - Get subscriptions for a user

## Benefits of Clean Architecture

1. **Independence**: Business logic is independent of frameworks, UI, and databases
2. **Testability**: Easy to test business logic without external dependencies
3. **Flexibility**: Easy to swap implementations (e.g., change database without affecting business logic)
4. **Maintainability**: Clear separation of concerns makes code easier to understand and maintain

