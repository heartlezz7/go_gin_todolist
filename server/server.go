package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heartlezz7/go_gin_todolist/database"
	"github.com/heartlezz7/go_gin_todolist/model"
	"github.com/heartlezz7/go_gin_todolist/router"
)

func RunServer() error {
	var err error

	engine := gin.New()

	server := model.Server{Engine: engine}

	err = database.ConnectPostgres()
	if err != nil {
		return err
	}
	defer database.DB.Close()

	router.Init(server)

	port := os.Getenv("PORT")
	err = server.Engine.Run(":" + port)
	if err != nil {
		return err
	}

	return nil
}
