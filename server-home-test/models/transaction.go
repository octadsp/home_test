package models

type Transaction struct {
	ID              uint    `json:"id" gorm:"primary_key:auto_increment"`
	ProductID       uint    `json:"productID"`
	Product         Product `json:"products"`
	CustomerName    string  `json:"customerName"`
	CustomerAddress string  `json:"customerAddress"`
	CustomerPhone   string  `json:"customerPhone"`
	UserID          uint    `json:"userID"`
	User            User    `json:"users"`
	Price           int     `json:"price"`
	Qty             int     `json:"qty"`
	Status          string  `json:"status"`
}

func (Transaction) TableName() string {
	return "transactions"
}
