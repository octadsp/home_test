package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	cartsdto "github.com/octadsp/server-home-test/dto/cart"
	dto "github.com/octadsp/server-home-test/dto/result"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/repositories"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) FindCarts(c *gin.Context) {
	carts, err := h.CartRepository.FindCarts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertFindCartsResponse(carts)})
}

func (h *handlerCart) AddCart(c *gin.Context) {
	var err error
	request := new(cartsdto.CartRequest)

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

	cart := models.Cart{
		ProductID:       request.ProductID,
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		CustomerPhone:   request.CustomerPhone,
		Price:           request.Price,
		Qty:             request.Qty,
		Status:          "pending",
	}

	cart, err = h.CartRepository.AddCart(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	cart, _ = h.CartRepository.GetCart(int(cart.ID))

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertAddCartResponse(cart)})
}

func (h *handlerCart) DeleteCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func convertFindCartsResponse(carts []models.Cart) []cartsdto.FindCartsResponse {
	responses := make([]cartsdto.FindCartsResponse, len(carts))
	for i, u := range carts {

		responses[i] = cartsdto.FindCartsResponse{
			ID:              u.ID,
			ProductID:       u.ProductID,
			CustomerName:    u.CustomerName,
			CustomerAddress: u.CustomerAddress,
			CustomerPhone:   u.CustomerPhone,
			Price:           u.Price,
			Qty:             u.Qty,
			Status:          u.Status,
		}
	}
	return responses
}

func convertAddCartResponse(cart models.Cart) cartsdto.AddCartResponse {

	response := cartsdto.AddCartResponse{
		ID:              cart.ID,
		ProductID:       cart.ProductID,
		CustomerName:    cart.CustomerName,
		CustomerAddress: cart.CustomerAddress,
		CustomerPhone:   cart.CustomerPhone,
		Price:           cart.Price,
		Qty:             cart.Qty,
		Status:          cart.Status,
	}

	return response
}
