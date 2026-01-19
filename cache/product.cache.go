package cache

import (
	"context"
	"ecommerce-api/models"
	"time"
)

type ProductCache interface {
	GetList(ctx context.Context, key string) ([]*models.Product, error)
	SetList(ctx context.Context, key string, product []*models.Product, ttl time.Duration) error
}
