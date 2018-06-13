package resolvers

import (
	graphql "github.com/poudre-aux-yeux/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// GameResolver : resolves tennis.Game
type GameResolver struct {
	game tennis.Game
}

// ID : resolves the ID
func (r *GameResolver) ID() graphql.ID {
	return r.game.ID
}

// HomePoints : resolves the HomePoints
func (r *GameResolver) HomePoints() *int32 {
	return &r.game.HomePoints
}

// AwayPoints : resolves the AwayPoints
func (r *GameResolver) AwayPoints() *int32 {
	return &r.game.AwayPoints
}
