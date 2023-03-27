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

	r.POST("/daftar", userHandler.Register)
	r.POST("/login", userHandler.Login)

	a := r.Group("/user", middleware.JWTMiddleware)
	a.GET("/", handler.Index)
	r.Run()

}
