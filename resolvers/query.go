// Package resolvers contains the data resolution layer
package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type queryArgs struct {
	ID graphql.ID
}

// Matches : resolves the Matches query
func (r *RootResolver) Matches() []*MatchResolver {
	return make([]*MatchResolver, 0)
}

// Match : resolves the Match query
func (r *RootResolver) Match(args *queryArgs) *MatchResolver {
	match, err := r.tennis.GetMatch(args.ID)

	if err != nil {
		// TODO : return an error
		return &MatchResolver{}
	}

	return &MatchResolver{match: match}
}

// Player : resolves the Player query
func (r *RootResolver) Player(args *queryArgs) *PlayerResolver {
	player := tennis.Player{}
	return &PlayerResolver{player: player}
}

// Stadium : resolves the Stadium query
func (r *RootResolver) Stadium(args *queryArgs) *StadiumResolver {
	stadium := tennis.Stadium{}
	return &StadiumResolver{stadium: stadium}
}

// RefereeTennis : resolves the RefereeTennis query
func (r *RootResolver) RefereeTennis(args *queryArgs) *TennisRefereeResolver {
	ref := tennis.Referee{}
	return &TennisRefereeResolver{ref: ref}
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
