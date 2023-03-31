package repository

import (
	"github.com/nanwp/api-sederhana/models/payments"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment payments.Payment) (payments.Payment, error)
	FindByID(id int) (payments.Payment, error)
	FindAll() ([]payments.Payment, error)
	Update(payment payments.Payment) (payments.Payment, error)
	Delete(payment payments.Payment) (payments.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment payments.Payment) (payments.Payment, error) {
	err := r.db.Create(&payment).Error
	return payment, err
}

func (r *paymentRepository) FindAll() ([]payments.Payment, error) {
	var payment []payments.Payment
	err := r.db.Find(&payment).Error
	return payment, err
}

func (r *paymentRepository) FindByID(id int) (payments.Payment, error) {
	var payment payments.Payment
	err := r.db.Where("id = ?", id).First(&payment).Error
	return payment, err
}

func (r *paymentRepository) Update(payment payments.Payment) (payments.Payment, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *paymentRepository) Delete(paymment payments.Payment) (payments.Payment, error) {
	err := r.db.Delete(paymment).Error
	return paymment, err
}
