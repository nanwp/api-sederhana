package orders

import (
	"time"

	"github.com/nanwp/api-sederhana/models/payments"
	"github.com/nanwp/api-sederhana/models/users"
)

type Order struct {
	ID          int
	UserId      string
	User        users.User `gorm:"foreignKey:UserId"`
	PaymentId   int
	Payment     payments.Payment `gorm:"foreignKey:PaymentId"`
	TotalPrice  int
	TotalPaid   int
	TotalReturn int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Order) TableName() string {
	return "tbl_order"
}
