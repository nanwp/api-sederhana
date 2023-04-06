package repository

import (
	orderproducts "github.com/nanwp/api-sederhana/models/order_products"
	"gorm.io/gorm"
)

type OrderProductRepository interface {
	Create(orderProduct orderproducts.OrderProduct) (orderproducts.OrderProduct, error)
	FindByID(id int) (orderproducts.OrderProduct, error)
	FindAll() ([]orderproducts.OrderProduct, error)
}

type orderProductRepository struct {
	db *gorm.DB
}

func NewOrderProductRepository(db *gorm.DB) *orderProductRepository {
	return &orderProductRepository{db}
}

func (r *orderProductRepository) Create(orderProduct orderproducts.OrderProduct) (orderproducts.OrderProduct, error) {
	err := r.db.Create(&orderProduct).Error
	return orderProduct, err
}

func (r *orderProductRepository) FindByID(id int) (orderproducts.OrderProduct, error) {
	var orderProduct orderproducts.OrderProduct
	err := r.db.Preload("Order.User").Preload("Order.Payment").Preload("Products.Category").Where("tbl_order_product.id = ?", id).First(&orderProduct).Error
	return orderProduct, err
}

func (r *orderProductRepository) FindAll() ([]orderproducts.OrderProduct, error) {
	var orderProducts []orderproducts.OrderProduct
	err := r.db.Preload("Order.User").Preload("Order.Payment").Preload("Products.Category").Find(&orderProducts).Error
	return orderProducts, err
}
