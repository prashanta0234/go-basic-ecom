package repository

import (
	"context"
	"e-com/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id string) (*domain.Order, error)
	FindByUserID(userID string) ([]*domain.Order, error)
	FindByOrderID(orderID string) (*domain.Order, error)
}

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) OrderRepository {
	return &orderRepository{
		collection: collection,
	}
}

func (r *orderRepository) Create(order *domain.Order) error {
	if order.ID.IsZero() {
		order.ID = primitive.NewObjectID()
	}

	order.CreatedAt = primitive.NewDateTimeFromTime(primitive.NewObjectID().Timestamp())

	_, err := r.collection.InsertOne(context.TODO(), order)
	if err != nil {
		return errors.New("failed to create order")
	}
	return nil
}

func (r *orderRepository) FindByID(id string) (*domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	var order domain.Order
	err = r.collection.FindOne(context.TODO(), primitive.M{"_id": objectID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		return nil, errors.New("failed to find order")
	}
	return &order, nil
}

func (r *orderRepository) FindByUserID(userID string) ([]*domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	cursor, err := r.collection.Find(context.TODO(), primitive.M{"userId": objectID})
	if err != nil {
		return nil, errors.New("failed to find orders")
	}
	defer cursor.Close(context.TODO())

	var orders []*domain.Order
	for cursor.Next(context.TODO()) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, errors.New("failed to decode order")
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *orderRepository) FindByOrderID(orderID string) (*domain.Order, error) {
	var order domain.Order
	err := r.collection.FindOne(context.TODO(), primitive.M{"orderId": orderID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		return nil, errors.New("failed to find order")
	}
	return &order, nil
}
