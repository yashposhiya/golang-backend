package services

import (
	"GoLang/models"
	"GoLang/repositories"
	"fmt"
)

var database []models.Product
var nextId = 1

func CreateProduct(req models.Product) (models.Product, error) {
	newProduct, err := repositories.CreateProduct(req)
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to create product")
	}
	return newProduct,nil
}

func GetProduct(id uint) (models.Product, error) {
	newProduct, err := repositories.GetProduct(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("product not found with id %d", id)
	}
	return newProduct, nil
}

func DeleteProduct(id uint) error {
	err := repositories.DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("product not found")
	}
	return nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return []models.Product{}, fmt.Errorf("no products found")
	}
	return products, nil
}

func UpdateProductFull(id uint, product models.Product) (models.Product, error) {
	return repositories.UpdateProductFull(id, product)
}

func ProductUpdatePartial(id uint, product map[string]any) (models.Product, error) {
	return repositories.UpdateProductPartial(id, product)
}
