package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// App is the API for all our front-end applications
type App struct {
	Router *gin.Engine
}

func (app *App) initializeRouter() {
	app.Router = gin.Default()
}

// Initialize the application
func (app *App) Initialize() {
	app.initializeRouter()
}

// Run the application
func (app *App) Run(addr string) {
	app.Router.Run(addr)
	fmt.Printf("API started at %s\n", addr)
}
