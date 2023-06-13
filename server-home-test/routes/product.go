package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/handlers"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/pkg/middleware"
	"github.com/octadsp/server-home-test/repositories"
)

func ProductRoutes(r *gin.RouterGroup) {
	productRepository := repositories.RepositoryProduct(connectiondb.DB)
	h := handlers.HandlerProduct(productRepository)

	r.GET("/products", h.FindProducts)
	r.GET("/product/:id", h.GetProduct)
	r.POST("/product", middleware.Auth, middleware.UploadFile, h.AddProduct)
	r.PATCH("/product/:id", middleware.UploadFile, h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)
}
