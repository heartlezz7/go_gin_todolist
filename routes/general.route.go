package route

import (
	"github.com/gin-gonic/gin"
	"github.com/heartlezz7/go_gin_todolist/handler"
)

func generalRoute(r *gin.Engine) {

	r.GET("/ping", handler.PingHandler)

	r.GET("/health", handler.HealthCheckHandler)

	r.GET("/api/v1/hello", handler.HelleWorld)
}
