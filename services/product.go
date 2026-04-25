package services

import (
	"GoLang/models"
	"GoLang/repositories"
	"GoLang/utils"
)

func CreateProduct(req models.Product) (models.Product, error) {
	newProduct, err := repositories.CreateProduct(req)
	if err != nil {
		return models.Product{}, &utils.AppError{
			StatusCode: 500,
			Message:    "Failed To Create Product",
		}
	}
	return newProduct, nil
}

func GetProduct(id uint) (models.Product, error) {
	newProduct, err := repositories.GetProduct(id)
	if err != nil {
		return models.Product{}, &utils.AppError{
			StatusCode: 404,
			Message:    "Product Not Found",
		}
	}
	return newProduct, nil
}

func DeleteProduct(id uint) error {
	err := repositories.DeleteProduct(id)
	if err != nil {
		return &utils.AppError{
			StatusCode: 404,
			Message:    "Product Not Found",
		}
	}
	return nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return []models.Product{}, &utils.AppError{
			StatusCode: 404,
			Message:    "No Products Found",
		}
	}
	return products, nil
}

func UpdateProductFull(id uint, product models.Product) (models.Product, error) {
	return repositories.UpdateProductFull(id, product)
}

func ProductUpdatePartial(id uint, product map[string]any) (models.Product, error) {
	return repositories.UpdateProductPartial(id, product)
}
