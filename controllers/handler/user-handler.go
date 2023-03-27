package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nanwp/api-sederhana/config"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/models/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var userRequest users.UserCreate

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on fieled %s, conditions: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	userCheck, _ := service.NewUserService(repository.NewUserRepository(config.ConnectDatabase().Debug())).FindAll()
	for _, a := range userCheck {
		if a.Username == userRequest.Username {
			c.JSON(http.StatusBadRequest, gin.H{
				"username": a.Username,
				"error":    "username telah digunakan",
			})
			return
		}
	}
	for _, b := range userCheck {
		if b.Email == userRequest.Email {
			c.JSON(http.StatusBadRequest, gin.H{
				"email": b.Email,
				"error": "email telah digunakan",
			})
			return
		}
	}

	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) Login(c *gin.Context) {

	var userInput users.UserLogin

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on fieled %s, conditions: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	userLogin, err := h.userService.FindByUsername(userInput.Username)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			userEmail, err := h.userService.FindByEmail(userInput.Username)
			if err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.JSON(http.StatusUnauthorized, gin.H{
						"message": "Username atau password salah",
					})
					return
				default:
					c.JSON(http.StatusInternalServerError, gin.H{
						"message login": err,
					})
					return
				}
			}
			userLogin = userEmail
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message login": err,
			})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username atau password salah",
		})
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: userLogin.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "api-sederhana",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   3600 * 24 * 30,
		Path:     "/",
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil login",
	})
}
