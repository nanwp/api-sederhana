package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanwp/api-sederhana/config"
	"github.com/nanwp/api-sederhana/controllers/handler"
	"github.com/nanwp/api-sederhana/controllers/repository"
	"github.com/nanwp/api-sederhana/controllers/service"
	"github.com/nanwp/api-sederhana/middleware"
)

var version = "dev"

func main() {
	r := gin.Default()
	db := config.ConnectDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Use(CORSMiddleware())

	r.POST("/daftar", userHandler.Register)
	r.POST("/login", userHandler.Login)

	a := r.Group("/users", middleware.JWTMiddleware)
	a.GET("/", handler.Index)
	r.Run(":8080")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
