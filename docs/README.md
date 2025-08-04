# API Documentation

This directory contains the Swagger/OpenAPI documentation for the Go Basic E-commerce API.

## Files

- `index.html` - Swagger UI interface for viewing the API documentation
- `swagger.yaml` - OpenAPI 3.0 specification file (located in project root)

## How to View the Documentation

### Option 1: Open HTML File Directly

1. Open `docs/index.html` in your web browser
2. The documentation will load automatically

### Option 2: Serve with HTTP Server

```bash
# Using Python
python -m http.server 8080

# Using Node.js
npx serve .

# Using Go
go install github.com/shurcooL/goexec@latest
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

Then visit `http://localhost:8080/docs/`

### Option 3: Use Swagger Editor Online

1. Go to [Swagger Editor](https://editor.swagger.io/)
2. Copy the contents of `swagger.yaml`
3. Paste into the editor

## Features

- **Interactive Documentation**: Try out API endpoints directly from the browser
- **Authentication Support**: JWT Bearer token authentication
- **Request/Response Examples**: Detailed examples for all endpoints
- **Schema Definitions**: Complete data models for User and Product
- **Error Responses**: Comprehensive error handling documentation

## API Overview

### Authentication

- JWT-based authentication
- Token obtained from `/registration` or `/login` endpoints
- Include token in `Authorization: Bearer <token>` header

### Endpoints

#### Health Check

- `GET /health` - Check API status

#### Authentication

- `POST /registration` - Register new user
- `POST /login` - Login user

#### Products (Protected)

- `GET /product` - Get all products (with optional filtering)
- `GET /product/{id}` - Get specific product
- `POST /product` - Create new product
- `PUT /product/{id}` - Update product
- `DELETE /product/{id}` - Delete product

#### Payments

- `POST /payment/checkout` - Create checkout session (returns URL)
- `POST /payment/checkout/redirect` - Create checkout session (redirects)
- `GET /payment/success` - Payment success callback
- `GET /payment/cancel` - Payment cancel callback

## Testing the API

1. **Start the server**:

   ```bash
   go run cmd/main.go
   ```

2. **Register a user**:

   ```bash
   curl -X POST http://localhost:5000/registration \
     -H "Content-Type: application/json" \
     -d '{
       "name": "John Doe",
       "email": "john@example.com",
       "password": "password123"
     }'
   ```

3. **Login to get token**:

   ```bash
   curl -X POST http://localhost:5000/login \
     -H "Content-Type: application/json" \
     -d '{
       "email": "john@example.com",
       "password": "password123"
     }'
   ```

4. **Use the token** in subsequent requests:
   ```bash
   curl -X GET http://localhost:5000/product \
     -H "Authorization: Bearer <your_token>"
   ```

## Environment Setup

Make sure your `.env` file is configured:

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_NAME=ecommerce_db
JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production
STRIPE_SECRET_KEY=your_stripe_secret_key_here
STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key_here
```

## Notes

- All product endpoints require authentication
- Payment endpoints use product IDs to fetch prices from database
- JWT tokens expire after 24 hours
- Use test Stripe keys for development
- MongoDB must be running for the API to work
