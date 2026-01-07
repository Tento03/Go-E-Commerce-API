package repositories

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func CreateUser(user *models.Auth) error {
	return config.DB.Create(user).Error
}

func IsUsernameExist(username string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Auth{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}
