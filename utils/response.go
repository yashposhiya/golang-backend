package utils

import (
	"github.com/gin-gonic/gin"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {

	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, gin.H{
			"success": false,
			"data":    nil,
			"error":   appErr.Message,
		})
		return
	}

	c.JSON(500, gin.H{
		"success": false,
		"data":nil,
		"error":"Internal Server Error",
	})
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"success": false,
		"data":    nil,
		"error":   msg,
	})
}