package repositories

import (
	"GoLang/config"
	"GoLang/models"
	"GoLang/utils"

	"gorm.io/gorm/clause"
)

func CreateProduct(product models.Product) (models.Product, error) {
	err := config.DB.Create(&product).Error
	return product, err
}

func GetProduct(id uint) (models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	return product, err
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}

func UpdateProductFull(id uint, product models.Product) (models.Product, error) {
	product.Id = id
	result := config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(&product)
	if result.RowsAffected == 0 {
		return models.Product{}, &utils.AppError{
			StatusCode: 404,
			Message:    "Product Not Found",
		}
	}
	return product, nil
}

func UpdateProductPartial(id uint, product map[string]any) (models.Product, error) {
	var newProduct models.Product

	res := config.DB.Model(&newProduct).Clauses(clause.Returning{}).Where("id = ?", id).Updates(product)

	if res.Error != nil {
		return models.Product{}, &utils.AppError{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}
	}
	if res.RowsAffected == 0 {
		return models.Product{}, &utils.AppError{
			StatusCode: 404,
			Message:    "Product Not Found",
		}
	}
	return newProduct, nil
}
