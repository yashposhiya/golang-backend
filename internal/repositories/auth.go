package repositories

import (
	"GoLang/config"
	"GoLang/internal/models"
)

func InsertUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserByUsername(name string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", name).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
