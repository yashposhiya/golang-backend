package utils

import "net/http"

// Internal Errors
func InternalServerError() *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	}
}

// Rate Limit Error
func TooManyRequests() *AppError {
	return &AppError{
		StatusCode: http.StatusTooManyRequests,
		Message:    "Too Many Requests",
	}
}

func InvalidID() *AppError {
	return &AppError{
		StatusCode: 400,
		Message:    "Invalid Id",
	}
}

func InvalidJSON() *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid JSON Request",
	}
}

// All the errors related to JWT
func InvalidJWT() *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid JWT Token",
	}
}

// All Error Related to Products
func ProductNotFound() *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    "Product Not Found",
	}
}

// All Error Related to User
func UserNotFound() *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    "User Not Found",
	}
}

func InvalidCredentials() *AppError {
	return &AppError{
		StatusCode: 404,
		Message:    "Invalid Credentials",
	}
}
