package route

import (
	"github.com/gin-gonic/gin"
	"github.com/heartlezz7/go_gin_todolist/handler"
)

func AuthRoute(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/register", handler.Register)
	}
}
