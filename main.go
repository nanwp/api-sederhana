package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jub0bs/fcors"
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

	cors, err := fcors.AllowAccess(
		fcors.FromAnyOrigin(),
		fcors.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			"UPDATE",
		),
		fcors.WithRequestHeaders(
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Max",
		),
		fcors.MaxAgeInSeconds(86400),
	)
	if err != nil {
		log.Fatal(err)
	}
	r.Use(adapter.Wrap(cors))

	r.POST("/daftar", userHandler.Register)
	r.POST("/login", userHandler.Login)

	a := r.Group("/users", middleware.JWTMiddleware)
	a.GET("/", handler.Index)
	r.Run(":8080")

}
