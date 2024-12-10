package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.Default()
	server := gin.New()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	port := os.Getenv("PORT")
	server.Run(":" + port)

}
