package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanwp/api-sederhana/config"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/middleware"
	"github.com/nanwp/api-sederhana/models/users"
)

func Index(c *gin.Context) {
	// data := []map[string]interface{}{
	// 	{
	// 		"id":           1,
	// 		"nama_product": "Kemeja",
	// 		"stok":         1000,
	// 	},
	// 	{
	// 		"id":           2,
	// 		"nama_product": "Celana",
	// 		"stok":         10000,
	// 	},
	// 	{
	// 		"id":           1,
	// 		"nama_product": "Sepatu",
	// 		"stok":         500,
	// 	},
	// }

	getUser, _ := service.NewUserService(repository.NewUserRepository(config.ConnectDatabase())).FindByID(middleware.Username)
	userLog := users.UserResponse{
		Email:    getUser.Email,
		Username: getUser.Username,
		Alamat:   getUser.Alamat,
		Phone:    getUser.Phone,
	}

	c.JSON(http.StatusOK, userLog)
}
