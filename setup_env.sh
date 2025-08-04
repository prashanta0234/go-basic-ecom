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
EOF

echo ".env file created successfully!"
echo ""
echo "Please make sure MongoDB is running on localhost:27017"
echo "You can now run: go run cmd/main.go" 