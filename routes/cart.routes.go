package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/cart")
	cart.Use(middleware.RequireAuth)
	{
		cart.GET("/", controllers.GetCart)
		cart.GET("/:id", controllers.GetCartById)
		cart.POST("/", controllers.CreateCart)
		cart.PUT("/:id", controllers.UpdateCart)
		cart.DELETE("/:id", controllers.DeleteCartItem)
	}
}
