package cache

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"encoding/json"
	"time"
)

func GetList(ctx context.Context, key string) ([]*models.Product, error) {
	val, err := config.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func SetList(ctx context.Context, key string, products []*models.Product, ttl time.Duration) error {
	bytes, _ := json.Marshal(products)
	return config.Client.Set(ctx, key, bytes, ttl).Err()
}
