# API Quick Reference

## Base URL

```
http://localhost:5000
```

## Authentication

Include JWT token in header:

```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### 🔐 Authentication

| Method | Endpoint        | Description       | Auth Required |
| ------ | --------------- | ----------------- | ------------- |
| POST   | `/registration` | Register new user | ❌            |
| POST   | `/login`        | Login user        | ❌            |

### 🏥 Health Check

| Method | Endpoint  | Description      | Auth Required |
| ------ | --------- | ---------------- | ------------- |
| GET    | `/health` | Check API status | ❌            |

### 📦 Products (Protected)

| Method | Endpoint        | Description        | Auth Required |
| ------ | --------------- | ------------------ | ------------- |
| GET    | `/product`      | Get all products   | ✅            |
| GET    | `/product/{id}` | Get product by ID  | ✅            |
| POST   | `/product`      | Create new product | ✅            |
| PUT    | `/product/{id}` | Update product     | ✅            |
| DELETE | `/product/{id}` | Delete product     | ✅            |

### 💳 Payments

| Method | Endpoint                     | Description                   | Auth Required |
| ------ | ---------------------------- | ----------------------------- | ------------- |
| POST   | `/payment/checkout`          | Create checkout (returns URL) | ❌            |
| POST   | `/payment/checkout/redirect` | Create checkout (redirects)   | ❌            |
| GET    | `/payment/success`           | Payment success callback      | ❌            |
| GET    | `/payment/cancel`            | Payment cancel callback       | ❌            |

## Request Examples

### Register User

```bash
curl -X POST http://localhost:5000/registration \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login User

```bash
curl -X POST http://localhost:5000/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Product

```bash
curl -X POST http://localhost:5000/product \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{
    "name": "MacBook Pro",
    "description": "High-performance laptop for professionals",
    "image": "https://example.com/macbook.jpg",
    "price": 1299.99
  }'
```

### Get Products

```bash
curl -X GET http://localhost:5000/product \
  -H "Authorization: Bearer <your_token>"
```

### Create Payment

```bash
curl -X POST http://localhost:5000/payment/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "507f1f77bcf86cd799439011",
    "currency": "usd"
  }'
```

## Response Examples

### Successful Registration/Login

```json
{
  "message": "Registration successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Product Response

```json
{
  "message": "Products fetched successfully!",
  "data": [
    {
      "_id": "507f1f77bcf86cd799439011",
      "name": "MacBook Pro",
      "description": "High-performance laptop for professionals",
      "image": "https://example.com/macbook.jpg",
      "price": 1299.99,
      "createdAt": "2024-01-15T10:30:00Z",
      "updatedAt": "2024-01-15T10:30:00Z",
      "createdBy": "507f1f77bcf86cd799439012"
    }
  ]
}
```

### Payment Response

```json
{
  "checkout_url": "https://checkout.stripe.com/pay/cs_test_...",
  "message": "Checkout session created successfully"
}
```

## Error Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 303  | Redirect (Payment)    |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 404  | Not Found             |
| 500  | Internal Server Error |

## Environment Variables

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_NAME=ecommerce_db
JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production
STRIPE_SECRET_KEY=your_stripe_secret_key_here
STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key_here
```

## Notes

- **JWT Tokens**: Expire after 24 hours
- **Product Prices**: Automatically fetched from database using product ID
- **Authentication**: Required for all product operations
- **Payment**: Uses Stripe for checkout processing
- **Database**: MongoDB required for operation
