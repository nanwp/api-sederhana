package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/helper"
	"github.com/nanwp/api-sederhana/models/category"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {

	var categoryRequest category.CategoryCreate

	err := c.ShouldBindJSON(&categoryRequest)
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

	category, err := h.categoryService.Create(categoryRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"add ": category.Name,
	})
}

func (h *categoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var categoriesResponse []category.CategoryResponse

	for _, cat := range categories {
		categoryResponse := helper.ConvertCategoryToResponse(cat)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesResponse,
	})
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	kategori, err := h.categoryService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	categoryResponse := helper.ConvertCategoryToResponse(kategori)

	c.JSON(http.StatusOK, gin.H{
		"data": categoryResponse,
	})
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var categoryUpdate category.CategoryUpdate

	err := c.ShouldBindJSON(&categoryUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	kategori, err := h.categoryService.Update(id, categoryUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"date": helper.ConvertCategoryToResponse(kategori),
	})

}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	kategori, err := h.categoryService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": helper.ConvertCategoryToResponse(kategori),
	})
}
