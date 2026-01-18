package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func FindAll() (*[]models.Product, error) {
	var product []models.Product
	err := config.DB.Find(&product).Error
	return &product, err
}

func FindById(productId string) (*models.Product, error) {
	var product models.Product
	err := config.DB.Model(&models.Product{}).Where("product_id = ?", productId).First(&product).Error
	return &product, err
}

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(productId string) error {
	return config.DB.Where("product_id = ?", productId).Delete(&models.Product{}).Error
}
