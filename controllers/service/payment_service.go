package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/payments"
)

type PaymentService interface {
	Create(payment payments.PaymentCreate) (payments.Payment, error)
	FindByID(id int) (payments.Payment, error)
	FindAll() ([]payments.Payment, error)
	Update(id int, paymentUpdate payments.PaymentUpdate) (payments.Payment, error)
	Delete(id int) (payments.Payment, error)
}

type paymentService struct {
	repository repository.PaymentRepository
}

func NewPaymentService(repository repository.PaymentRepository) *paymentService {
	return &paymentService{repository}
}

func (s *paymentService) Create(payment payments.PaymentCreate) (payments.Payment, error) {
	payments := payments.Payment{
		Name: payment.Name,
		Type: payment.Type,
		Logo: payment.Logo,
	}
	newPayment, err := s.repository.Create(payments)
	return newPayment, err
}

func (s *paymentService) FindByID(id int) (payments.Payment, error) {
	return s.repository.FindByID(id)
}

func (s *paymentService) FindAll() ([]payments.Payment, error) {
	return s.repository.FindAll()
}

func (s *paymentService) Update(id int, paymentUpdate payments.PaymentUpdate) (payments.Payment, error) {
	payment, err := s.repository.FindByID(id)
	if err != nil {
		return payment, err
	}

	if paymentUpdate.Name != "" {
		payment.Name = paymentUpdate.Name
	}

	if paymentUpdate.Type != "" {
		payment.Type = paymentUpdate.Name
	}

	if paymentUpdate.Logo != "" {
		payment.Logo = paymentUpdate.Logo
	}

	return s.repository.Update(payment)
}

func (s *paymentService) Delete(id int) (payments.Payment, error) {
	payment, err := s.repository.FindByID(id)
	if err != nil {
		return payment, err
	}
	return s.repository.Delete(payment)
}
