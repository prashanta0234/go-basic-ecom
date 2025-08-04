package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduct(productID string, userID string) error {
	collection := bootstrap.DB.Collection("products")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.New("invalid product ID format")
	}

	var existingProduct domain.Products
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&existingProduct)
	if err != nil {
		return errors.New("product not found")
	}

	if existingProduct.CreatedBy != userID {
		return errors.New("unauthorized: you can only delete your own products")
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return errors.New("failed to delete product")
	}

	if result.DeletedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}
