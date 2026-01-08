package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userId string, username string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      time.Now().Add(duration).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(userId string, username string) (string, error) {
	return GenerateToken(userId, username, 15*time.Minute)
}

func GenerateRefreshToken(userId string, username string) (string, error) {
	return GenerateToken(userId, username, 7*24*time.Hour)
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if !token.Valid || err != nil {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}
