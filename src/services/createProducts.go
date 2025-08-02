package services

import (
	"context"
	"e-com/src/dto"
	"e-com/src/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProductsService(data dto.ProductsSchema, userID string) (*models.Products, error) {
	collection := DB.Collection("products")

	product := models.Products{
		Id:          primitive.NewObjectID(),
		Name:        data.Name,
		Description: data.Description,
		Image:       data.Image,
		Price:       data.Price,
		CreatedAt:   time.Now(),
		CreatedBy:   userID,
	}

	_, err := collection.InsertOne(context.TODO(), &product)

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &product, nil
}
