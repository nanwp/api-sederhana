package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nanwp/rknet/controllers/handler"
	"github.com/nanwp/rknet/controllers/repository"
	"github.com/nanwp/rknet/controllers/service"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(authService)

	r := router.Group("api")

	r.POST("/register", userHandler.Registrasi)
	r.POST("/login", userHandler.Login)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Page not found",
		})
	})

	return router

}
