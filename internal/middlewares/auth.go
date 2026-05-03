package middlewares

import (
	"GoLang/internal/repositories"
	"GoLang/internal/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("Bearer")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error")
			}
			return utils.Secret, nil
		})

		if err != nil || !token.Valid {
			fmt.Println("not valid or err", err)
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("No claims")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		jtiRaw, exists := claims["jti"]
		if !exists || jtiRaw == nil {
			fmt.Println("jti not there")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		jti, ok := jtiRaw.(string)
		if !ok {
			fmt.Println("jti not in format")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		// claims := token.Claims.(jwt.MapClaims)

		// jti := claims["jti"].(string)
		err = repositories.CheckBlacklistToken(jti)
		if err != nil {
			fmt.Println("not blacklist")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		usernameRaw, exists := claims["username"]
		if !exists || usernameRaw == nil {
			fmt.Println("no username")
			utils.HandleError(ctx, utils.InvalidJWT())
			ctx.Abort()
			return
		}

		username, ok := usernameRaw.(string)
		if !ok {
			fmt.Println("username format not right")
			utils.HandleError(ctx, utils.InvalidJWT())
		}

		//store in context for future reference or use
		ctx.Set("username", username)
		ctx.Set("jti", jti)

		ctx.Next()
	}
}
