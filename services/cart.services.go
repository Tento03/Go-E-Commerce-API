package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
	"errors"
)

var ErrCartNotFound = errors.New("Cart not found")

func GetCart() (*models.Cart, error) {
	cart, err := repository.GetCart()
	if err != nil {
		return nil, ErrCartNotFound
	}

	return cart, nil
}
