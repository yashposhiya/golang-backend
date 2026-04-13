package handlers

import (
	"GoLang/models"
	"GoLang/services"
	"GoLang/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		utils.Error(ctx, 400, "invalid data")
		return
	}
	newUser, err := services.RegisterUser(&user)
	if err != nil {
		utils.Error(ctx, 500, err.Error())
		return
	}
	utils.Success(ctx, 200, newUser)
}

func LoginUser(ctx *gin.Context) {
	var user models.UserLogin
	if err := ctx.BindJSON(&user); err != nil {
		utils.Error(ctx, 400, "invalid data")
		return
	}
	newUser, err := services.LoginUser(&user)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		utils.Error(ctx, 500, "Internal server error")
		fmt.Println("JWT Token Generation Error:", err)
		return
	}
	resp := models.UserLoginResponse{
		Id:    newUser.Id,
		Name:  newUser.Name,
		Token: token,
	}
	utils.Success(ctx, 200, resp)
}
