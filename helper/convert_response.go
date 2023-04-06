package helper

import (
	"github.com/nanwp/api-sederhana/models/category"
	orderproducts "github.com/nanwp/api-sederhana/models/order_products"
	"github.com/nanwp/api-sederhana/models/orders"
	"github.com/nanwp/api-sederhana/models/payments"
	"github.com/nanwp/api-sederhana/models/products"
	"github.com/nanwp/api-sederhana/models/users"
)

func ConvertOrderToResponse(o orderproducts.OrderProduct) orderproducts.OrderResponse {

	userResponse := ConvertUserToResponse(o.Order.User)
	productResponse := ConvertProductToResponse(o.Products)
	paymentResponse := ConvertPaymentToResponse(o.Order.Payment)

	oResponse := orders.OrderGetResponse{
		ID:          o.Order.ID,
		User:        userResponse,
		TotalPrice:  o.Order.TotalPrice,
		TotalPaid:   o.Order.TotalPaid,
		TotalReturn: o.Order.TotalReturn,
		Payment:     paymentResponse,
	}
	orderResponse := orderproducts.OrderResponse{
		ID:       o.ID,
		Qty:      o.Qty,
		Order:    oResponse,
		Products: productResponse,
	}
	return orderResponse
}

func ConvertCategoryToResponse(c category.Category) category.CategoryResponse {
	categoryResponse := category.CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
	return categoryResponse
}

func ConvertProductToResponse(p products.Product) products.ProductResponse {

	categoryResponse := category.CategoryResponse{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}

	productResponse := products.ProductResponse{
		ID:       p.ID,
		SKU:      p.SKU,
		Name:     p.Name,
		Stock:    p.Stock,
		Price:    p.Price,
		Image:    p.Image,
		Category: categoryResponse,
	}
	return productResponse
}

func ConvertProductUpdateToResponse(p products.Product) products.ProductUpdateResponse {
	productResponse := products.ProductUpdateResponse{
		SKU:        p.SKU,
		Name:       p.Name,
		Stock:      p.Stock,
		Price:      p.Price,
		Image:      p.Image,
		CategoryId: p.CategoryId,
	}
	return productResponse
}

func ConvertUserToResponse(u users.User) users.UserResponse {
	userResponse := users.UserResponse{
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
		Role:     u.Role,
	}
	return userResponse
}

func ConvertPaymentToResponse(p payments.Payment) payments.PaymentResponse {
	paymentResponse := payments.PaymentResponse{
		ID:   p.ID,
		Name: p.Name,
		Type: p.Type,
		Logo: p.Logo,
	}
	return paymentResponse
}
