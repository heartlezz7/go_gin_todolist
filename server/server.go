package server

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/heartlezz7/go_gin_todolist/database"
	_ "github.com/heartlezz7/go_gin_todolist/docs"
	"github.com/heartlezz7/go_gin_todolist/model"
	route "github.com/heartlezz7/go_gin_todolist/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() error {
	var err error

	engine := gin.Default()

	server := model.Server{Engine: engine}

	server.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Specify allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},   // Specify allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Specify allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Specify headers to expose
		AllowCredentials: true,                                                // Allow credentials (cookies, authorization headers, etc.)
		MaxAge:           12 * time.Hour,                                      // Specify the duration for which preflight requests can be cached
	}))

	// connect DB
	err = database.ConnectPostgres()
	if err != nil {
		return err
	}
	defer database.DB.Close()

	route.Init(server)

	// Swagger documentation route
	server.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	err = server.Engine.Run(":" + port)
	if err != nil {
		return err
	}

	return nil
}
