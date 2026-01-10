package repositories

import (
	"ecommerce-api/config"
	models "ecommerce-api/models/auth"
)

func CreateUser(user *models.Auth) error {
	return config.DB.Create(user).Error
}

func FindByUsername(username string) (*models.Auth, error) {
	var user models.Auth
	err := config.DB.Model(&models.Auth{}).Where("username = ?", username).First(&user).Error
	return &user, err
}

func IsUsernameExist(username string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Auth{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}
