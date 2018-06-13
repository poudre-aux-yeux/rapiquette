package resolvers

import (
	"context"

	graphql "github.com/poudre-aux-yeux/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// RaquetteRefereeResolver : resolves raquette.Referee
type RaquetteRefereeResolver struct {
	ref    *raquette.Referee
	tennis *tennis.Client
}

// ID : resolves the ID
func (r *RaquetteRefereeResolver) ID() graphql.ID {
	return r.ref.ID
}

// PasswordHash : resolves the PasswordHash
func (r *RaquetteRefereeResolver) PasswordHash() string {
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

// Ref : resolves the Referee
func (r *RaquetteRefereeResolver) Ref(ctx context.Context) (*TennisRefereeResolver, error) {
	if r.ref.Ref.ID != "" {
		return &TennisRefereeResolver{ref: r.ref.Ref}, nil
	}

	ref, err := r.tennis.GetRefereeByID(r.ref.RefLink)
	return &TennisRefereeResolver{ref: ref}, err
}
