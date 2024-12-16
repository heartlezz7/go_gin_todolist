package main

import (
	"log"
	"os"

	"github.com/heartlezz7/go_gin_todolist/server"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a Go Gin application.

// @host localhost:8888
// @BasePath /

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatal("Run server failed!", err)
		os.Exit(1)
	}

}
