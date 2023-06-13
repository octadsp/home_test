package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productsdto "github.com/octadsp/server-home-test/dto/product"
	dto "github.com/octadsp/server-home-test/dto/result"
	transactionsdto "github.com/octadsp/server-home-test/dto/transaction"
	usersdto "github.com/octadsp/server-home-test/dto/user"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/repositories"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(c *gin.Context) {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertFindUsersResponse(users)})
}

func convertFindUsersResponse(users []models.User) []usersdto.UserResponse {
	responses := make([]usersdto.UserResponse, len(users))
	for i, u := range users {

		products := make([]productsdto.ProductResponse, len(u.Product))
		for j, m := range u.Product {
			products[j] = productsdto.ProductResponse{
				ID:          m.ID,
				Name:        m.Name,
				Description: m.Description,
				Price:       m.Price,
				Qty:         m.Qty,
				Image:       m.Image,
			}
		}

		transactions := make([]transactionsdto.TransactionResponse, len(u.Transaction))
		for j, m := range u.Transaction {
			transactions[j] = transactionsdto.TransactionResponse{
				ID:              m.ID,
				ProductID:       m.ProductID,
				CustomerName:    m.CustomerName,
				CustomerAddress: m.CustomerAddress,
				CustomerPhone:   m.CustomerPhone,
				UserID:          m.UserID,
				Price:           m.Price,
				Qty:             m.Qty,
				Status:          m.Status,
			}
		}

		responses[i] = usersdto.UserResponse{
			ID:           u.ID,
			Email:        u.Email,
			Name:         u.Name,
			Role:         u.Role,
			Products:     products,
			Transactions: transactions,
		}
	}
	return responses
}
