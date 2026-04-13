package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}

func Error(c *gin.Context, code int, msg string){
	c.JSON(code, gin.H{
		"success": false,
		"data":    nil,
		"error":   msg,
	})
}
