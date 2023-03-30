package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/category"
)

type CategoryService interface {
	Create(category category.CategoryCreate) (category.Category, error)
	FindByID(id int) (category.Category, error)
	FindAll() ([]category.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *categoryService {
	return &categoryService{repository}
}

func (s *categoryService) Create(categorys category.CategoryCreate) (category.Category, error) {
	category := category.Category{
		Name: categorys.Name,
	}

	newCategory, err := s.repository.Create(category)

	return newCategory, err
}

func (s *categoryService) FindByID(id int) (category.Category, error) {
	return s.repository.FindByID(id)
}

func (s *categoryService) FindAll() ([]category.Category, error) {
	return s.repository.FindAll()
}
