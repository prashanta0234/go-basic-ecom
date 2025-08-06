package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"e-com/internal"
	"e-com/internal/cache"
	"errors"
	"log"
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

	cacheService := cache.NewCacheService()
	if err := cacheService.InvalidateProductCaches(); err != nil {
		log.Printf("Failed to invalidate product caches: %v", err)
	}

	return &product, nil
}
