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

	r := gin.Default()

	// Auth
	r.POST("/user/register", handlers.RegisterUser)
	r.POST("/user/login", handlers.LoginUser)

	r.Use(middlewares.AuthMiddleware())
	// Products
	r.GET("/products/:id", handlers.GetProduct)
	r.POST("/products/", handlers.InsertProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.GET("/products", handlers.GetAllProducts)
	r.PUT("/products/:id", handlers.UpdateProductPut)
	r.PATCH("/products/:id", handlers.ProductUpdatePartial)
	fmt.Println("Routes inserted")

	r.Run(":"+config.AppConfig.Port)
}
