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
