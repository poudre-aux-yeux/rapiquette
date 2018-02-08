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

// Stadium : resolves the Stadium query
func (r *RootResolver) Stadium(args *queryArgs) *StadiumResolver {
	stadium := tennis.Stadium{}
	return &StadiumResolver{stadium: stadium}
}

// RefereeTennis : resolves the RefereeTennis query
func (r *RootResolver) RefereeTennis(ctx context.Context, args *queryArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.GetRefereeByID(args.ID)
	return &TennisRefereeResolver{ref: ref}, err
}

// Set : resolves the Set query
func (r *RootResolver) Set(args *queryArgs) *SetResolver {
	set := tennis.Set{}
	return &SetResolver{set: set}
}

// Game : resolves the Game query
func (r *RootResolver) Game(args *queryArgs) *GameResolver {
	game := tennis.Game{}
	return &GameResolver{game: game}
}

// Admin : resolves the Admin query
func (r *RootResolver) Admin(args *queryArgs) *AdminResolver {
	admin := raquette.Admin{}
	admin.ID = "abcd"
	admin.Email = "a@b.fr"
	admin.PasswordHash = "2130192UIUFDISFU"
	return &AdminResolver{admin: admin}
}

// RefereeRaquette : resolves the RefereeRaquette query
func (r *RootResolver) RefereeRaquette(args *queryArgs) *RaquetteRefereeResolver {
	ref := raquette.Referee{}
	return &RaquetteRefereeResolver{ref: ref}
}
