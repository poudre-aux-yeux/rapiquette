package resolvers

import (
	"context"
	"fmt"

	"github.com/poudre-aux-yeux/rapiquette/tennis"

	graphql "github.com/neelance/graphql-go"
)

// MatchResolver : resolves tennis.Match
type MatchResolver struct {
	match tennis.Match
}

// ID : resolves the ID
func (r *MatchResolver) ID() graphql.ID {
	return r.match.ID
}

// Date : resolves the Date
func (r *MatchResolver) Date() graphql.Time {
	return r.match.Date
}

// HomePlayers : resolves the home players
func (r *MatchResolver) HomePlayers(ctx context.Context) ([]*PlayerResolver, error) {
	for _, id := range r.match.HomePlayersLinks {
		fmt.Print(id)
	}

	return make([]*PlayerResolver, 0), nil
}

// AwayPlayers : resolves the home players
func (r *MatchResolver) AwayPlayers(ctx context.Context) ([]*PlayerResolver, error) {
	for _, id := range r.match.AwayPlayersLinks {
		fmt.Print(id)
	}

	return make([]*PlayerResolver, 0), nil
}

// Referee : resolves the tennis Referee
func (r *MatchResolver) Referee() *TennisRefereeResolver {
	// TODO get the ref
	return &TennisRefereeResolver{}
}

// Stadium : resolves the Stadium
func (r *MatchResolver) Stadium() *StadiumResolver {
	// TODO get the stadium
	return &StadiumResolver{}
}

// Sets : resolves the sets
func (r *MatchResolver) Sets() []*SetResolver {
	// Todo get the sets
	return make([]*SetResolver, 0)
}

// Service : resolves the service
func (r *MatchResolver) Service() *bool {
	return &r.match.Service
}
