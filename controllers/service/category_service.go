package service

import (
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/category"
)

type CategoryService interface {
	Create(category category.CategoryCreate) (category.Category, error)
	FindByID(id int) (category.Category, error)
	FindAll() ([]category.Category, error)
	Update(ID int, categoryUpdate category.CategoryUpdate) (category.Category, error)
	Delete(ID int) (category.Category, error)
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

func (s *categoryService) Update(ID int, categoryUpdate category.CategoryUpdate) (category.Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	if categoryUpdate.Name != "" {
		category.Name = categoryUpdate.Name
	}

	updateCategory, err := s.repository.Update(category)

	return updateCategory, err
}

func (s *categoryService) Delete(ID int) (category.Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	deleteCategory, err := s.repository.Delete(category)
	return deleteCategory, err
}
