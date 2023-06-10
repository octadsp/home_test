package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
}

func (Product) TableName() string {
	return "products"
}
