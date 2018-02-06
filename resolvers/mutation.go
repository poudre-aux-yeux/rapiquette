package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Date    graphql.Time
	Players []graphql.ID
	Referee graphql.ID
}
type createPlayerArgs struct {
	Player tennis.Player
}
type startMatchArgs struct {
	ID graphql.ID
}

// CreateMatch : mutation to create a match
func (r *RootResolver) CreateMatch(ctx context.Context, args *createMatchArgs) (*MatchResolver, error) {
	ref := tennis.Referee{ID: args.Referee}
	players := make([]*tennis.Player, len(args.Players))

	for i, p := range args.Players {
		player := tennis.Player{ID: p}
		players[i] = &player
	}

	m := tennis.Match{
		Date:    args.Date,
		Ref:     &ref,
		Players: players,
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
