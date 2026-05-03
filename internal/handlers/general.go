package handlers

import(
	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message":"Hello From Go Backend",
	})
}