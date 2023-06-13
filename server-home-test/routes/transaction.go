package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/handlers"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/repositories"
)

func TransactionRoutes(r *gin.RouterGroup) {
	transactionRepository := repositories.RepositoryTransaction(connectiondb.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.GET("/transactions", h.FindTransactions)
	r.GET("/transaction/:id", h.GetTransaction)
	r.POST("/transaction", h.AddTransaction)
	r.PATCH("/transaction/:id", h.UpdateTransaction)
}
