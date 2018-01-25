package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/resolvers"
	"github.com/poudre-aux-yeux/rapiquette/schema"
)

// Server is the API for all our front-end applications
type Server struct {
	Router *gin.Engine
	Schema *graphql.Schema
}

func (server *Server) initializeGraphQL() {
	var rootResolver resolvers.Resolver
	server.Schema = graphql.MustParseSchema(schema.String(), rootResolver)
}

func (server *Server) initializeRouter() {
	server.Router = gin.Default()
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
