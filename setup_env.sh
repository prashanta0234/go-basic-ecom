#!/bin/bash

echo "Creating .env file with default values..."

cat > .env << EOF
# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_NAME=ecommerce_db

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production

# Server Configuration
PORT=5000

# Stripe Configuration
STRIPE_SECRET_KEY=your_stripe_secret_key_here
STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key_here
EOF

echo ".env file created successfully!"
echo ""
echo "Please make sure MongoDB is running on localhost:27017"
echo "You can now run: go run cmd/main.go" 