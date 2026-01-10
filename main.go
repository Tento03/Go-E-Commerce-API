package main

import (
	"ecommerce-api/config"
	models "ecommerce-api/models/auth"
	models2 "ecommerce-api/models/product"

	"ecommerce-api/routes"
	"ecommerce-api/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Auth{}, &models.Refresh{}, &models2.Product{})
	r := gin.Default()
	routes.AuthRoutes(r)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validators.PasswordValidator)
	}
	r.Run(":8080")
}
