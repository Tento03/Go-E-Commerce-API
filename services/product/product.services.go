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
		return nil, err
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

func CreateProduct(productId, title string, description string, price string, types string, path string) (*models.Product, error) {
	products := &models.Product{
		ProductID:   productId,
		Title:       title,
		Description: description,
		Price:       price,
		Type:        types,
		Path:        path,
	}

	_ = repositories.CreateProduct(products)
	return products, nil
}

func UpdateProduct(productId string, title string, description string, price string, types string, path string) (*models.Product, error) {
	input, err := repositories.FindById(productId)
	if err != nil {
		return nil, ErrNotFound
	}

	product := &models.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Type:        types,
		Path:        path,
	}

	_ = repositories.UpdateProduct(input, product)
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
