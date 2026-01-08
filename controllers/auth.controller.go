package controllers

import (
	"ecommerce-api/requests"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"
	"os"

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
	c.JSON(http.StatusCreated,
		gin.H{
			"message":  "register success",
			"userId":   user.UserID,
			"username": user.Username,
		})
}

func Login(c *gin.Context) {
	var req requests.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationError(err)})
		return
	}

	acccessToken, refreshToken, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secured := os.Getenv("APP_ENV") == "production"

	c.SetCookie("accessToken", acccessToken, 15*60, "/", "", secured, true)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", secured, true)

	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}
