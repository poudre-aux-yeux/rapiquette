package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// PlayerResolver : resolves tennis.Player
type PlayerResolver struct {
	player tennis.Player
}

// ID : resolves the ID
func (r *PlayerResolver) ID() graphql.ID {
	return r.player.ID
}

// Name : resolves the Name
func (r *PlayerResolver) Name() string {
	return r.player.Name
}

// Image : URL of the image
func (r *PlayerResolver) Image() string {
	return r.player.Image
}
