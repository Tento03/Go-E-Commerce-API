package repositories

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func FindByUsername(username string) (*models.Auth, error) {
	var auth models.Auth
	err := config.DB.Where("username = ?", username).Error
	return &auth, err
}
