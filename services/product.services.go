package services

import (
	"ecommerce-api/cache"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"ecommerce-api/repository"
	"ecommerce-api/requests"
	"errors"
	"fmt"
	"log"
	"time"
)

var ErrNotFound = errors.New("product not found")

func GetAllProducts() (*[]models.Product, error) {
	var req requests.GetAllProducts

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}

	cacheKey := fmt.Sprintf(
		"products:page=%d:limit=%d",
		req.Page,
		req.Limit,
	)

	products, err := cache.GetList(config.Ctx, cacheKey)
	if err == nil && len(*products) > 0 {
		log.Println("GET ALL -> CACHE HIT:", cacheKey)
		return products, nil
	}

	log.Println("GET ALL -> CACHE MISS:", cacheKey)

	product, err := repository.FindAll()
	if err != nil {
		return nil, ErrNotFound
	}

	_ = cache.SetList(config.Ctx, cacheKey, product, 5*time.Minute)
	return product, nil
}

func GetById(productId string) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:id=%s", productId)

	productCache, err := cache.GetById(config.Ctx, cacheKey)
	if err == nil && productCache != nil {
		log.Println("GET BY ID -> CACHE HIT:", productId)
		return productCache, nil
	}

	log.Println("GET BY ID -> CACHE MISS:", productId)

	product, err := repository.FindById(productId)
	if err != nil {
		return nil, ErrNotFound
	}

	_ = cache.SetById(config.Ctx, cacheKey, product, 5*time.Minute)
	return product, nil
}

func CreateProduct(productId string, title string, description string, price string, types string, path string) (*models.Product, error) {
	product := &models.Product{
		ProductID:   productId,
		Title:       title,
		Description: description,
		Price:       price,
		Type:        types,
		Path:        path,
	}
	err := repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(productId string, title string, description string, price string, types string, path string) (*models.Product, error) {
	product, err := repository.FindById(productId)
	if err != nil {
		return nil, ErrNotFound
	}

	product.Title = title
	product.Description = description
	product.Price = price
	product.Type = types
	product.Path = path

	if err := repository.UpdateProduct(product); err != nil {
		return nil, err
	}

	_ = cache.Delete(config.Ctx, product.ProductID)

	return product, nil
}

func DeleteProduct(productId string) error {
	_, err := repository.FindById(productId)
	if err != nil {
		return ErrNotFound
	}

	_ = repository.DeleteProduct(productId)
	return nil
}
