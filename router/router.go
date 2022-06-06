package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graph/gqlgen-todos/graph"
	"github.com/graph/gqlgen-todos/graph/generated"
	"github.com/graph/gqlgen-todos/router/context"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RouterSetup() {
	r := gin.Default()

	r.Use(cors.Default()) // 允许所有来源

	r.Use(context.GinContextToContextMiddleware())

	r.GET("/", playgroundHandler())
	r.POST("/graphql", graphqlHandler())
	r.Run()
}
