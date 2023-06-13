package cartsdto

type CartRequest struct {
	ProductID       uint   `json:"productID" form:"productID" validate:"required"`
	CustomerName    string `json:"customerName" form:"customerName" validate:"required"`
	CustomerAddress string `json:"customerAddress" form:"customerAddress" validate:"required"`
	CustomerPhone   string `json:"customerPhone" form:"customerPhone" validate:"required"`
	Price           int    `json:"price" form:"price" validate:"required"`
	Qty             int    `json:"qty" form:"qty" validate:"required"`
	Status          string `json:"status"`
}

type UpdateCartRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Image       string `json:"image" form:"image"`
}
