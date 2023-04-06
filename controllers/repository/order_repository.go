package repository

import (
	"github.com/nanwp/api-sederhana/models/orders"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order orders.Order) (orders.Order, error)
	FindByID(id int) (orders.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order orders.Order) (orders.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) FindByID(id int) (orders.Order, error) {
	var order orders.Order
	err := r.db.Preload("Payment").Preload("User").Where("tbl_order.id = ?", id).First(&order).Error
	return order, err
}
