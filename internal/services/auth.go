package services

import (
	"GoLang/internal/models"
	"GoLang/internal/repositories"
	"GoLang/internal/utils"
)

func RegisterUser(user models.User) (models.User, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hash
	newUser, err := repositories.InsertUser(user)
	if err != nil {
		return models.User{}, utils.InternalServerError()
	}
	return newUser, nil
}

func LoginUser(user models.User) (models.User, error) {
	newUser, err := repositories.GetUserByUsername(user.Username)
	if err != nil {
		return models.User{}, utils.InvalidCredentials()
	}
	err = utils.CheckPassword(newUser.Password, user.Password)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}
