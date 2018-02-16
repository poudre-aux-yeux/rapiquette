// Package tennis contains the business models
package tennis

import (
	graphql "github.com/neelance/graphql-go"
)

// Match : metadata about a match
type Match struct {
	ID               graphql.ID   `json:"id"`
	HomePlayersLinks []graphql.ID `json:"homePlayers"`
	AwayPlayersLinks []graphql.ID `json:"awayPlayers"`
	StdLink          graphql.ID   `json:"stadium"`
	RefLink          graphql.ID   `json:"referee"`
	Date             graphql.Time `json:"date"`
	HomePlayers      []*Player
	AwayPlayers      []*Player
	Std              *Stadium
	Ref              *Referee
	Score
}

// GetType returns the type of the struct
func (s *Match) GetType() string {
	return "Match"
}

// GetType returns the type of the struct
func (s *Player) GetType() string {
	return "Player"
}

// GetType returns the type of the struct
func (s *Referee) GetType() string {
	return "TennisReferee"
}

// GetType returns the type of the struct
func (s *Stadium) GetType() string {
	return "Stadium"
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
	ID    graphql.ID `json:"id"`
	Name  string     `json:"name"`
	Image string     `json:"image"`
}

// Stadium : metadata about a stadium
type Stadium struct {
	ID         graphql.ID `json:"id"`
	GroundType string     `json:"surface"`
	Name       string     `json:"name"`
	City       string     `json:"city"`
	Image      string     `json:"image"`
}

// Referee : tennis referee
type Referee struct {
	ID    graphql.ID `json:"id"`
	Name  string     `json:"name"`
	Image string     `json:"image"`
}

// Set : a set of games
type Set struct {
	ID    graphql.ID `json:"id"`
	Games []*Game    `json:"games"`
}

// Game : Current points of a game
type Game struct {
	ID         graphql.ID `json:"id"`
	HomePoints int32      `json:"home"`
	AwayPoints int32      `json:"away"`
}
