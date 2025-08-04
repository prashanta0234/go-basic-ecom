package main

import (
	"e-com/api/route"
	"e-com/bootstrap"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	bootstrap.InitMongoDB()

	// Setup routes
	r := route.SetupRoutes()

	fmt.Println("Server is running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
