package controllers

import (
	"ecommerce-api/requests"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req requests.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": utils.ValidationError(err)})
		return
	}

	user, err := services.Register(req.Username, req.Password)
	if err != nil {
		if err == services.ErrUsernameExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"userId":   user.UserID,
		"username": user.Username,
	})
}
