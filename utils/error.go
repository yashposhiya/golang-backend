package utils

func InvalidID() *AppError {
	return &AppError{
		StatusCode: 400,
		Message:    "Invalid Id",
	}
}

func InvalidJSON() *AppError {
	return &AppError{
		StatusCode: 400,
		Message:    "Invalid JSON Request",
	}
}
