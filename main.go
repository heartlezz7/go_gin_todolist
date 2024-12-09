package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.Default()
	server := gin.New()

	port := os.Getenv("PORT")
	server.Run(":" + port)

}
