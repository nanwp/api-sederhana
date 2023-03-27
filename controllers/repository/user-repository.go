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

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(users users.User) (users.User, error) {
	err := r.db.Table("tbl_user").Create(&users).Error
	return users, err
}

func (r *repository) FindByUsername(username string) (users.User, error) {
	var user users.User
	err := r.db.Table("tbl_user").Where("username = ?", username).First(&user).Error
	return user, err
}
func (r *repository) FindByEmail(email string) (users.User, error) {
	var user users.User
	err := r.db.Table("tbl_user").Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) FindByID(id string) (users.User, error) {
	var user users.User
	err := r.db.Table("tbl_user").Where("id = ?", id).First(&user).Error
	return user, err
}
func (r *repository) FindAll() ([]users.User, error) {
	var user []users.User
	err := r.db.Table("tbl_user").Find(&user).Error
	return user, err
}
