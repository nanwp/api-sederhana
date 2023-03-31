package repository

import (
	"github.com/nanwp/api-sederhana/models/category"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category category.Category) (category.Category, error)
	FindByID(id int) (category.Category, error)
	FindAll() ([]category.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category category.Category) (category.Category, error) {
	err := r.db.Table("tbl_category").Create(&category).Error
	return category, err
}

func (r *categoryRepository) FindByID(id int) (category.Category, error) {
	var category category.Category
	err := r.db.Table("tbl_category").Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryRepository) FindAll() ([]category.Category, error) {
	var category []category.Category
	err := r.db.Table("tbl_category").Find(&category).Error
	return category, err
}
