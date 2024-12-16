package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler 	godoc
// @Summary 	Get a pong message
// @Description Returns a simple pong message
// @Tags 		tags
// @Accept  	json
// @Produce  	json
// @response 	200 {string} string "OK"
// @Router 		/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})

}

// HealthCheckHandler 	godoc
// @Summary 			Health Check
// @Description 		Health checking for the service
// @Tags 				tags
// @Accept  			json
// @Produce  			json
// @response 			200 {string} string "OK"
// @Router 				/health [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// @Summary 	Get a hello world message
// @Description Returns a simple hello world message
// @Tags 		tags
// @Accept  	json
// @Produce  	json
// @response 	200 {string} string "OK"
// @Router 		/api/v1/hello [get]
func HelleWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
