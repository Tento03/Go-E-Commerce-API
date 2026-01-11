package routes

import (
	controllers "ecommerce-api/controllers/product"
	"ecommerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/product")
	product.Use(middleware.RequireAuth)
	{
		product.GET("/", controllers.GetAll)
		product.GET("/:id", controllers.GetById)
		product.POST("/", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}
}
