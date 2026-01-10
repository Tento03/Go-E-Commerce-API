package routes

import (
	controllers "ecommerce-api/controllers/auth"
	"ecommerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/refresh", middleware.RequireAuth, controllers.Refresh)
		auth.POST("/logout", controllers.Logout)
	}
}
