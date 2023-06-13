package productsdto

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
}

type FindProductsResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	UserID      uint   `json:"userID"`
}
type GetProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	UserID      uint   `json:"userID"`
}
type AddProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	UserID      uint   `json:"userID"`
}
