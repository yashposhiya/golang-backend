package handlers

import (
	"GoLang/helpers"
	"GoLang/models"
	"GoLang/services"
	"GoLang/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InsertProduct(ctx *gin.Context) {
	fmt.Println("Inside InsertProduct()")
	var product models.Product

	if err := ctx.BindJSON(&product); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}
	newProduct, err := services.CreateProduct(product)
	if err != nil {
		utils.Error(ctx, 500, err.Error())
	}
	utils.Success(ctx, 201, newProduct)
}

func GetProduct(ctx *gin.Context) {
	// uname, _ := ctx.Get("username")
	// fmt.Println("logged in user:", uname)
	idStr := ctx.Param("id")
	id, _ := helpers.StringToUint(idStr)
	product, err := services.GetProduct(id)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	utils.Success(ctx, 200, product)
}

func DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := helpers.StringToUint(idStr)
	err := services.DeleteProduct(id)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	ctx.JSON(204, gin.H{
		"success": true,
		"error":   nil,
		"data":    "no data",
	})
}

func GetAllProducts(ctx *gin.Context) {
	var data []models.Product
	data, err := services.GetAllProducts()
	if err != nil {
		utils.Error(ctx, 404, err.Error())
	}
	ctx.JSON(200, gin.H{"data": data})
}

func UpdateProductPut(ctx *gin.Context) {
	var product models.Product
	idStr := ctx.Param("id")
	id, _ := helpers.StringToUint(idStr)

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalud json"})
		return
	}

	newProduct, err := services.UpdateProductFull(id, product)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	utils.Success(ctx, 200, newProduct)
}

func ProductUpdatePartial(ctx *gin.Context) {
	var product models.ProductUpdatePartial
	idStr := ctx.Param("id")
	id, _ := helpers.StringToUint(idStr)
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(400, "Invalid json")
		return
	}

	newProduct, err := services.ProductUpdatePartial(id, product)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	utils.Success(ctx, 200, newProduct)
}
