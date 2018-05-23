package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/auth"
	"github.com/poudre-aux-yeux/rapiquette/handler"
	"github.com/poudre-aux-yeux/rapiquette/kvs"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/resolvers"
	"github.com/poudre-aux-yeux/rapiquette/schema"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// App is the API for all our front-end applications
type App struct {
	Server   *http.Server
	GraphQL  handler.GraphQL
	Tennis   *tennis.Client
	Raquette *raquette.Client
}

func (app *App) initializeTennisKeyValueStore() {
	host := os.Getenv("TENNIS_HOST")
	if host == "" {
		host = ":6379"
	}
	redis := kvs.NewRedis(host)
	client, err := tennis.New(redis)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.Tennis = client
}

func (app *App) initializeRaquetteKeyValueStore() {
	host := os.Getenv("RAQUETTE_HOST")
	if host == "" {
		host = ":6379"
	}
	redis := kvs.NewRedis(host)
	client, err := raquette.New(redis)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.Raquette = client
}

func (app *App) initializeGraphQL() {
	resolver, err := resolvers.NewRoot(app.Tennis, app.Raquette)

	if err != nil {
		fmt.Println("Error creating the Query resolver")
		os.Exit(1)
	}

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
	mux.HandleFunc("/login", app.loginFunc)

	app.Server.Handler = mux
}

func (app *App) loginFunc(w http.ResponseWriter, req *http.Request) {
	var jwtSecret = "secretodelavega"
	decoder := json.NewDecoder(req.Body)
	payload := auth.LoginPayload{}
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	fmt.Println(payload)
	user, err := auth.ConfirmLogin(&payload, app.Raquette)
	if err != nil {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	expires := time.Now().Add(time.Hour * 1).Unix()
	claims := auth.Claims{
		user.ID,
		true,
		false,
		jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer:    "rapiquette",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)
	tokenResponse := struct {
		Token string `json:"token"`
	}{signedToken}
	json.NewEncoder(w).Encode(tokenResponse)
}

// Initialize the application
func (app *App) Initialize() {
	app.initializeTennisKeyValueStore()
	app.initializeRaquetteKeyValueStore()
	app.initializeGraphQL()
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
