package authenticationdto

type RegisterRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
