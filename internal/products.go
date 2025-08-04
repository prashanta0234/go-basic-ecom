package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsSchema struct {
	Name        string  `json:"name" binding:"required,min=3"`
	Description string  `json:"description" binding:"required,min=20"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required"`
}

type Products struct {
	Id          primitive.ObjectID
	Name        string
	Description string
	Image       string
	Price       float64
	CreatedAt   time.Time
}
