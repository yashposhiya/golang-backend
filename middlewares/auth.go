package middlewares

import (
	"GoLang/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			utils.Error(ctx, 401, "Missing token")
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(ctx, 401, "Invalid token format")
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return utils.Secret, nil
		})

		if err != nil || !token.Valid {
			utils.Error(ctx, 401, "Invalid token")
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		username := claims["username"].(string)

		//store in context for future reference or use
		ctx.Set("username",username)

		ctx.Next()
	}
}
