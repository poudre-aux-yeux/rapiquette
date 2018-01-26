package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Date    graphql.Time
	Players []graphql.ID
	Referee graphql.ID
}
type startMatchArgs struct {
	ID graphql.ID
}

// CreateMatch : mutation to create a match
func (r *RootResolver) CreateMatch(args *createMatchArgs) *MatchResolver {
	ref := tennis.Referee{ID: args.Referee}
	players := make([]*tennis.Player, len(args.Players))

	for i, p := range args.Players {
		player := tennis.Player{ID: p}
		players[i] = &player
	}

	match := tennis.Match{
		Date:    args.Date,
		Ref:     &ref,
		Players: players,
	}
	// insert in db
	return &MatchResolver{match: match}
}

// StartMatch : starts a match
func (r *RootResolver) StartMatch(args *startMatchArgs) *MatchResolver {
	// retrieve the match
	match := tennis.Match{}
	// start the match and update db
	return &MatchResolver{match: match}
}
