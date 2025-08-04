# Go Basic E-commerce API

A clean architecture implementation of a basic e-commerce API built with Go.

## Project Structure

This project follows Clean Architecture principles with the following structure:

```
go-basic-ecom/
├── api/                    # API Layer
│   ├── controller/         # HTTP Controllers
│   ├── middleware/         # HTTP Middleware
│   └── route/             # Route Definitions
├── bootstrap/              # Application Bootstrap
├── cmd/                    # Application Entry Point
├── domain/                 # Domain Models/Entities
├── internal/               # Internal Utilities & DTOs
├── repository/             # Data Access Layer
├── usecase/                # Business Logic Layer
├── go.mod
├── go.sum
└── README.md
```

## Architecture Layers

### 1. API Layer (`api/`)

- **Controllers**: Handle HTTP requests and responses
- **Middleware**: Authentication, CORS, logging, etc.
- **Routes**: Route definitions and setup

### 2. Domain Layer (`domain/`)

- **Models**: Core business entities (User, Product, Order)
- Pure business logic without external dependencies

### 3. Use Case Layer (`usecase/`)

- **Business Logic**: Application-specific business rules
- Orchestrates domain objects and repositories
- Implements use cases like user registration, product management

### 4. Repository Layer (`repository/`)

- **Data Access**: Abstracts data persistence
- Implements repository pattern for database operations
- Handles MongoDB operations

### 5. Internal Layer (`internal/`)

- **Utilities**: Helper functions, DTOs, JWT handling
- Internal packages not meant for external use

### 6. Bootstrap Layer (`bootstrap/`)

- **Initialization**: Database connections, configuration
- Application startup and configuration

### 7. Command Layer (`cmd/`)

- **Entry Point**: Main application entry point
- Server setup and initialization

## Features

- **User Management**: Registration and login with JWT authentication
- **Product Management**: CRUD operations for products
- **Authentication**: JWT-based authentication middleware
- **Database**: MongoDB integration
- **Clean Architecture**: Separation of concerns with clear layer boundaries

## API Endpoints

### Authentication

- `POST /registration` - User registration
- `POST /login` - User login

### Products (Protected Routes)

- `GET /product` - Get all products (with optional name filter)
- `GET /product/{id}` - Get product by ID
- `POST /product` - Create new product
- `PUT /product/{id}` - Update product
- `DELETE /product/{id}` - Delete product

### Health Check

- `GET /health` - Health check endpoint

## Getting Started

1. **Install Dependencies**

   ```bash
   go mod tidy
   ```

2. **Environment Setup**

   **Option 1: Use the setup script**

   ```bash
   # Windows
   setup_env.bat

   # Unix/Linux/Mac
   chmod +x setup_env.sh
   ./setup_env.sh
   ```

   **Option 2: Create manually**
   Create a `.env` file with:

   ```
   MONGODB_URI=mongodb://localhost:27017
   MONGODB_NAME=ecommerce_db
   JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production
   PORT=5000
   ```

3. **Run the Application**

   ```bash
   go run cmd/main.go
   ```

4. **Access the API**
   The server will run on `http://localhost:5000`

## Clean Architecture Benefits

- **Separation of Concerns**: Each layer has a specific responsibility
- **Testability**: Easy to unit test each layer independently
- **Maintainability**: Clear structure makes code easy to understand and modify
- **Scalability**: Easy to add new features without affecting existing code
- **Dependency Inversion**: High-level modules don't depend on low-level modules

## Dependencies

- **MongoDB Driver**: For database operations
- **JWT**: For authentication
- **bcrypt**: For password hashing
- **godotenv**: For environment variable management
