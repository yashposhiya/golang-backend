package services

import (
	"GoLang/models"
	"GoLang/repositories"
	"GoLang/utils"
	"fmt"
)

func RegisterUser(user models.User) (models.User, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hash
	newUser, err := repositories.InsertUser(user)
	if err != nil {
		return models.User{}, fmt.Errorf("registration failed")
	}
	return newUser, nil
}

func LoginUser(user models.User) (models.User, error) {
	newUser, err := repositories.GetUserByUsername(user.Username)
	if err != nil {
		return models.User{}, fmt.Errorf("User not found")
	}
	// if newUser.Password == user.Password {
	// 	return newUser, nil
	// }
	err = utils.CheckPassword(newUser.Password, user.Password)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}
