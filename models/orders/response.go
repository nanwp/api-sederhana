package orders

import (
	"time"

	"github.com/nanwp/api-sederhana/models/payments"
	"github.com/nanwp/api-sederhana/models/products"
	"github.com/nanwp/api-sederhana/models/users"
)

type OrderCreateResponse struct {
	ID          int                `json:"id"`
	UserId      int                `json:"user_id"`
	PaymentId   int                `json:"payment_id"`
	TotalPrice  int                `json:"total_price"`
	TotalPaid   int                `json:"total_paid"`
	TotalReturn int                `json:"total_return"`
	Products    []products.Product `json:"products"`
	Payment     payments.Payment   `json:"payment_type"`
	CreatedAt   time.Time          `json:"created_at"`
}

type OrderGetResponse struct {
	ID          int                      `json:"id"`
	User        users.UserResponse       `json:"user"`
	TotalPrice  int                      `json:"total_price"`
	TotalPaid   int                      `json:"total_paid"`
	TotalReturn int                      `json:"total_return"`
	Payment     payments.PaymentResponse `json:"payment_type"`
	CreatedAt   time.Time                `json:"created_at"`
}
