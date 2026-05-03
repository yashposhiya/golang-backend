package handlers

import (
	_ "GoLang/docs"
	"GoLang/helpers"
	"GoLang/internal/dto"
	"GoLang/internal/services"
	"GoLang/internal/utils"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
)

// @Summary Create Product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product Data"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} utils.AppError
// @Router /products [post]
func InsertProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}

	// Passing DTO to model as data validation is there
	newProduct, err := services.CreateProduct(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	// still service will return model for future flexibility
	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id:    newProduct.Id,
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}

	utils.Success(ctx, 201, resp)
}

// @Summary Get Product
// @Description Get a product with an id
// @Tags Products
// @Produce json
// @Param id path int true "Product Id"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} utils.AppError
// @Router /products/{id} [get]
func GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := helpers.StringToUint(idStr)
	if err != nil {
		utils.HandleError(ctx, utils.InvalidID())
		return
	}

	product, err := services.GetProduct(id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
	utils.Success(ctx, 200, resp)
}

// @Summary Delete Product
// @Description Delete product by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} utils.AppError
// @Router /products/{id} [delete]new(type)
func DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := helpers.StringToUint(idStr)
	if err != nil {
		utils.HandleError(ctx, utils.InvalidID())
		return
	}
	err = services.DeleteProduct(id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(204, gin.H{
		"success": true,
		"error":   nil,
		"data":    "no data",
	})
}

// @Summary Get All Products
// @Description Get all the products
// @Tags Products
// @Produce json
// @Success 200 {object} dto.GetAllProductsResponse
// @Failure 400 {object} utils.AppError
// @Router /products [get]
func GetAllProducts(ctx *gin.Context) {
	data, err := services.GetAllProducts()
	if err != nil {
		utils.HandleError(ctx, err)
	}
	products := make([]dto.ProductResponse, 0, len(data))
	for _, p := range data {
		resp := dto.ProductResponse{
			Id:    p.Id,
			Name:  p.Name,
			Price: p.Price,
		}
		products = append(products, resp)
	}
	// time.Sleep(3 * time.Second)
	ctx.JSON(200, gin.H{"data": products})
}

// @Summary Update Product Full
// @Description Update full product using product id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Param product body dto.UpdateProductRequest true "Product Update Request Body Put"
// @Success 200 {object} dto.ProductResponse
// @Failure 404 {object} utils.AppError
// @Router /products/{id} [put]
func UpdateProductFull(ctx *gin.Context) {
	var req dto.UpdateProductRequest
	idStr := ctx.Param("id")
	id, err := helpers.StringToUint(idStr)
	if err != nil {
		utils.HandleError(ctx, utils.InvalidID())
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}

	newProduct, err := services.UpdateProductFull(id, req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	//Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id:    newProduct.Id,
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}
	utils.Success(ctx, 200, resp)
}

// @Summary Update Product Partial
// @Description Update partial product using product id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Param product body dto.UpdateProductRequest true "Product Update Request Body Patch"
// @Success 200 {object} dto.PatchUpdateProductRequest
// @Failure 404 {object} utils.AppError
// @Router /products/{id} [patch]
func ProductUpdatePartial(ctx *gin.Context) {
	var req dto.PatchUpdateProductRequest
	idStr := ctx.Param("id")
	id, err := helpers.StringToUint(idStr)
	if err != nil {
		utils.HandleError(ctx, utils.InvalidID())
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}

	newProduct, err := services.ProductUpdatePartial(id, req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	// Mapping Model -> Resp DTO
	resp := dto.ProductResponse{
		Id:    newProduct.Id,
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}
	utils.Success(ctx, 200, resp)
}

func GetAllProductsPaginated(ctx *gin.Context) {
	var pageStr = ctx.Query("page")
	var limitStr = ctx.Query("limit")

	page := 1
	limit := 20

	if pageStr != ""{
		if p,err := helpers.StringToInt(pageStr); err == nil && p>0{
			page = p
		}
	}
	if limitStr != "" {
        if l, err := helpers.StringToInt(limitStr); err == nil && l > 0 {
            limit = l
        }
    }

	offset := (page-1)*limit

	data,err := services.GetAllProductsPaginated(limit,offset)
	if err!= nil{
		utils.HandleError(ctx, err)
		return
	}
	
	products := make([]dto.ProductResponse,0,limit)
	for _,p := range data{
		product := dto.ProductResponse{
			Name: p.Name,
			Price: p.Price,
			Id: p.Id,
		}
		products = append(products, product)
	}
	response := dto.GetAllProductsPaginatedResponse{
		Limit: limit,
		Page: page,
		Data: products,
	}
	time.Sleep(time.Second*1)
	utils.Success(ctx, 200, response)
}