package repository

import (
	"github.com/nanwp/api-sederhana/models/products"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product products.Product) (products.Product, error)
	FindByID(id int) (products.Product, error)
	FindAll() ([]products.Product, error)
	Update(product products.Product) (products.Product, error)
	Delete(product products.Product) (products.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product products.Product) (products.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) FindByID(id int) (products.Product, error) {
	var product products.Product
	err := r.db.Preload("Category").Joins("JOIN tbl_category on tbl_category.id=tbl_product.category_id").Where("tbl_product.id = ?", id).First(&product).Error
	return product, err
}

func (r *productRepository) FindAll() ([]products.Product, error) {
	var product []products.Product
	err := r.db.Preload("Category").Joins("JOIN tbl_category on tbl_category.id=tbl_product.category_id").Find(&product).Error
	return product, err
}

func (r *productRepository) Update(product products.Product) (products.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepository) Delete(product products.Product) (products.Product, error) {
	err := r.db.Delete(product).Error
	return product, err
}
