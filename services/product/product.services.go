package product

import (
	"ecommerce-api/models"
	repositories "ecommerce-api/repositories/product"
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("product not found")

func GetAllProducts() (*[]models.Product, error) {
	product, err := repositories.FindAll()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetByTitle(title string) (*[]models.Product, error) {
	product, err := repositories.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetById(userId string) (*models.Product, error) {
	product, err := repositories.FindById(userId)
	if err != nil {
		return nil, ErrNotFound
	}
	return product, nil
}

func CreateProduct(title string, description string, price string, path string) (*models.Product, error) {
	products := &models.Product{
		UserID:      uuid.NewString(),
		Title:       title,
		Description: description,
		Price:       price,
		Path:        path,
	}

	_ = repositories.CreateProduct(products)
	return products, nil
}

func UpdateProduct(userId string, title string, description string, price string, path string) (*models.Product, error) {
	input, err := repositories.FindById(userId)
	if err != nil {
		return nil, ErrNotFound
	}

	product := &models.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Path:        path,
	}

	_ = repositories.UpdateProduct(input, product)
	return product, nil
}

func DeleteProduct(userId string) error {
	product, err := repositories.FindById(userId)
	if err != nil {
		return ErrNotFound
	}

	_ = repositories.DeleteProduct(product, userId)
	return nil
}
