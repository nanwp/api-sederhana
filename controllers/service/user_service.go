package service

import (
	"github.com/google/uuid"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/models/users"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(users users.UserCreate) (users.User, error)
	FindByUsername(username string) (users.User, error)
	FindByEmail(email string) (users.User, error)
	FindByID(id string) (users.User, error)
	FindAll() ([]users.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Create(user users.UserCreate) (users.User, error) {

	uuidGenerate := uuid.New()
	stringUuid := uuidGenerate.String()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	usr := users.User{
		ID:       stringUuid,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: string(hashPassword),
		Role:     user.Role,
	}

	newUser, err := s.repository.Create(usr)

	return newUser, err
}

func (s *userService) FindByUsername(username string) (users.User, error) {
	return s.repository.FindByUsername(username)
}
func (s *userService) FindByEmail(email string) (users.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *userService) FindByID(id string) (users.User, error) {
	return s.repository.FindByID(id)
}
func (s *userService) FindAll() ([]users.User, error) {
	return s.repository.FindAll()
}
