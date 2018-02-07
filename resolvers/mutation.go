package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Date         graphql.Time
	PlayersLinks []graphql.ID
	RefereeLink  graphql.ID
	StadiumLink  graphql.ID
}
type createPlayerArgs struct {
	Player tennis.Player
}
type startMatchArgs struct {
	ID graphql.ID
}

// CreateMatch : mutation to create a match
func (r *RootResolver) CreateMatch(ctx context.Context, args *createMatchArgs) (*MatchResolver, error) {
	m := tennis.Match{
		Date:         args.Date,
		RefLink:      args.RefereeLink,
		PlayersLinks: args.PlayersLinks,
		StdLink:      args.StadiumLink,
	}

	match, err := r.tennis.CreateMatch(m)

	return &MatchResolver{match: match}, err
}

// CreatePlayer creates a new Player and returns it
func (r *RootResolver) CreatePlayer(ctx context.Context, args *createPlayerArgs) (*PlayerResolver, error) {
	player, err := r.tennis.CreatePlayer(args.Player)

	return &PlayerResolver{player: player}, err
}

// StartMatch : starts a match
func (r *RootResolver) StartMatch(args *startMatchArgs) *MatchResolver {
	// retrieve the match
	match := tennis.Match{}
	// start the match and update db
	return &MatchResolver{match: match}
}
