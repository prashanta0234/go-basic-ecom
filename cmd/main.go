// @title E-Commerce API
// @version 1.0
// @description A complete e-commerce platform API built with Go, featuring user authentication, product management, payment processing with Stripe, and order management.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5050
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"e-com/api/route"
	"e-com/bootstrap"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	bootstrap.InitMongoDB()
	bootstrap.InitRedis()

	r := route.SetupRoutes()

	fmt.Println("Server is running at http://localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", r))
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
