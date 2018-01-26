package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
)

// RaquetteRefereeResolver : resolves raquette.Referee
type RaquetteRefereeResolver struct {
	ref raquette.Referee
}

// ID : resolves the ID
func (r *RaquetteRefereeResolver) ID() graphql.ID {
	return r.ref.ID
}

// Hash : resolves the Hash
func (r *RaquetteRefereeResolver) Hash() string {
	return r.ref.PasswordHash
}

// Username : resolves the Username
func (r *RaquetteRefereeResolver) Username() string {
	return r.ref.Username
}

// Email : resolves the Email
func (r *RaquetteRefereeResolver) Email() string {
	return r.ref.Email
}
