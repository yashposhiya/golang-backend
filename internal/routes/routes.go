package routes

import (
	"GoLang/internal/handlers"
	"GoLang/internal/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/", handlers.Health)

	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/register", handlers.RegisterUser)
		userRoutes.POST("/login", handlers.LoginUser)
	}

	// Protected routes using the auth middleware
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/auth/logout", handlers.LogoutUser)

		// Product routes nested under a group
		products := auth.Group("/products")
		// products.Use(middlewares.RateLimiterMiddleware())
		{
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.InsertProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
			products.GET("", handlers.GetAllProducts)
			products.GET("/paginated",handlers.GetAllProductsPaginated)
			products.PUT("/:id", handlers.UpdateProductFull)
			products.PATCH("/:id", handlers.ProductUpdatePartial)
		}
	}
	fmt.Println("Routes Registered")
}
