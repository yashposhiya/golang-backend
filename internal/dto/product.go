package dto

type ProductResponse struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PatchUpdateProductRequest struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price" binding:"omitempty,gt=0"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type CreateProductRequest struct {
	Name  string  `json:"name"  binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type GetAllProductsResponse struct {
	Count int               `json:"count"`
	Data  []ProductResponse `json:"data"`
}

type GetAllProductsPaginatedResponse struct {
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
	Data  []ProductResponse `json:"data"`
}
