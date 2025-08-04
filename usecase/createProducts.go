package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"e-com/internal"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProductsService(data internal.ProductsSchema, userID string) (*domain.Products, error) {
	collection := bootstrap.DB.Collection("products")

	product := domain.Products{
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
