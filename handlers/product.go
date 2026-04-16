package handlers

import (
	"GoLang/dto"
	"GoLang/helpers"
	"GoLang/models"
	"GoLang/services"
	"GoLang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertProduct(ctx *gin.Context) {
	fmt.Println("Inside InsertProduct()")
	var req dto.CreateProductRequest

	if err := ctx.BindJSON(&req); err != nil {
		utils.Error(ctx, 400, err.Error())
		return
	}

	//Mapping Req DTO -> Model
	product := models.Product{
		Name: req.Name,
		Price: req.Price,
	}

	newProduct, err := services.CreateProduct(product)
	if err != nil {
		utils.Error(ctx, 500, err.Error())
		return
	}

	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id: newProduct.Id,
		Name: newProduct.Name,
		Price: newProduct.Price,
	}

	utils.Success(ctx, 201, resp)
}

func GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id,err := helpers.StringToUint(idStr)
	if err != nil{
		utils.Error(ctx,401,"Invalid Id")
	}

	product, err := services.GetProduct(id)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}

	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
	}
	utils.Success(ctx, 200, resp)
}

func DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id,err := helpers.StringToUint(idStr)
	if err != nil{
		utils.Error(ctx,401,"Invalid Id")
		return
	}
	err = services.DeleteProduct(id)
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
	data, err := services.GetAllProducts()
	if err != nil {
		utils.Error(ctx, 404, err.Error())
	}
	products := make([]dto.ProductResponse,0,len(data))
	for _,p := range data{
		resp := dto.ProductResponse{
			Id: p.Id,
			Name: p.Name,
			Price: p.Price,
		}
		products = append(products, resp)
	}
	ctx.JSON(200, gin.H{"data": products})
}

func UpdateProductPut(ctx *gin.Context) {
	var req dto.UpdateProductRequest
	idStr := ctx.Param("id")
	id,err := helpers.StringToUint(idStr)
	if err != nil{
		utils.Error(ctx,401,"Invalid Id")
		return
	}

	if err = ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalud json"})
		return
	}

	//Mapping Req DTO -> Model
	product := models.Product{
		Name: req.Name,
		Price: req.Price,
	}
	newProduct, err := services.UpdateProductFull(id, product)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}

	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id: newProduct.Id,
		Name: newProduct.Name,
		Price: newProduct.Price,
	}
	utils.Success(ctx, 200, resp)
}

func ProductUpdatePartial(ctx *gin.Context) {
	var req dto.PatchUpdateProductRequest
	idStr := ctx.Param("id")
	id, err := helpers.StringToUint(idStr)
	if err != nil{
		utils.Error(ctx,401,"Invalid Id")
		return
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, "Invalid json")
		return
	}

	// Mapping Req DTO -> Map For Partial Updation
	product := make(map[string]any)
	if req.Name != nil{
		product["name"] = req.Name
	}
	if req.Price != nil{
		product["price"] = req.Price
	}
	newProduct, err := services.ProductUpdatePartial(id, product)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}

	// Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id: newProduct.Id,
		Name: newProduct.Name,
		Price: newProduct.Price,
	}
	utils.Success(ctx, 200, resp)
}