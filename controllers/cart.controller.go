package controllers

import (
	"ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCart(c *gin.Context) {
	cart, err := services.GetCart()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success", "data": gin.H{"cartId": cart.Items.CartID}})
}

func GetCartById(c *gin.Context) {
	userId := c.GetString("userId")

	cart, err := services.GetCartById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": gin.H{"cartId": cart.Items.CartID}})
}

func GetCartItem(c *gin.Context) {
	userId := c.GetString("userId")
	cart, err := services.GetCartById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": gin.H{
		"CartId":    cart.Items.CartID,
		"ProductId": cart.Items.ProductID,
		"Quantity":  cart.Items.Qty,
	}})
}

func CreateCart(c *gin.Context) {
	userId := c.GetString("userId")
	cartId := uuid.NewString()

	cart, err := services.CreateCart(userId, cartId)
}
