package service

import (
	"errors"

	"github.com/nanwp/rknet/controllers/repository"
	"github.com/nanwp/rknet/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(user models.Login) (models.User, error)
	Registrasi(user models.User) (bool, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authService {
	return &authService{repository}
}

func (s *authService) Registrasi(user models.User) (bool, error) {
	checkUsername, _ := s.userRepo.FindAll()
	for _, a := range checkUsername {
		if a.Username == user.Username {
			return false, errors.New("username sudah ada")
		}
	}

	user, err := s.userRepo.Create(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *authService) Login(user models.Login) (models.User, error) {

	usrData, err := s.userRepo.FindByUsername(user.Username)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			err := errors.New("User Tidak Ada")
			return models.User{}, err
		default:
			return models.User{}, err
		}

	}

	if err := bcrypt.CompareHashAndPassword([]byte(usrData.Password), []byte(user.Password)); err != nil {
		err := errors.New("Password Salah")
		return models.User{}, err
	}

	return usrData, nil
}
