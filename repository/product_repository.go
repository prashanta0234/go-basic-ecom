package repository

import (
	"context"
	"e-com/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Create(product *domain.Products) error
	FindByID(id string) (*domain.Products, error)
	FindAll(nameFilter string) ([]*domain.Products, error)
	Update(id string, product *domain.Products) error
	Delete(id string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

func (r *productRepository) Create(product *domain.Products) error {
	_, err := r.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return errors.New("failed to create product")
	}
	return nil
}

func (r *productRepository) FindByID(id string) (*domain.Products, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	var product domain.Products
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (r *productRepository) FindAll(nameFilter string) ([]*domain.Products, error) {
	var filter bson.M
	if nameFilter != "" {
		filter = bson.M{"name": bson.M{"$regex": nameFilter, "$options": "i"}}
	} else {
		filter = bson.M{}
	}

	cursor, err := r.collection.Find(context.TODO(), filter)
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

func (r *productRepository) Update(id string, product *domain.Products) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid product ID format")
	}

	updateData := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"image":       product.Image,
			"price":       product.Price,
			"updatedAt":   product.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		updateData,
	)

	if err != nil {
		return errors.New("failed to update product")
	}

	if result.MatchedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid product ID format")
	}

	result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return errors.New("failed to delete product")
	}

	if result.DeletedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}
