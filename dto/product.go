package dto

type ProductResponse struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PatchUpdateProductRequest struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetAllProductsResponse struct {
	Count int               `json:"count"`
	Data  []ProductResponse `json:"data"`
}
