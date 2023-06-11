package productsdto

type ProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Qty         int    `json:"qty" form:"qty" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Image       string `json:"image" form:"image"`
}
