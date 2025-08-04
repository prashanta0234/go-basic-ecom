package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(nameFilter string) ([]*domain.Products, error) {

	collection := bootstrap.DB.Collection("products")

	var filter bson.M
	if nameFilter != "" {
		filter = bson.M{"name": bson.M{"$regex": nameFilter, "$options": "i"}}
	} else {
		filter = bson.M{}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}
	defer cursor.Close(context.TODO())

	var products []*domain.Products
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, errors.New("failed to decode products")
	}

	return products, nil
}

func GetProductByID(productID string) (*domain.Products, error) {
	collection := bootstrap.DB.Collection("products")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	var product domain.Products
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return &product, nil
}
