package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// url := "postgres://postgres:Latihan#@103.171.182.206:5432/api-sederhana"
	url := "postgres://postgres:r00tLatihan@103.171.182.206:5432/api-sederhana"
	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return database

}