package orderproducts

import (
	"time"

	"github.com/nanwp/api-sederhana/models/orders"
	"github.com/nanwp/api-sederhana/models/products"
)

type OrderProduct struct {
	ID         int
	OrderId    int
	Order      orders.Order     `gorm:"foreignKey:OrderId"`
	ProductId  int              `json:"product_id"`
	Products   products.Product `gorm:"foreignKey:ProductId"`
	Qty        int              `json:"qty" binding:"required"`
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


func (OrderProduct) TableName() string {
	return "tbl_order_product"
}
