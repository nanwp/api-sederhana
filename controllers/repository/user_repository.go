package repository

import (
	"github.com/nanwp/api-sederhana/models/users"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user users.User) (users.User, error)
	FindByUsername(username string) (users.User, error)
	FindByEmail(email string) (users.User, error)
	FindByID(id string) (users.User, error)
	FindAll() ([]users.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(users users.User) (users.User, error) {
	err := r.db.Create(&users).Error
	return users, err
}

func (r *userRepository) FindByUsername(username string) (users.User, error) {
	var user users.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}
func (r *userRepository) FindByEmail(email string) (users.User, error) {
	var user users.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id string) (users.User, error) {
	var user users.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
func (r *userRepository) FindAll() ([]users.User, error) {
	var user []users.User
	err := r.db.Find(&user).Error
	return user, err
}
