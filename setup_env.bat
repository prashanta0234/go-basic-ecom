@echo off
echo Creating .env file with default values...

echo # MongoDB Configuration > .env
echo MONGODB_URI=mongodb://localhost:27017 >> .env
echo MONGODB_NAME=ecommerce_db >> .env
echo. >> .env
echo # JWT Configuration >> .env
echo JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production >> .env
echo. >> .env
echo # Server Configuration >> .env
echo PORT=5000 >> .env
echo. >> .env
echo # Stripe Configuration >> .env
echo STRIPE_SECRET_KEY=your_stripe_secret_key_here >> .env
echo STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key_here >> .env

echo .env file created successfully!
echo.
echo Please make sure MongoDB is running on localhost:27017
echo You can now run: go run cmd/main.go
pause 