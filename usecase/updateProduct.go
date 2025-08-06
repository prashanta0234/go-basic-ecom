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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProduct(productID string, data internal.ProductsSchema, userID string) (*domain.Products, error) {
	collection := bootstrap.DB.Collection("products")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	var existingProduct domain.Products
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&existingProduct)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if existingProduct.CreatedBy != userID {
		return nil, errors.New("unauthorized: you can only update your own products")
	}

	updateData := bson.M{
		"$set": bson.M{
			"name":        data.Name,
			"description": data.Description,
			"image":       data.Image,
			"price":       data.Price,
			"updatedAt":   time.Now(),
		},
	}

	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		updateData,
	)

	if err != nil {
		return nil, errors.New("failed to update product")
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("product not found")
	}

	var updatedProduct domain.Products
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&updatedProduct)
	if err != nil {
		return nil, errors.New("failed to fetch updated product")
	}

	cacheService := cache.NewCacheService()
	if err := cacheService.InvalidateSpecificProductCache(productID); err != nil {
		log.Printf("Failed to invalidate product cache: %v", err)
	}

	return &updatedProduct, nil
}
