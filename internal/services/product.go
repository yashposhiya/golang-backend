package services

import (
	"GoLang/internal/dto"
	"GoLang/internal/models"
	"GoLang/internal/repositories"
	"GoLang/internal/utils"
)

func CreateProduct(req dto.CreateProductRequest) (models.Product, error) {

	//Mapping Req DTO -> Model
	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	newProduct, err := repositories.CreateProduct(product)
	if err != nil {
		return models.Product{}, utils.InternalServerError()
	}
	return newProduct, nil
}

func GetProduct(id uint) (models.Product, error) {
	newProduct, err := repositories.GetProduct(id)
	if err != nil {
		return models.Product{}, utils.ProductNotFound()
	}
	return newProduct, nil
}

func DeleteProduct(id uint) error {
	err := repositories.DeleteProduct(id)
	if err != nil {
		return utils.ProductNotFound()
	}
	return nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return []models.Product{}, utils.ProductNotFound()
	}
	return products, nil
}

func UpdateProductFull(id uint, req dto.UpdateProductRequest) (models.Product, error) {
	//Mapping Req DTO -> Model
	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	return repositories.UpdateProductFull(id, product)
}

func ProductUpdatePartial(id uint, req dto.PatchUpdateProductRequest) (models.Product, error) {
	// Mapping Req DTO -> Map For Partial Updation
	product := make(map[string]any)
	if req.Name != nil {
		product["name"] = req.Name
	}
	if req.Price != nil {
		product["price"] = req.Price
	}
	return repositories.UpdateProductPartial(id, product)
}

func GetAllProductsPaginated(limit int, offset int)([]models.Product,error){
	return repositories.GetAllProductsPaginated(limit, offset)
}