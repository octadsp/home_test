package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	jwtToken "github.com/octadsp/server-home-test/pkg/jwt"
	authenticationdto "github.com/octadsp/server-home-test/dto/authentication"
	dto "github.com/octadsp/server-home-test/dto/result"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/pkg/bcrypt"
	repository "github.com/octadsp/server-home-test/repositories"
)

type handlerAuthentication struct {
	AuthenticationRepository repository.AuthenticationRepository
}

func HandlerAuthentication(AuthenticationRepository repository.AuthenticationRepository) *handlerAuthentication {
	return &handlerAuthentication{AuthenticationRepository}
}

func (h *handlerAuthentication) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(authenticationdto.RegisterRequest)

		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		validation := validator.New()
		if err := validation.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		password, err := bcrypt.HashPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		user := models.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: password,
			Role:     request.Role,
		}

		data, err := h.AuthenticationRepository.Register(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		registerResponse := authenticationdto.RegisterResponse{
			Name:  data.Name,
			Email: data.Email,
			Role:  data.Role,
		}
		c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: registerResponse})
	}
}

func (h *handlerAuthentication) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(authenticationdto.LoginRequest)
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		user := models.User{
			Email:    request.Email,
			Password: request.Password,
		}

		// Check email
		user, err := h.AuthenticationRepository.Login(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		// Check password
		isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
		if !isValid {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "wrong email or password"})
			return
		}

		// Generate token
		claims := jwt.MapClaims{}
		claims["id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours expired

		token, errGenerateToken := jwtToken.GenerateToken(&claims)
		if errGenerateToken != nil {
			log.Println(errGenerateToken)
			c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized"})
			return
		}

		loginResponse := authenticationdto.LoginResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Token: token,
			Role:  user.Role,
		}

		c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: loginResponse})
	}
}
