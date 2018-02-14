package resolvers

import (
	"context"
	"errors"
	"fmt"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Date        graphql.Time
	HomePlayers []graphql.ID
	AwayPlayers []graphql.ID
	Referee     graphql.ID
	Stadium     graphql.ID
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
	if args.Referee == "" {
		return nil, errors.New("referee is required")
	}

	if args.Stadium == "" {
		return nil, errors.New("stadium is required")
	}

	nbHomePlayers := len(args.HomePlayers)
	nbAwayPlayers := len(args.AwayPlayers)

	if nbHomePlayers < 1 || nbAwayPlayers < 1 {
		return nil, errors.New("not enough players")
	}

	if nbHomePlayers > 2 || nbAwayPlayers > 2 {
		return nil, errors.New("too many players")
	}

	if nbHomePlayers != nbAwayPlayers {
		return nil, errors.New("the teams are not balanced")
	}

	allPlayers := append(args.HomePlayers, args.AwayPlayers...)

	if !containsUniqueKeys(allPlayers) {
		return nil, errors.New("duplicate players found")
	}

	refType := (&tennis.Referee{}).GetType()
	if err := r.existsInSet(refType, args.Referee); err != nil {
		return nil, err
	}

	stdType := (&tennis.Stadium{}).GetType()
	if err := r.existsInSet(stdType, args.Stadium); err != nil {
		return nil, err
	}

	plType := (&tennis.Player{}).GetType()
	for _, p := range allPlayers {
		if err := r.existsInSet(plType, p); err != nil {
			return nil, err
		}
	}

	m := tennis.Match{
		Date:             args.Date,
		RefLink:          args.Referee,
		HomePlayersLinks: args.HomePlayers,
		AwayPlayersLinks: args.AwayPlayers,
		StdLink:          args.Stadium,
	}

	match, err := r.tennis.CreateMatch(m)

	return &MatchResolver{match: match, tennis: r.tennis}, err
}

// CreatePlayer creates a new Player and returns it
func (r *RootResolver) CreatePlayer(ctx context.Context, args *createPlayerArgs) (*PlayerResolver, error) {
	player, err := r.tennis.CreatePlayer(args.Player)

	return &PlayerResolver{player: player}, err
}

// CreateTennisReferee creates a new Referee and returns it
func (r *RootResolver) CreateTennisReferee(ctx context.Context, args *createTennisRefereeArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.CreateReferee(args.Referee)

	return &TennisRefereeResolver{ref: &ref}, err
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

func (r *RootResolver) existsInSet(set string, key graphql.ID) error {
	if ex, err := r.tennis.KeyExistsInSet(set, string(key)); err != nil {
		return err
	} else if !ex {
		return fmt.Errorf("%v: key: %s, set: %s", ErrKeyDoesNotExist, key, set)
	}

	return nil
}

func containsUniqueKeys(keys []graphql.ID) bool {
	kmap := make(map[graphql.ID]bool)
	for _, key := range keys {
		kmap[key] = true
	}

	return len(kmap) == len(keys)
}
