package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
	"errors"
)

var ErrNotFound = errors.New("product not found")

func GetAllProducts() (*[]models.Product, error) {
	product, err := repository.FindAll()
	if err != nil {
		return nil, ErrNotFound
	}
	return product, nil
}

func GetById(productId string) (*models.Product, error) {
	product, err := repository.FindById(productId)
	if err != nil {
		return nil, ErrNotFound
	}
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
