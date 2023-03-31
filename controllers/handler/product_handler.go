package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/helper"
	"github.com/nanwp/api-sederhana/models/products"
	"gorm.io/gorm"
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

func (h *productHandler) GetProducts(c *gin.Context) {
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

func (h *productHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	produk, err := h.productService.FindByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "record not found",
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}

	productResponse := helper.ConvertProductToResponse(produk)

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var productUpdate products.ProductUpdate

	err := c.ShouldBindJSON(&productUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	produk, err := h.productService.Update(id, productUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": helper.ConvertProductUpdateToResponse(produk),
	})
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	produk, err := h.productService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": helper.ConvertProductToResponse(produk),
	})
}
