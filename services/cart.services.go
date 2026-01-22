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

func GetCartById(userId string) (*models.Cart, error) {
	cart, err := repository.GetCartById(userId)
	if err != nil {
		return nil, ErrCartNotFound
	}
	return cart, nil
}

func GetCartItem(cartId string, productId string) (*models.CartItem, error) {
	cart, err := repository.GetCartItem(cartId, productId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func CreateCart(userId string, cartId string, productId string, qty int) (*models.Cart, error) {
	items := &models.CartItem{
		CartID:    cartId,
		ProductID: productId,
		Qty:       qty,
	}

	cart := &models.Cart{
		UserID: userId,
		Items:  items,
	}

	if err := repository.CreateCart(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func UpdateCart(userId string, qty int) (*models.Cart, error) {
	cart, err := repository.GetCartById(userId)
	if err != nil {
		return nil, err
	}

	cart.Items.Qty = qty

	if err := repository.UpdateCart(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func DeleteCart(userId string) error {
	_, err := repository.GetCartById(userId)
	if err != nil {
		return err
	}
	return repository.DeleteCart(userId)
}

func DeleteCartItem(cartId string, productId string) error {
	_, err := repository.GetCartItem(cartId, productId)
	if err != nil {
		return err
	}

	return repository.DeleteCartItem(productId)
}
