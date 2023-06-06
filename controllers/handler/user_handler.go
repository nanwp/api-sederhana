package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nanwp/rknet/config"
	"github.com/nanwp/rknet/controllers/service"
	"github.com/nanwp/rknet/models"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	authService service.AuthService
}

func NewUserHandler(authService service.AuthService) *userHandler {
	return &userHandler{authService}
}

func (h *userHandler) Registrasi(c *gin.Context) {
	var userRequest models.UserCreate

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": errorMessages,
		})

		return
	}

	uuidGenerate := uuid.New()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	userReg := models.User{
		ID:       uuidGenerate.String(),
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: string(hashPassword),
		Role:     userRequest.Role,
	}

	userCreate, err := h.authService.Registrasi(userReg)

	if err != nil || !userCreate {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success create user",
	})
	return
}

func (h *userHandler) Login(c *gin.Context) {
	var userInput models.Login

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": errorMessages,
		})

		return
	}

	userLogin, err := h.authService.Login(userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	expTime := time.Now().Add(time.Hour * 24)
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
		"code":    200,
		"message": "berhasil login",
	})
}
