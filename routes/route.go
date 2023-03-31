package routes

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

func Route() {
	r := gin.Default()
	db := config.ConnectDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productHandler := handler.NewProductHandler(service.NewProductService(repository.NewProductRepository(db)))

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

	// r.GET("/users", userHandler.GetUsers)

	a := r.Group("/api", middleware.JWTMiddleware)
	a.GET("/users", middleware.AdminAuth, userHandler.GetUsers)
	a.POST("/addcategory", middleware.AdminAuth, categoryHandler.CreateCategory)
	a.POST("/addproduct", middleware.AdminAuth, productHandler.CreateProduct)
	a.GET("/products", productHandler.GetProduct)
	r.Run(":8080")
}
