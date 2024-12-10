package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heartlezz7/go_gin_todolist/model"
)

type route func(*gin.Engine)

var router = map[string]route{"general": generalRoute}

func Init(server model.Server) {
	for key, route := range router {
		route(server.Engine)
		fmt.Println(key, "route active")
	}
}
