package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	dto "github.com/octadsp/server-home-test/dto/result"
	transactionsdto "github.com/octadsp/server-home-test/dto/transaction"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/repositories"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(transactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{transactionRepository}
}

func (h *handlerTransaction) FindTransactions(c *gin.Context) {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertFindTransactionsResponse(transactions)})
}

func (h *handlerTransaction) GetTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertGetTransactionResponse(transaction)})
}

func (h *handlerTransaction) AddTransaction(c *gin.Context) {
	var err error
	request := new(transactionsdto.TransactionRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	transaction := models.Transaction{
		ProductID:       request.ProductID,
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		CustomerPhone:   request.CustomerPhone,
		UserID:          request.UserID,
		Price:           request.Price,
		Qty:             request.Qty,
		Status:          "pending",
	}

	transaction, err = h.TransactionRepository.AddTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	transaction, _ = h.TransactionRepository.GetTransaction(int(transaction.ID))

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertAddTransactionResponse(transaction)})
}

func (h *handlerTransaction) UpdateTransaction(c *gin.Context) {
	var err error
	request := new(transactionsdto.TransactionRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	transaction := models.Transaction{
		ID:              request.ID,
		ProductID:       request.ProductID,
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		CustomerPhone:   request.CustomerPhone,
		UserID:          request.UserID,
		Price:           request.Price,
		Qty:             request.Qty,
		Status:          request.Status,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err = h.TransactionRepository.GetTransaction(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.ProductID != 0 {
		transaction.ProductID = request.ProductID
	}

	if request.CustomerName != "" {
		transaction.CustomerName = request.CustomerName
	}

	if request.CustomerAddress != "" {
		transaction.CustomerAddress = request.CustomerAddress
	}

	if request.CustomerPhone != "" {
		transaction.CustomerPhone = request.CustomerPhone
	}

	if request.UserID != 0 {
		transaction.UserID = request.UserID
	}

	if request.Price != 0 {
		transaction.Price = request.Price
	}

	if request.Qty != 0 {
		transaction.Qty = request.Qty
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertUpdateTransactionResponse(data)})
}

func convertFindTransactionsResponse(transactions []models.Transaction) []transactionsdto.FindTransactionsResponse {
	responses := make([]transactionsdto.FindTransactionsResponse, len(transactions))
	for i, u := range transactions {

		responses[i] = transactionsdto.FindTransactionsResponse{
			ID:              u.ID,
			ProductID:       u.ProductID,
			CustomerName:    u.CustomerName,
			CustomerAddress: u.CustomerAddress,
			CustomerPhone:   u.CustomerPhone,
			UserID:          u.UserID,
			Price:           u.Price,
			Qty:             u.Qty,
			Status:          u.Status,
		}
	}
	return responses
}

func convertGetTransactionResponse(transaction models.Transaction) transactionsdto.GetTransactionResponse {

	response := transactionsdto.GetTransactionResponse{
		ID:              transaction.ID,
		ProductID:       transaction.ProductID,
		CustomerName:    transaction.CustomerName,
		CustomerAddress: transaction.CustomerAddress,
		CustomerPhone:   transaction.CustomerPhone,
		UserID:          transaction.UserID,
		Price:           transaction.Price,
		Qty:             transaction.Qty,
		Status:          transaction.Status,
	}

	return response
}

func convertAddTransactionResponse(transaction models.Transaction) transactionsdto.AddTransactionResponse {

	response := transactionsdto.AddTransactionResponse{
		ID:              transaction.ID,
		ProductID:       transaction.ProductID,
		CustomerName:    transaction.CustomerName,
		CustomerAddress: transaction.CustomerAddress,
		CustomerPhone:   transaction.CustomerPhone,
		UserID:          transaction.UserID,
		Price:           transaction.Price,
		Qty:             transaction.Qty,
		Status:          transaction.Status,
	}

	return response
}

func convertUpdateTransactionResponse(transaction models.Transaction) transactionsdto.UpdateTransactionResponse {

	response := transactionsdto.UpdateTransactionResponse{
		ID:     transaction.ID,
		Status: transaction.Status,
	}

	return response
}
