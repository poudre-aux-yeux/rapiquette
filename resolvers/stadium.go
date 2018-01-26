package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// StadiumResolver : resolves tennis.Stadium
type StadiumResolver struct {
	stadium tennis.Stadium
}

// ID : resolves the ID
func (r *StadiumResolver) ID() graphql.ID {
	return r.stadium.ID
}

// Name : resolves the Name
func (r *StadiumResolver) Name() string {
	return r.stadium.Name
}

// City : resolves the City
func (r *StadiumResolver) City() string {
	return r.stadium.City
}
