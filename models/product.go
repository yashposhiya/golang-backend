package models

type Product struct {
	Id    uint    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price float64    `json:"price"`
}

type ProductUpdatePartial struct {
	Name  *string `json:"name"`
	Price *float64    `json:"price"`
}
