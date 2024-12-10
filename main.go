package main

import (
	"log"
	"os"

	"github.com/heartlezz7/go_gin_todolist/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatal("Run server failed!", err)
		os.Exit(1)
	}

}
