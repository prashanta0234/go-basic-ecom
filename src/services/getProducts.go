package services

import (
	"context"
	"e-com/src/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(nameFilter string) ([]*models.Products, error) {

	collection := DB.Collection("products")

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

	var products []*models.Products
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, errors.New("failed to decode products")
	}

	return products, nil
}
