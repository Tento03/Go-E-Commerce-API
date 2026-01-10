package product

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func FindAll() (*[]models.Product, error) {
	var product []models.Product
	err := config.DB.Find(&product).Error
	return &product, err
}

func FindById(userId string) (*models.Product, error) {
	var product models.Product
	err := config.DB.Model(&models.Product{}).Where("user_id = ?", userId).Error
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

func UpdateProduct(userId string, input *models.Product) error {
	var product models.Product
	err := config.DB.Model(input).Where("user_id = ?", userId).Updates(&product).Error
	return err
}

func DeleteProduct(input *models.Product, userId string) error {
	err := config.DB.Delete(input, userId).Error
	return err
}
