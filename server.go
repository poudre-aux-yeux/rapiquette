package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/handler"
	"github.com/poudre-aux-yeux/rapiquette/resolvers"
	"github.com/poudre-aux-yeux/rapiquette/schema"
)

// Server is the API for all our front-end applications
type Server struct {
	Router  *gin.Engine
	GraphQL handler.GraphQL
}

func (server *Server) initializeGraphQL() {
	resolver, _ := resolvers.NewRoot()
	server.GraphQL.Schema = graphql.MustParseSchema(schema.String(), resolver)
}

func (server *Server) initializeRouter() {
	server.Router = gin.Default()
}

func (server *Server) initializeRoutes() {
	server.Router.GET("/graphql")
}

// Initialize the application
func (server *Server) Initialize() {
	server.initializeGraphQL()
	server.initializeRouter()
}

// Run the application
func (server *Server) Run(addr string) {
	server.Router.Run(addr)
	fmt.Printf("API started at %s\n", addr)
}
