package controllers

import (
	"ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	cart, err := services.GetCart()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"product name": cart.Items.Product.Title,
			"quantity":     cart.Items.Qty,
		}})
}
