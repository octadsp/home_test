package models

type User struct {
	ID           uint          `json:"id" gorm:"primary_key:auto_increment"`
	Email        string        `json:"email" gorm:"unique;not null"`
	Password     string        `json:"password"`
	Name         string        `json:"name"`
	Role         string        `json:"role"`
	Product     []Product     `json:"products"`
	Transaction []Transaction `json:"transactions"`
}

func (User) TableName() string {
	return "users"
}
