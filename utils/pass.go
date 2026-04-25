package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &AppError{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}
	}

	return string(hash), nil
}


func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if err != nil{
		return &AppError{
			StatusCode: 404,
			Message: "Invalid Password",
		}
	}
	return nil
}
