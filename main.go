package main

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"ecommerce-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Auth{})
	r := gin.Default()
	routes.AuthRoutes(r)
	r.Run(":8080")
}
