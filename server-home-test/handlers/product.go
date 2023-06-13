package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	productsdto "github.com/octadsp/server-home-test/dto/product"
	dto "github.com/octadsp/server-home-test/dto/result"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/repositories"
)

var path_file = "http://localhost:5000/api/v1/uploads"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(productRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{productRepository}
}

func (h *handlerProduct) FindProducts(c *gin.Context) {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertFindProductsResponse(products)})
}

func (h *handlerProduct) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	product.Image = path_file + product.Image

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertGetProductResponse(product)})
}

func (h *handlerProduct) AddProduct(c *gin.Context) {
	var err error
	dataFile := c.GetString("dataFile")

	price, _ := strconv.Atoi(c.PostForm("price"))
	qty, _ := strconv.Atoi(c.PostForm("qty"))

	request := productsdto.ProductRequest{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Price:       price,
		Qty:         qty,
		Image:       dataFile,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	userLogin, _ := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Qty:         request.Qty,
		Image:       request.Image,
		UserID:      uint(userId),
	}

	product, err = h.ProductRepository.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	product, _ = h.ProductRepository.GetProduct(int(product.ID))

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertAddProductResponse(product)})
}

func (h *handlerProduct) UpdateProduct(c *gin.Context) {
	var err error
	dataFile := c.GetString("dataFile")

	price, _ := strconv.Atoi(c.PostForm("price"))
	qty, _ := strconv.Atoi(c.PostForm("qty"))

	request := productsdto.UpdateProductRequest{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Price:       price,
		Image:       dataFile,
		Qty:         qty,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.Image != "" {
		product.Image = request.Image
	}

	if request.Qty != 0 {
		product.Qty = request.Qty
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerProduct) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func convertFindProductsResponse(products []models.Product) []productsdto.FindProductsResponse {
	responses := make([]productsdto.FindProductsResponse, len(products))
	for i, u := range products {

		responses[i] = productsdto.FindProductsResponse{
			ID:          u.ID,
			Name:        u.Name,
			Description: u.Description,
			Price:       u.Price,
			Qty:         u.Qty,
			Image:       u.Image,
			UserID:      u.UserID,
		}
	}
	return responses
}

func convertGetProductResponse(product models.Product) productsdto.GetProductResponse {

	response := productsdto.GetProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
		Image:       product.Image,
		UserID:      product.UserID,
	}

	return response
}

func convertAddProductResponse(product models.Product) productsdto.AddProductResponse {

	response := productsdto.AddProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
		Image:       product.Image,
		UserID:      product.UserID,
	}

	return response
}
