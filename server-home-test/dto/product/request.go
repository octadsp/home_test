package productsdto

type ProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Qty         int    `json:"qty" form:"qty" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}

