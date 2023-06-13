package transactionsdto

type TransactionRequest struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"productID"`
	CustomerName    string `json:"customerName"`
	CustomerAddress string `json:"customerAddress"`
	CustomerPhone   string `json:"customerPhone"`
	UserID          uint   `json:"userID"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Status          string `json:"status"`
}
