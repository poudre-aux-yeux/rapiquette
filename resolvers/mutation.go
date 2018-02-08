package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Date             graphql.Time
	HomePlayersLinks []graphql.ID
	AwayPlayersLinks []graphql.ID
	RefereeLink      graphql.ID
	StadiumLink      graphql.ID
}
type createPlayerArgs struct {
	Player tennis.Player
}
type createTennisRefereeArgs struct {
	Referee tennis.Referee
}
type createStadiumArgs struct {
	Stadium tennis.Stadium
}
type startMatchArgs struct {
	ID graphql.ID
}

// CreateMatch : mutation to create a match
func (r *RootResolver) CreateMatch(ctx context.Context, args *createMatchArgs) (*MatchResolver, error) {
	m := tennis.Match{
		Date:             args.Date,
		RefLink:          args.RefereeLink,
		HomePlayersLinks: args.HomePlayersLinks,
		AwayPlayersLinks: args.AwayPlayersLinks,
		StdLink:          args.StadiumLink,
	}

	match, err := r.tennis.CreateMatch(m)

	return &MatchResolver{match: match}, err
}

// CreatePlayer creates a new Player and returns it
func (r *RootResolver) CreatePlayer(ctx context.Context, args *createPlayerArgs) (*PlayerResolver, error) {
	player, err := r.tennis.CreatePlayer(args.Player)

	return &PlayerResolver{player: player}, err
}

// CreateTennisReferee creates a new Referee and returns it
func (r *RootResolver) CreateTennisReferee(ctx context.Context, args *createTennisRefereeArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.CreateReferee(args.Referee)

	return &TennisRefereeResolver{ref: ref}, err
}

// CreateStadium creates a new Stadium and returns it
func (r *RootResolver) CreateStadium(ctx context.Context, args *createStadiumArgs) (*StadiumResolver, error) {
	stadium, err := r.tennis.CreateStadium(args.Stadium)

	return &StadiumResolver{stadium: stadium}, err
}

// StartMatch : starts a match
func (r *RootResolver) StartMatch(args *startMatchArgs) *MatchResolver {
	// retrieve the match
	match := tennis.Match{}
	// start the match and update db
	return &MatchResolver{match: match}
}
