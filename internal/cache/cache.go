package cache

import (
	"context"
	"crypto/md5"
	"e-com/bootstrap"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	DefaultTTL = 15 * time.Minute
	ProductTTL = 30 * time.Minute
	UserTTL    = 10 * time.Minute
)

const (
	ProductsListKey  = "products:list:%s"
	ProductDetailKey = "product:detail:%s"
	UserOrdersKey    = "user:orders:%s"
	ProductSearchKey = "products:search:%s"
)

type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService() *CacheService {
	return &CacheService{
		client: bootstrap.RedisClient,
		ctx:    context.Background(),
	}
}

func (c *CacheService) Set(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(c.ctx, key, data, ttl).Err()
}

func (c *CacheService) Get(key string, dest interface{}) error {
	data, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (c *CacheService) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

func (c *CacheService) DeletePattern(pattern string) error {
	keys, err := c.client.Keys(c.ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return c.client.Del(c.ctx, keys...).Err()
	}

	return nil
}

func (c *CacheService) Exists(key string) bool {
	result, err := c.client.Exists(c.ctx, key).Result()
	return err == nil && result > 0
}

func (c *CacheService) GenerateProductsListKey(name string, page, limit int) string {
	hash := c.generateHash(fmt.Sprintf("%s:%d:%d", name, page, limit))
	return fmt.Sprintf(ProductsListKey, hash)
}

func (c *CacheService) GenerateProductDetailKey(productID string) string {
	return fmt.Sprintf(ProductDetailKey, productID)
}

func (c *CacheService) GenerateUserOrdersKey(userID string) string {
	return fmt.Sprintf(UserOrdersKey, userID)
}

func (c *CacheService) generateHash(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func (c *CacheService) InvalidateProductCaches() error {
	patterns := []string{
		"products:list:*",
		"products:search:*",
		"product:detail:*",
	}

	for _, pattern := range patterns {
		if err := c.DeletePattern(pattern); err != nil {
			return err
		}
	}

	return nil
}

func (c *CacheService) InvalidateUserOrdersCache(userID string) error {
	key := c.GenerateUserOrdersKey(userID)
	return c.Delete(key)
}

func (c *CacheService) InvalidateSpecificProductCache(productID string) error {
	productKey := c.GenerateProductDetailKey(productID)
	if err := c.Delete(productKey); err != nil {
		return err
	}

	return c.DeletePattern("products:list:*")
}
