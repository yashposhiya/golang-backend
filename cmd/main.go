package main

import (
	"GoLang/config"
	"GoLang/internal/models"
	"GoLang/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	SwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()
	config.InitRedis()

	config.DB.AutoMigrate(&models.Product{}, &models.User{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(SwaggerFiles.Handler))
	routes.RegisterRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
