package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func generalRoute(c *gin.Engine) {

	c.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	c.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

}
