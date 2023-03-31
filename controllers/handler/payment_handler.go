package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/helper"
	"github.com/nanwp/api-sederhana/models/payments"
	"gorm.io/gorm"
)

type paymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h *paymentHandler) CreatePayment(c *gin.Context) {
	var paymentRequest payments.PaymentCreate

	err := c.ShouldBindJSON(&paymentRequest)
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

	payment, err := h.paymentService.Create(paymentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"add": gin.H{
			"data": payment,
		},
	})
}

func (h *paymentHandler) GetPayments(c *gin.Context) {
	payment, err := h.paymentService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var paymentsResponse []payments.PaymentResponse

	for _, pay := range payment {
		paymentResponse := helper.ConvertPaymentToResponse(pay)
		paymentsResponse = append(paymentsResponse, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"payments": paymentsResponse,
		},
	})
}

func (h *paymentHandler) GetPayment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	payment, err := h.paymentService.FindByID(id)
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

	paymentResponse := helper.ConvertPaymentToResponse(payment)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"payment": paymentResponse,
		},
	})
}

func (h *paymentHandler) UpdatePayment(c *gin.Context) {
	var paymentUpdate payments.PaymentUpdate

	err := c.ShouldBindJSON(&paymentUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	payment, err := h.paymentService.Update(id, paymentUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	paymentResponse := helper.ConvertPaymentToResponse(payment)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"payment": paymentResponse,
		},
	})
}

func (h *paymentHandler) DeletePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	payment, err := h.paymentService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	paymentResponse := helper.ConvertPaymentToResponse(payment)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"payment": paymentResponse,
		},
	})
}
