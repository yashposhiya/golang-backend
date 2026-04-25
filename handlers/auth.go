package handlers

import (
	"GoLang/dto"
	"GoLang/models"
	"GoLang/repositories"
	"GoLang/services"
	"GoLang/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var req dto.UserRegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
		utils.Error(ctx, 400, "invalid data")
		return
	}

	//Mapping Req DTO -> Model
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := services.RegisterUser(user)
	if err != nil {
		utils.Error(ctx, 500, err.Error())
		return
	}

	//Mapping Model -> Response DTO
	resp := dto.UserResponse{
		Id:       newUser.Id,
		Username: newUser.Username,
	}
	utils.Success(ctx, 200, resp)
}

func LoginUser(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		utils.Error(ctx, 400, "invalid data")
		return
	}
	//Mapping Req DTO -> Model
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := services.LoginUser(user)
	if err != nil {
		utils.Error(ctx, 404, err.Error())
		return
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.Error(ctx, 500, "Internal server error")
		fmt.Println("JWT Token Generation Error:", err)
		return
	}

	// Mapping Model -> Resp DTO With Token
	resp := dto.UserResponse{
		Id:       newUser.Id,
		Username: newUser.Username,
		Token:    token,
	}
	utils.Success(ctx, 200, resp)
}

func LogoutUser(ctx *gin.Context) {
	val, exists := ctx.Get("jti")
	if !exists {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	jti, ok := val.(string)
	if !ok {
		ctx.JSON(401, gin.H{"error": "invalid token data"})
		return
	}
	if err := repositories.BlackListToken(jti); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Logout Success",
	})

}
