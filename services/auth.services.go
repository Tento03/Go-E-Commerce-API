package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"errors"

	"github.com/google/uuid"
)

var ErrUsernameExists = errors.New("username already exist")

func Register(username string, password string) (*models.Auth, error) {
	exist, err := repositories.IsUsernameExist(username)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, ErrUsernameExists
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.Auth{
		UserID:   uuid.NewString(),
		Username: username,
		Password: hashed,
	}

	if err := repositories.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
