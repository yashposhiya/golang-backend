package utils

import (
	"GoLang/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte(config.AppConfig.Secretkey)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret)
}
