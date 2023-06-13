package models

type Product struct {
	ID          uint   `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	UserID      uint   `json:"userID"`
	User        User   `json:"users"`
}

func (Product) TableName() string {
	return "products"
}
