package resolvers

import (
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

// Players : resolves the players
func (r *MatchResolver) Players() []*PlayerResolver {
	return make([]*PlayerResolver, 0)
}

// Referee : resolves the tennis Referee
func (r *MatchResolver) Referee() *TennisRefereeResolver {
	return &TennisRefereeResolver{}
}

// Stadium : resolves the Stadium
func (r *MatchResolver) Stadium() *StadiumResolver {
	return &StadiumResolver{}
}

// Sets : resolves the sets
func (r *MatchResolver) Sets() []*SetResolver {
	return make([]*SetResolver, 0)
}

// Service : resolves the service
func (r *MatchResolver) Service() *bool {
	return &r.match.Service
}
