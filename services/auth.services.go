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

	return accessToken, refreshToken, err
}

func Refresh(userId string, refreshToken string) (string, string, error) {
	old, err := repositories.FindValidRefreshToken(userId, refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = repositories.RevokeAllUser(userId)
			return "", "", ErrRefreshReuse
		}
		return "", "", err
	}

	if err := repositories.RevokeToken(old); err != nil {
		return "", "", err
	}

	newAccessToken, _ := utils.GenerateAccessToken(userId)
	newRefreshToken, _ := utils.GenerateRefreshToken(userId)

	refresh := &models.Refresh{
		UserID:    userId,
		Token:     newRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := repositories.SaveRefreshToken(refresh); err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
