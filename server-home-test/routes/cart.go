package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/handlers"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/repositories"
)

func CartRoutes(r *gin.RouterGroup) {
	cartRepository := repositories.RepositoryCart(connectiondb.DB)
	h := handlers.HandlerCart(cartRepository)

	r.GET("/carts", h.FindCarts)
	r.POST("/cart", h.AddCart)
	r.DELETE("/cart/:id", h.DeleteCart)
}
