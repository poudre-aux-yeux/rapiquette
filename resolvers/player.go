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

// Birth : date of birth
func (r *PlayerResolver) Birth() graphql.Time {
	return r.player.Birth
}

func (r *PlayerResolver) Nationality() string {
	return r.player.Nationality
}

// Weight : weight in kilograms
func (r *PlayerResolver) Weight() *int32 {
	return r.player.Weight
}

// Ranking : ATP ranking
func (r *PlayerResolver) Ranking() *int32 {
	return r.player.Ranking
}

// Titles : number of titles won
func (r *PlayerResolver) Titles() *int32 {
	return r.player.Titles
}

// Height :height in centimeters
func (r *PlayerResolver) Height() *int32 {
	return r.player.Height
}
