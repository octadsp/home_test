package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	productsdto "github.com/octadsp/server-home-test/dto/product"
	dto "github.com/octadsp/server-home-test/dto/result"
	"github.com/octadsp/server-home-test/models"
	"github.com/octadsp/server-home-test/repositories"
)

var path_file = "http://localhost:5000/api/v1/uploads"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
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

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: products})
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

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: product})
}

func (h *handlerProduct) AddProduct(c *gin.Context) {
	var err error
	dataFile := c.GetString("dataFile")

	price, _ := strconv.Atoi(c.PostForm("price"))
	qty, _ := strconv.Atoi(c.PostForm("qty"))

	request := productsdto.UpdateProductRequest{
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

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Qty:         request.Qty,
		Image:       request.Image,
	}

	product, err = h.ProductRepository.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	product, _ = h.ProductRepository.GetProduct(int(product.ID))

	c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: product})
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
