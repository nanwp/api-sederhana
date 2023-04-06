package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nanwp/api-sederhana/config"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/helper"
	"github.com/nanwp/api-sederhana/middleware"
	orderproducts "github.com/nanwp/api-sederhana/models/order_products"
	"github.com/nanwp/api-sederhana/models/orders"
	"github.com/nanwp/api-sederhana/models/products"
	"gorm.io/gorm"
)

type orderHandelr struct {
	orderService service.OrderProductService
}

func NewOrederHandler(orderService service.OrderProductService) *orderHandelr {
	return &orderHandelr{orderService}
}

func (h *orderHandelr) CreateOrder(c *gin.Context) {

	var orderCreate orderproducts.OrderProductCreate

	err := c.ShouldBindJSON(&orderCreate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	productt, err := service.NewProductService(repository.NewProductRepository(config.ConnectDatabase())).FindByID(orderCreate.ProductId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Product not found",
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}

	if productt.Stock < orderCreate.Qty {
		c.JSON(http.StatusBadRequest, gin.H{
			"gagal": "stok tidak cukup",
		})
		return
	}

	productUp := products.ProductUpdate{
		Stock: productt.Stock - orderCreate.Qty,
	}
	productUpdate, err := service.NewProductService(repository.NewProductRepository(config.ConnectDatabase())).Update(orderCreate.ProductId, productUp)

	payment, err := service.NewPaymentService(repository.NewPaymentRepository(config.ConnectDatabase())).FindByID(orderCreate.PaymentId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Product not found",
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}

	totalPrice := productt.Price * orderCreate.Qty

	userId1 := middleware.Username

	orderReqCreate := orders.OrderCreate{
		UserId:      userId1,
		PaymentId:   orderCreate.PaymentId,
		TotalPrice:  totalPrice,
		TotalPaid:   orderCreate.TotalPaid,
		TotalReturn: orderCreate.TotalPaid - totalPrice,
	}

	if orderCreate.TotalPaid < totalPrice {
		c.JSON(http.StatusBadRequest, gin.H{
			"gagal": "uang kurang",
		})
		return
	}

	order := service.NewOrderService(repository.NewOrderRepository(config.ConnectDatabase()))
	orderApp, err := order.Create(orderReqCreate)

	orderProduct := orderproducts.OrderProduct{
		OrderId:    orderApp.ID,
		ProductId:  orderCreate.ProductId,
		Qty:        orderCreate.Qty,
		TotalPrice: totalPrice,
	}

	orderSukses, err := h.orderService.Create(orderProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	responseOrder := orderproducts.OrderCreateResponse{
		ID:          orderSukses.ID,
		UserId:      userId1,
		PaymentId:   orderCreate.PaymentId,
		TotalPrice:  totalPrice,
		TotalPaid:   orderCreate.TotalPaid,
		TotalReturn: orderCreate.TotalPaid - totalPrice,
		Products:    productUpdate,
		Payment:     payment,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"add": gin.H{
			"data": responseOrder,
		},
	})
}

func (h *orderHandelr) GetOrders(c *gin.Context) {
	order, err := h.orderService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var orderResponse []orderproducts.OrderResponse

	for _, p := range order {
		produk := helper.ConvertOrderToResponse(p)
		orderResponse = append(orderResponse, produk)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"orders": orderResponse,
		},
	})
}

func (h *orderHandelr) GetOrder(c *gin.Context) {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	order, err := h.orderService.FindByID(idInt)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "order not found",
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data": gin.H{
			"order": helper.ConvertOrderToResponse(order),
		},
	})
}
