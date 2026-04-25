package main

import (
	"GoLang/config"
	"GoLang/handlers"
	"GoLang/middlewares"
	"GoLang/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{}, &models.User{})

	config.InitRedis()

	r := gin.Default()

	// Auth
	r.POST("/users/register", handlers.RegisterUser)
	r.POST("/users/login", handlers.LoginUser)

	r.Use(middlewares.AuthMiddleware(), middlewares.RateLimiterMiddleware())
	//Auth - Logout
	r.POST("/users/logout", handlers.LogoutUser)

	// Products
	r.GET("/products/:id", handlers.GetProduct)
	r.POST("/products/", handlers.InsertProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.GET("/products", handlers.GetAllProducts)
	r.PUT("/products/:id", handlers.UpdateProductPut)
	r.PATCH("/products/:id", handlers.ProductUpdatePartial)
	fmt.Println("Routes inserted")

	r.Run(":" + config.AppConfig.Port)
}
