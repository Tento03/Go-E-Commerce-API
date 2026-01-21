package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func GetCart() (*models.Cart, error) {
	var cart models.Cart
	err := config.DB.Find(&cart).Error
	return &cart, err
}

func GetCartById(userId string) (*models.Cart, error) {
	var cart models.Cart
	err := config.DB.Model(&models.Cart{}).Where("user_id = ?", userId).First(&cart).Error
	return &cart, err
}

func GetCartItem(cartId string, productId string) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := config.DB.Model(&models.CartItem{}).Where("cart_id =? AND product_id =?", cartId, productId).Error
	return &cartItem, err
}

func CreateCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}

func UpdateCart(cart *models.Cart) error {
	return config.DB.Save(cart).Error
}

func DeleteCart(userId string) error {
	return config.DB.Where("user_id = ?", userId).Delete(&models.Cart{}).Error
}

func DeleteCartItem(productId string) error {
	return config.DB.Where("product_id = ?", productId).Delete(models.CartItem{}).Error
}
