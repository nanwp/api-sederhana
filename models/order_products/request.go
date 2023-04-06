package orderproducts

import "github.com/nanwp/api-sederhana/models/products"

type OrderProductCreate struct {
	PaymentId int              `json:"payment_id" binding:"required"`
	TotalPaid int              `json:"total_paid" binding:"required"`
	ProductId int              `json:"product_id"`
	Products  products.Product `gorm:"foreignKey:ProductId"`
	Qty       int              `json:"qty" binding:"required"`
}
