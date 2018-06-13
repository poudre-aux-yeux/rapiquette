package resolvers

import (
	"context"
	"fmt"

	"github.com/poudre-aux-yeux/rapiquette/tennis"

	graphql "github.com/poudre-aux-yeux/graphql-go"
)

// MatchResolver : resolves tennis.Match
type MatchResolver struct {
	match  *tennis.Match
	tennis *tennis.Client
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
	players := make([]*PlayerResolver, 0)

	for _, id := range r.match.HomePlayersLinks {
		p, err := r.tennis.GetPlayerByID(id)

		if err != nil {
			return nil, fmt.Errorf("couldn't get the home player %s: %v", id, err)
		}

		players = append(players, &PlayerResolver{player: p})
	}

	return players, nil
}

// AwayPlayers : resolves the home players
func (r *MatchResolver) AwayPlayers(ctx context.Context) ([]*PlayerResolver, error) {
	players := make([]*PlayerResolver, 0)

	for _, id := range r.match.AwayPlayersLinks {
		p, err := r.tennis.GetPlayerByID(id)

		if err != nil {
			return nil, fmt.Errorf("couldn't get the away player %s: %v", id, err)
		}

		players = append(players, &PlayerResolver{player: p})
	}

	return players, nil
}

// Referee : resolves the tennis Referee
func (r *MatchResolver) Referee(ctx context.Context) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.GetRefereeByID(r.match.RefLink)
	return &TennisRefereeResolver{ref: ref}, err
}

// Stadium : resolves the Stadium
func (r *MatchResolver) Stadium(ctx context.Context) (*StadiumResolver, error) {
	stadium, err := r.tennis.GetStadiumByID(r.match.StdLink)
	return &StadiumResolver{stadium: stadium}, err
}

// Sets : resolves the sets
func (r *MatchResolver) Sets() ([]*SetResolver, error) {
	// Todo get the sets
	return make([]*SetResolver, 0), nil
}

// Service : resolves the service
func (r *MatchResolver) Service() *bool {
	return &r.match.Service
}
