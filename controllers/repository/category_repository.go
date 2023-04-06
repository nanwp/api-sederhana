package repository

import (
	"github.com/nanwp/api-sederhana/models/category"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category category.Category) (category.Category, error)
	FindByID(id int) (category.Category, error)
	FindAll() ([]category.Category, error)
	Update(category category.Category) (category.Category, error)
	Delete(category category.Category) (category.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category category.Category) (category.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) FindByID(id int) (category.Category, error) {
	var category category.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryRepository) FindAll() ([]category.Category, error) {
	var category []category.Category
	err := r.db.Find(&category).Error
	return category, err
}

func (r *categoryRepository) Update(category category.Category) (category.Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *categoryRepository) Delete(category category.Category) (category.Category, error) {
	err := r.db.Delete(category).Error
	return category, err
}
