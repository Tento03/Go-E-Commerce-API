package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrUsernameExists = errors.New("username already exist")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrRefreshReuse = errors.New("refresh token expired or reused")

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

func Login(username string, password string) (string, string, error) {
	user, err := repositories.FindByUsername(username)
	if err != nil {
		return "", "", err
	}

	if !utils.ComparePassword(user.Password, password) {
		return "", "", ErrInvalidCredentials
	}

	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.UserID)
	if err != nil {
		return "", "", err
	}

	hashRT := utils.HashToken(refreshToken)
	refresh := &models.Refresh{
		UserID:    user.UserID,
		Token:     hashRT,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := repositories.SaveRefreshToken(refresh); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
