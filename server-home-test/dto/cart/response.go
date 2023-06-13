package cartsdto

type CartResponse struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"productID"`
	CustomerName    string `json:"customerName"`
	CustomerAddress string `json:"customerAddress"`
	CustomerPhone   string `json:"customerPhone"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Status          string `json:"status"`
}

type FindCartsResponse struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"productID"`
	CustomerName    string `json:"customerName"`
	CustomerAddress string `json:"customerAddress"`
	CustomerPhone   string `json:"customerPhone"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Status          string `json:"status"`
}
type GetCartResponse struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"productID"`
	CustomerName    string `json:"customerName"`
	CustomerAddress string `json:"customerAddress"`
	CustomerPhone   string `json:"customerPhone"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Status          string `json:"status"`
}
type AddCartResponse struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"productID"`
	CustomerName    string `json:"customerName"`
	CustomerAddress string `json:"customerAddress"`
	CustomerPhone   string `json:"customerPhone"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Status          string `json:"status"`
}
