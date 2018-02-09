// Package resolvers contains the data resolution layer
package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type queryArgs struct {
	ID graphql.ID
}

// Matches : resolves the Matches query
func (r *RootResolver) Matches(ctx context.Context) ([]*MatchResolver, error) {
	matches, err := r.tennis.GetAllMatches()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*MatchResolver, len(matches))

	for i, match := range matches {
		resolvers[i] = &MatchResolver{match: match, tennis: r.tennis}
	}

	return resolvers, nil
}

// Match : resolves the Match query
func (r *RootResolver) Match(ctx context.Context, args *queryArgs) (*MatchResolver, error) {
	match, err := r.tennis.GetMatchByID(args.ID)

	if err != nil {
		return nil, err
	}

	return &MatchResolver{match: match, tennis: r.tennis}, nil
}

// Players : resolves the Players query
func (r *RootResolver) Players(ctx context.Context) ([]*PlayerResolver, error) {
	players, err := r.tennis.GetAllPlayers()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*PlayerResolver, len(players))

	for i, player := range players {
		resolvers[i] = &PlayerResolver{player: player}
	}

	return resolvers, nil
}

// Player : resolves the Player query
func (r *RootResolver) Player(ctx context.Context, args *queryArgs) (*PlayerResolver, error) {
	player, err := r.tennis.GetPlayerByID(args.ID)

	return &PlayerResolver{player: player}, err
}

// Stadiums : resolves the Stadiums query
func (r *RootResolver) Stadiums(ctx context.Context) ([]*StadiumResolver, error) {
	stds, err := r.tennis.GetAllStadiums()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*StadiumResolver, len(stds))

	for i, stadium := range stds {
		resolvers[i] = &StadiumResolver{stadium: stadium}
	}

	return resolvers, nil
}

// Stadium : resolves the Stadium query
func (r *RootResolver) Stadium(ctx context.Context, args *queryArgs) (*StadiumResolver, error) {
	stadium, err := r.tennis.GetStadiumByID(args.ID)
	return &StadiumResolver{stadium: stadium}, err
}

// TennisReferees : resolves the TennisReferees query
func (r *RootResolver) TennisReferees(ctx context.Context) ([]*TennisRefereeResolver, error) {
	refs, err := r.tennis.GetAllReferees()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*TennisRefereeResolver, len(refs))

	for i, ref := range refs {
		resolvers[i] = &TennisRefereeResolver{ref: ref}
	}

	return resolvers, nil
}

// TennisReferee : resolves the TennisReferee query
func (r *RootResolver) TennisReferee(ctx context.Context, args *queryArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.GetRefereeByID(args.ID)
	return &TennisRefereeResolver{ref: ref}, err
}

// Set : resolves the Set query
func (r *RootResolver) Set(ctx context.Context, args *queryArgs) (*SetResolver, error) {
	set := tennis.Set{}
	return &SetResolver{set: set}, ErrNotImplemented
}

// Game : resolves the Game query
func (r *RootResolver) Game(ctx context.Context, args *queryArgs) (*GameResolver, error) {
	game := tennis.Game{}
	return &GameResolver{game: game}, ErrNotImplemented
}

// Admin : resolves the Admin query
func (r *RootResolver) Admin(ctx context.Context, args *queryArgs) (*AdminResolver, error) {
	admin := raquette.Admin{}
	return &AdminResolver{admin: admin}, ErrNotImplemented
}

// RefereeRaquette : resolves the RefereeRaquette query
func (r *RootResolver) RefereeRaquette(ctx context.Context, args *queryArgs) (*RaquetteRefereeResolver, error) {
	ref := raquette.Referee{}
	return &RaquetteRefereeResolver{ref: ref}, ErrNotImplemented
}
