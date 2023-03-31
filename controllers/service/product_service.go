package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/products"
)

type ProductService interface {
	Create(product products.ProductCreate) (products.Product, error)
	FindByID(id int) (products.Product, error)
	FindAll() ([]products.Product, error)
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{repository}
}

func (s *productService) Create(product products.ProductCreate) (products.Product, error) {

	prdct := products.Product{
		SKU:        product.SKU,
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
		Image:      product.Image,
		CategoryId: product.CategoryId,
	}

	newProduct, err := s.repository.Create(prdct)
	return newProduct, err
}

func (s *productService) FindByID(id int) (products.Product, error) {
	return s.repository.FindByID(id)
}

func (s *productService) FindAll() ([]products.Product, error) {
	return s.repository.FindAll()
}
