package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanwp/api-sederhana/config"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/controllers/service"
)

func AdminAuth(c *gin.Context) {
	userLogin, err := service.NewUserService(repository.NewUserRepository(config.ConnectDatabase())).FindByID(Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		c.Abort()
		return
	}
	if userLogin.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"gagal": "anda bukan admin",
		})
		c.Abort()
		return
	}
	c.Next()
}
