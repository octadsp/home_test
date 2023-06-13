package transactionsdto

type TransactionResponse struct {
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

type FindTransactionsResponse struct {
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

type GetTransactionResponse struct {
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

type AddTransactionResponse struct {
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

type UpdateTransactionResponse struct {
	ID              uint   `json:"id"`
	Status          string `json:"status"`
}
