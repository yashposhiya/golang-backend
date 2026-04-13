package repositories

import(
	"GoLang/config"
	"GoLang/models"
)

func CreateUser(user models.User)(models.User,error){
	err := config.DB.Save(&user).Error

	return user,err
}