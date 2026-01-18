package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func GetCart() (*models.Cart, error) {
	var cart models.Cart
	err := config.DB.Find(&cart).Error
	return &cart, err
}

func CreateCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}
