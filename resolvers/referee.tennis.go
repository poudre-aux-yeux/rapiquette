package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// TennisRefereeResolver resolves tennis.Referee
type TennisRefereeResolver struct {
	ref tennis.Referee
}

// ID : resolves the ID
func (r *TennisRefereeResolver) ID() graphql.ID {
	return r.ref.ID
}

// Name : resolves the Name
func (r *TennisRefereeResolver) Name() string {
	return r.ref.Name
}
