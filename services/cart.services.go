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

func CreateCart(userId string, cartId string, item *[]models.CartItem) (*models.Cart, error) {
	cart := &models.Cart{
		UserID: userId,
		Items:  item,
	}

	if err := repository.CreateCart(cart); err != nil {
		return nil, err
	}

	return cart, nil
}
