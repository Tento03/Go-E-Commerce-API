package product

import (
	models "ecommerce-api/models/product"
	repositories "ecommerce-api/repositories/product"
	"errors"
)

var ErrNotFound = errors.New("product not found")

func GetAllProducts() (*[]models.Product, error) {
	product, err := repositories.FindAll()
	if err != nil {
		return nil, ErrNotFound
	}
	return product, nil
}

func GetById(productId string) (*models.Product, error) {
	product, err := repositories.FindById(productId)
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
	err := repositories.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(productId string, title string, description string, price string, types string, path string) (*models.Product, error) {
	product, err := repositories.FindById(productId)
	if err != nil {
		return nil, ErrNotFound
	}

	product.Title = title
	product.Description = description
	product.Price = price
	product.Type = types
	product.Path = path

	if err := repositories.UpdateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(productId string) error {
	_, err := repositories.FindById(productId)
	if err != nil {
		return ErrNotFound
	}

	_ = repositories.DeleteProduct(productId)
	return nil
}
