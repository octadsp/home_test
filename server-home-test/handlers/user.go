package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *handlerUser) FindUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.UserRepository.FindUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Status": http.StatusOK,
			"Data":   convertFindUsersResponse(users),
		})
	}
}

func convertFindUsersResponse(users []models.User) []usersdto.UserResponse {
	responses := make([]usersdto.UserResponse, len(users))
	for i, u := range users {
		responses[i] = usersdto.UserResponse{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
			Role:  u.Role,
		}
	}
	return responses
}
