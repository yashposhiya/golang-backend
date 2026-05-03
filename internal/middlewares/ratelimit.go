package middlewares

import (
	"GoLang/internal/repositories"
	"GoLang/internal/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid, _ := ctx.Get("userid")
		key := fmt.Sprintf("rate_limit:user:%d", userid)
		_, err := repositories.AllowRequest(key, 5, 60*time.Second)
		if err != nil {
			utils.HandleError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
