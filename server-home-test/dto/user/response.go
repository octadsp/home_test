package usersdto

import (
	productsdto "github.com/octadsp/server-home-test/dto/product"
	transactionsdto "github.com/octadsp/server-home-test/dto/transaction"
	
)
type UserResponse struct {
	ID          uint                          `json:"id"`
	Email       string                        `json:"email"`
	Name        string                        `json:"name"`
	Role        string                        `json:"role"`
	Products     []productsdto.ProductResponse `json:"products"`
	Transactions []transactionsdto.TransactionResponse `json:"transactions"`
}
