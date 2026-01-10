package product

import (
	"ecommerce-api/config"
	models "ecommerce-api/models/product"
)

func FindAll() (*[]models.Product, error) {
	var product []models.Product
	err := config.DB.Find(&product).Error
	return &product, err
}

func FindById(productId string) (*models.Product, error) {
	var product models.Product
	err := config.DB.Model(&models.Product{}).Where("product_id = ?", productId).Error
	return &product, err
}

func FindByTitle(title string) (*[]models.Product, error) {
	var product []models.Product
	err := config.DB.Model(&models.Product{}).Where("title = ?", title).First(&product).Error
	return &product, err
}

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func UpdateProduct(input *models.Product, product *models.Product) error {
	err := config.DB.Model(input).Updates(&product).Error
	return err
}

func DeleteProduct(input *models.Product, productId string) error {
	err := config.DB.Delete(input, productId).Error
	return err
}
