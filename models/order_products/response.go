package orderproducts

import (
	"time"

	"github.com/nanwp/api-sederhana/models/orders"
	"github.com/nanwp/api-sederhana/models/payments"
	"github.com/nanwp/api-sederhana/models/products"
)

type OrderCreateResponse struct {
	ID          int              `json:"id"`
	UserId      string           `json:"user_id"`
	PaymentId   int              `json:"payment_id"`
	TotalPrice  int              `json:"total_price"`
	TotalPaid   int              `json:"total_paid"`
	TotalReturn int              `json:"total_return"`
	Products    products.Product `json:"product"`
	Payment     payments.Payment `json:"payment" gorm:"foreignKey:PaymentId"`
}

type OrderResponse struct {
	ID        int                      `json:"id"`
	Qty       int                      `json:"qty"`
	Products  products.ProductResponse `json:"product"`
	Order     orders.OrderGetResponse  `json:"order_detail"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
