package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	orderproducts "github.com/nanwp/api-sederhana/models/order_products"
)

type OrderProductService interface {
	Create(orderProduct orderproducts.OrderProduct) (orderproducts.OrderProduct, error)
	FindByID(id int) (orderproducts.OrderProduct, error)
	FindAll() ([]orderproducts.OrderProduct, error)
}

type orderProductService struct {
	repository repository.OrderProductRepository
}

func NewOrderProductService(repository repository.OrderProductRepository) *orderProductService {
	return &orderProductService{repository}
}

func (s *orderProductService) Create(orderProduct orderproducts.OrderProduct) (orderproducts.OrderProduct, error) {

	newOrders, err := s.repository.Create(orderProduct)
	return newOrders, err

}

func (s *orderProductService) FindAll() ([]orderproducts.OrderProduct, error) {

	newOrders, err := s.repository.FindAll()
	return newOrders, err

}
func (s *orderProductService) FindByID(id int) (orderproducts.OrderProduct, error) {

	newOrders, err := s.repository.FindByID(id)
	return newOrders, err

}
