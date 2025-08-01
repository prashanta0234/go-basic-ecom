package main

import (
	"e-com/src/services"
	"fmt"
)

func main() {
	collection := services.DB.Collection("orders")
	fmt.Println("📦 Collection reference ready:", collection.Name())
}

func init() {
	services.InitMongoDB()
}
