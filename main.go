package main

import (
	gqlrouter "gcim/example/pkg/adapter/api/graphql/router"
	restapi "gcim/example/pkg/adapter/api/rest"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()
	gqlrouter.Route(router)
	restapi.Route(router)

	router.Run(":" + port)
}
