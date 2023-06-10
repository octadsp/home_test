package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/handlers"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/repositories"
)

func AuthenticationRoutes(r *gin.RouterGroup) {
	authenticationRepository := repositories.RepositoryAuthentication(connectiondb.DB)
	h := handlers.HandlerAuthentication(authenticationRepository)

	r.POST("/register", h.Register())
	r.POST("/login", h.Login())
}
