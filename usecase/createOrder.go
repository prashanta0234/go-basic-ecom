package usecase

import (
	"e-com/bootstrap"
	"e-com/domain"
	"e-com/internal/cache"
	"e-com/repository"
	"errors"
	"fmt"
	"log"
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

	cacheService := cache.NewCacheService()
	if err := cacheService.InvalidateUserOrdersCache(userID); err != nil {
		log.Printf("Failed to invalidate user orders cache: %v", err)
	}

	return order, nil
}

func GetOrderByID(orderID string) (*domain.Order, error) {
	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	return orderRepo.FindByID(orderID)
}

func GetOrdersByUserID(userID string) ([]*domain.Order, error) {
	cacheService := cache.NewCacheService()

	cacheKey := cacheService.GenerateUserOrdersKey(userID)

	var cachedOrders []*domain.Order
	if err := cacheService.Get(cacheKey, &cachedOrders); err == nil {
		log.Printf("Cache HIT for user orders: %s", userID)
		return cachedOrders, nil
	}

	log.Printf("Cache MISS for user orders: %s", userID)

	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	orders, err := orderRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if err := cacheService.Set(cacheKey, orders, cache.UserTTL); err != nil {
		log.Printf("Failed to cache user orders: %v", err)
	}

	return orders, nil
}

func GetOrderByOrderID(orderID string) (*domain.Order, error) {
	orderRepo := repository.NewOrderRepository(bootstrap.DB.Collection("orders"))
	return orderRepo.FindByOrderID(orderID)
}
