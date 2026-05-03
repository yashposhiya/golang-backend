package utils

import (
	"GoLang/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var Secret = []byte(config.AppConfig.Secretkey)

func GenerateToken(username string) (string, error) {
	fmt.Println(config.AppConfig.JWT_XPR)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
		"jti":      uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret)
}
