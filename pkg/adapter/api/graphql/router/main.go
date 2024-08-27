package gqlrouter

import (
	"gcim/example/pkg/adapter/api/graphql"
	gqlresolver "gcim/example/pkg/adapter/api/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.GET("/playground", playgroundHandler())
	router.POST("/query", graphqlHandler())
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &gqlresolver.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
