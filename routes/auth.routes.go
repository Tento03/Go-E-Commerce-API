package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", middleware.LoginRateLimiter(5, 2*time.Minute), controllers.Login)
		auth.POST("/refresh", middleware.RequireAuth, middleware.RefreshTokenLimiter(5, 2*time.Minute), controllers.Refresh)
		auth.POST("/logout", controllers.Logout)
	}
}
