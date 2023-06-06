package main

import (
	"fmt"

	"github.com/nanwp/rknet/config"
	"github.com/nanwp/rknet/pkg/middleware"
)

func main() {
	db := config.Connect()

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	router := middleware.InitRouter(db)

	if err := router.Run(":8080"); err != nil {
		panic(fmt.Errorf("failed to start server: %s", err))
	}
}

// ubah
