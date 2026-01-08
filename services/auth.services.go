package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrUsernameExists = errors.New("username already exist")
var ErrInvalidCredentials = errors.New("invalid credentials")

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
		Password: string(hashed),
	}

	if err := repositories.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Login(username string, password string) (string, error) {
	user, err := repositories.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
	}

	if !utils.ComparePassword(user.Password, password) {
		return "", ErrInvalidCredentials
	}

	return utils.GenerateToken(user.UserID, user.Username, 15*time.Minute)
}
