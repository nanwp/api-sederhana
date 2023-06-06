package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	url := "postgres://postgres:NewVPSNanda@103.161.184.72:5432/rk-net"
	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return database
}
