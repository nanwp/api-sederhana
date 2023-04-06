package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/orders"
)

type OrderService interface {
	Create(Order orders.OrderCreate) (orders.Order, error)
	FindByID(id int) (orders.Order, error)
}

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) *orderService {
	return &orderService{repository}
}

func (s *orderService) Create(order orders.OrderCreate) (orders.Order, error) {
	orders := orders.Order{
		UserId:      order.UserId,
		PaymentId:   order.PaymentId,
		TotalPrice:  order.TotalPrice,
		TotalPaid:   order.TotalPaid,
		TotalReturn: order.TotalReturn,
	}

	newOrder, err := s.repository.Create(orders)
	return newOrder, err

}

func (s *orderService) FindByID(id int) (orders.Order, error) {
	return s.repository.FindByID(id)
}
