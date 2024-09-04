package main

import (
	"gcim/example/config"
	gqlrouter "gcim/example/pkg/adapter/api/graphql/router"
	restapi "gcim/example/pkg/adapter/api/rest"
	"gcim/example/pkg/adapter/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "3000"

func main() {
	// Load .env file
	config.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Config router
	router := gin.Default()
	router.Use(middleware.ErrorHandler)

	gqlrouter.Route(router)
	restapi.Route(router)

	router.Run(":" + port)
}
