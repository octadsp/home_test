package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	dto "github.com/octadsp/server-home-test/dto/result"
	jwtToken "github.com/octadsp/server-home-test/pkg/jwt"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Auth Function
func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Unauthorized"})
		c.Abort()
		return
	}

	token = strings.Split(token, " ")[1]
	claims, err := jwtToken.DecodeToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("userLogin", claims)
	c.Next()
}
