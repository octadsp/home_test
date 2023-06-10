package authenticationdto

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `gorm:"type: varchar(20)" json:"role"`
}
