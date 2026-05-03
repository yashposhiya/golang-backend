package handlers

import (
	"GoLang/internal/dto"
	"GoLang/internal/models"
	"GoLang/internal/repositories"
	"GoLang/internal/services"
	"GoLang/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Register User
// @Description Register or Signup User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterRequest true "User Request Body"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} utils.AppError
// @Router /auth/register [post]
func RegisterUser(ctx *gin.Context) {
	var req dto.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}

	//Mapping Req DTO -> Model
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := services.RegisterUser(user)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	//Mapping Model -> Response DTO
	resp := dto.UserResponse{
		Id:       newUser.Id,
		Username: newUser.Username,
	}
	utils.Success(ctx, 200, resp)
}

// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserLoginRequest true "User Login Request Body"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} utils.AppError
// @Router /auth/login [post]
func LoginUser(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}
	//Mapping Req DTO -> Model
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := services.LoginUser(user)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.HandleError(ctx, utils.InternalServerError())
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

// @Summary Logout User
// @Description User Logout
// @Tags Auth
// @Produce json
// @Success 200 {object} map[any]interface{}
// @Failure 400 {object} utils.AppError
// @Router /auth/logout [post]
func LogoutUser(ctx *gin.Context) {
	val, exists := ctx.Get("jti")
	if !exists {
		utils.HandleError(ctx, utils.InvalidJSON())
		return
	}

	jti, ok := val.(string)
	if !ok {
		utils.HandleError(ctx, utils.InvalidJSON())
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
