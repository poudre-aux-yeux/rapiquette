package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/handler"
	"github.com/poudre-aux-yeux/rapiquette/kvs"
	"github.com/poudre-aux-yeux/rapiquette/redis"
	"github.com/poudre-aux-yeux/rapiquette/resolvers"
	"github.com/poudre-aux-yeux/rapiquette/schema"
)

// App is the API for all our front-end applications
type App struct {
	Server  *http.Server
	GraphQL handler.GraphQL
	Pool    *redis.Pool
}

func (app *App) initializeGraphQL() {
	resolver, _ := resolvers.NewRoot()
	app.GraphQL.Schema = graphql.MustParseSchema(schema.String(), resolver)
}

func (app *App) initializeServer() {
	var (
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	app.Server = &http.Server{
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
}

func (app *App) initializeRoutes() {
	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/graphql", app.GraphQL)
	mux.Handle("/graphql/", app.GraphQL)

	app.Server.Handler = mux
}

func (app *App) initializeKeyValueStore() {

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = ":6379"
	}
	defer cleanupHook()
	app.Pool = kvs.NewPool(host)
}

// Intercept the exit signal and close the pool properly before exiting
func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

// Initialize the application
func (app *App) Initialize() {
	app.initializeGraphQL()
	app.initializeKeyValueStore()
	app.initializeServer()
	app.initializeRoutes()
}

// Run the application
func (app *App) Run(addr string) {
	app.Server.Addr = addr
	fmt.Printf("API started at %s\n", addr)

	if err := app.Server.ListenAndServe(); err != nil {
		log.Println("app.ListenAndServe:", err)
	}

	fmt.Println("API shut down")
}
