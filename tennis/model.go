// Package tennis contains the business models
package tennis

import (
	"time"

	graphql "github.com/neelance/graphql-go"
)

// Match : metadata about a match
type Match struct {
	ID      graphql.ID `json:"id"`
	Players []*Player  `json:"players"`
	Std     *Stadium   `json:"stadium"`
	Ref     *Referee   `json:"referee"`
	Date    time.Time  `json:"date"`
	Score
}

// Score : current score of a match
type Score struct {
	Sets []*Set `json:"sets"`
	// Service :
	//   true => service to the team 1
	//   false => service to the team 2
	Service bool `json:"service"`
}

// Player : tennis player
type Player struct {
	ID   graphql.ID `json:"id"`
	Name string     `json:"name"`
}

// Stadium : metadata about a stadium
type Stadium struct {
	GroundType string `json:"surface"`
	Name       string `json:"name"`
	City       string `json:"city"`
}

// Referee : tennis referee
type Referee struct {
	ID   graphql.ID `json:"id"`
	Name string     `json:"name"`
}

// Set : a set of games
type Set struct {
	Games []*Game `json:"games"`
}

// Game : Current points of a game
type Game struct {
	HomePoints int `json:"home"`
	AwayPoints int `json:"away"`
}
