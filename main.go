package main

import (
	"e-com/src/controllers"
	"e-com/src/middleware"
	"e-com/src/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)

	json.NewEncoder(w).Encode("Health is good bro")
}

func main() {
	collection := services.DB.Collection("orders")
	fmt.Println("Collection reference ready:", collection.Name())

	r := http.NewServeMux()

	
	r.HandleFunc("/health", HandleRoot)
	
	r.HandleFunc("/registration", controllers.RegisterUserController)
	r.HandleFunc("/login", controllers.LoginController)
	
	r.HandleFunc("/product", middleware.AuthMiddleware(controllers.Products))
	r.HandleFunc("/product/", middleware.AuthMiddleware(controllers.Products))

	fmt.Println("Server is running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func init() {
	services.InitMongoDB()

}
