package utils

import (
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

func GenerateAccessToken(userId string, username string, duration time.Duration) (string, error) {
	return GenerateToken(userId, username, 15*time.Minute)
}

func GenerateRefreshToken(userId string, username string, duration time.Duration) (string, error) {
	return GenerateToken(userId, username, duration)
}
