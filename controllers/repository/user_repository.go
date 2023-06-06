package repository

import (
	"github.com/nanwp/rknet/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByID(id string) (models.User, error)
	FindAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(users models.User) (models.User, error) {
	err := r.db.Create(&users).Error
	return users, err
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
func (r *userRepository) FindAll() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error
	return user, err
}
