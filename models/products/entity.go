package products

import (
	"time"

	"github.com/nanwp/api-sederhana/models/category"
)

type Product struct {
	ID         int
	SKU        string
	Name       string
	Stock      int
	Price      int
	Image      string
	CategoryId int
	Category   category.Category `gorm:"foreignKey:CategoryId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Product) TableName() string {
	return "tbl_product"
}
