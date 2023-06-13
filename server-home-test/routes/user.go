package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/handlers"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/repositories"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(connectiondb.DB)
	h := handlers.HandlerUser(userRepository)

	r.GET("/users", h.FindUsers)
}
