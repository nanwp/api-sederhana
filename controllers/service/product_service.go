package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/products"
)

type ProductService interface {
	Create(product products.ProductCreate) (products.Product, error)
	FindByID(id int) (products.Product, error)
	FindAll() ([]products.Product, error)
	Update(ID int, productUpdate products.ProductUpdate) (products.Product, error)
	Delete(ID int) (products.Product, error)
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

func (s *productService) Update(ID int, productUpdate products.ProductUpdate) (products.Product, error) {
	produk, err := s.repository.FindByID(ID)
	if err != nil {
		return produk, err
	}

	if productUpdate.SKU != "" {
		produk.SKU = productUpdate.SKU
	}

	if productUpdate.Name != "" {
		produk.Name = productUpdate.Name
	}

	if productUpdate.Stock != 0 {
		produk.Stock = productUpdate.Stock
	}

	if productUpdate.Price != 0 {
		produk.Price = productUpdate.Price
	}

	if productUpdate.Image != "" {
		produk.Image = productUpdate.Image
	}

	if productUpdate.CategoryId != 0 {
		produk.Category.ID = productUpdate.CategoryId
	}

	updateProduct, err := s.repository.Update(produk)

	return updateProduct, err
}

func (s *productService) Delete(ID int) (products.Product, error) {
	produk, err := s.repository.FindByID(ID)
	if err != nil {
		return produk, err
	}

	deleteProduk, err := s.repository.Delete(produk)
	return deleteProduk, err
}
