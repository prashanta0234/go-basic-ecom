package usecase

import (
	"e-com/bootstrap"
	"e-com/domain"
	"e-com/repository"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(userID, productID string, sessionID string) (*domain.Order, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	if productID == "" {
		return nil, errors.New("product ID is required")
	}

	if sessionID == "" {
		return nil, errors.New("session ID is required")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	productObjectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	product, err := GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))

	orderID := fmt.Sprintf("ORD-%d-%s", time.Now().Unix(), sessionID[len(sessionID)-8:])

	order := &domain.Order{
		OrderID:      orderID,
		UserID:       userObjectID,
		ProductID:    productObjectID,
		ProductPrice: product.Price,
		PaidPrice:    product.Price,
	}

	err = orderRepo.Create(order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return order, nil
}

func GetOrderByID(orderID string) (*domain.Order, error) {
	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	return orderRepo.FindByID(orderID)
}

func GetOrdersByUserID(userID string) ([]*domain.Order, error) {
	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	return orderRepo.FindByUserID(userID)
}

func GetOrderByOrderID(orderID string) (*domain.Order, error) {
	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	return orderRepo.FindByOrderID(orderID)
}
