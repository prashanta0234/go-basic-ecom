package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"e-com/internal/cache"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filter struct {
	Name  string `json:"name"`
	Page  int    `json:"page"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

type ProductsResponse struct {
	Data       []*domain.Products `json:"data"`
	Pagination PaginationMeta     `json:"pagination"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrev     bool  `json:"has_prev"`
}

func GetProducts(fil Filter) (*ProductsResponse, error) {
	cacheService := cache.NewCacheService()

	cacheKey := cacheService.GenerateProductsListKey(fil.Name, fil.Page, fil.Limit)

	var cachedResponse ProductsResponse
	if err := cacheService.Get(cacheKey, &cachedResponse); err == nil {
		log.Printf("Cache HIT for products list: %s", cacheKey)
		return &cachedResponse, nil
	}

	log.Printf("Cache MISS for products list: %s", cacheKey)

	collection := bootstrap.DB.Collection("products")

	var filter bson.M
	if fil.Name != "" {
		filter = bson.M{"name": bson.M{"$regex": fil.Name, "$options": "i"}}
	} else {
		filter = bson.M{}
	}

	totalCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, errors.New("failed to count products")
	}

	skip := (fil.Page - 1) * fil.Limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(fil.Limit))

	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}
	defer cursor.Close(context.TODO())

	var products []*domain.Products
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, errors.New("failed to decode products")
	}

	totalPages := int((totalCount + int64(fil.Limit) - 1) / int64(fil.Limit))
	hasNext := fil.Page < totalPages
	hasPrev := fil.Page > 1

	response := &ProductsResponse{
		Data: products,
		Pagination: PaginationMeta{
			CurrentPage: fil.Page,
			PerPage:     fil.Limit,
			Total:       totalCount,
			TotalPages:  totalPages,
			HasNext:     hasNext,
			HasPrev:     hasPrev,
		},
	}

	if err := cacheService.Set(cacheKey, response, cache.ProductTTL); err != nil {
		log.Printf("Failed to cache products list: %v", err)
	}

	return response, nil
}

func GetProductByID(productID string) (*domain.Products, error) {
	cacheService := cache.NewCacheService()

	cacheKey := cacheService.GenerateProductDetailKey(productID)

	var cachedProduct domain.Products
	if err := cacheService.Get(cacheKey, &cachedProduct); err == nil {
		log.Printf("Cache HIT for product detail: %s", productID)
		return &cachedProduct, nil
	}

	log.Printf("Cache MISS for product detail: %s", productID)

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

	if err := cacheService.Set(cacheKey, product, cache.ProductTTL); err != nil {
		log.Printf("Failed to cache product detail: %v", err)
	}

	return &product, nil
}
