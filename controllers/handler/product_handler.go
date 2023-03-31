package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/helper"
	"github.com/nanwp/api-sederhana/models/products"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var productRequest products.ProductCreate

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditions: %s", e.Field(), e.Error())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	product, err := h.productService.Create(productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	response := products.ProductCreateResponse{
		SKU:        product.SKU,
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
		Image:      product.Image,
		CategoryId: product.CategoryId,
	}

	c.JSON(http.StatusOK, gin.H{
		"add": response,
	})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	allProduct, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors 1": err,
		})
		return
	}

	var productsResponse []products.ProductResponse

	for _, p := range allProduct {
		pr := helper.ConvertProductToResponse(p)
		productsResponse = append(productsResponse, pr)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}
