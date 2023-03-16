package main

import (
	"github.com/gin-gonic/gin"
)

var version = "dev"

func main() {
	r := gin.Default()
	r.GET("/api-sederhana", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pesan": "Masuk",
		})
	})

	r.Run()
}
