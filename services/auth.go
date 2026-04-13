package services

import (
	"GoLang/models"
	"GoLang/repositories"
	"fmt"
)

func RegisterUser(user *models.User) (*models.User, error) {
	newUser, err := repositories.InsertUser(user)
	if err != nil {
		return nil, fmt.Errorf("registration failed")
	}
	return newUser, nil
}

func LoginUser(user *models.UserLogin) (*models.User,error) {
	newUser,err := repositories.GetUserByUsername(user.Name)
	if err != nil{
		return nil, fmt.Errorf("User not found")
	}
	if newUser.Password == user.Password{
		return newUser,nil
	}
	return nil,fmt.Errorf("wrong credentials")
}

