package resolvers

import (
	"context"
	"errors"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

type createMatchArgs struct {
	Match tennis.Match
}
type updateMatchArgs struct {
	Match tennis.Match
}

type createPlayerArgs struct {
	Player struct {
		Name        string       `json:"name"`
		Image       string       `json:"image"`
		Birth       graphql.Time `json:"birth"`
		Nationality string       `json:"nationality"`
		Weight      *float64     `json:"weight"`
		Ranking     *float64     `json:"ranking"`
		Titles      float64      `json:"titles"`
		Height      *float64     `json:"height"`
	} `json:"player"`
}
type updatePlayerArgs struct {
	Player struct {
		ID          graphql.ID   `json:"id"`
		Name        string       `json:"name"`
		Image       string       `json:"image"`
		Birth       graphql.Time `json:"birth"`
		Nationality string       `json:"nationality"`
		Weight      *float64     `json:"weight"`
		Ranking     *float64     `json:"ranking"`
		Titles      float64      `json:"titles"`
		Height      *float64     `json:"height"`
	} `json:"player"`
}

type createTennisRefereeArgs struct {
	Referee tennis.Referee
}
type updateTennisRefereeArgs struct {
	Referee tennis.Referee
}

type createStadiumArgs struct {
	Stadium tennis.Stadium
}
type updateStadiumArgs struct {
	Stadium tennis.Stadium
}

// CreateMatch : mutation to create a match
func (r *RootResolver) CreateMatch(ctx context.Context, args *createMatchArgs) (*MatchResolver, error) {
	if args.Match.RefLink == "" {
		return nil, errors.New("referee is required")
	}

	if args.Match.StdLink == "" {
		return nil, errors.New("stadium is required")
	}

	nbHomePlayers := len(args.Match.HomePlayersLinks)
	nbAwayPlayers := len(args.Match.AwayPlayersLinks)

	if nbHomePlayers < 1 || nbAwayPlayers < 1 {
		return nil, errors.New("not enough players")
	}

	if nbHomePlayers > 2 || nbAwayPlayers > 2 {
		return nil, errors.New("too many players")
	}

	if nbHomePlayers != nbAwayPlayers {
		return nil, errors.New("the teams are not balanced")
	}

	allPlayers := append(args.Match.HomePlayersLinks, args.Match.AwayPlayersLinks...)

	if !containsUniqueKeys(allPlayers) {
		return nil, errors.New("duplicate players found")
	}

	refType := (&tennis.Referee{}).GetType()
	if err := r.existsInSet(refType, args.Match.RefLink); err != nil {
		return nil, err
	}

	stdType := (&tennis.Stadium{}).GetType()
	if err := r.existsInSet(stdType, args.Match.StdLink); err != nil {
		return nil, err
	}

	plType := (&tennis.Player{}).GetType()
	for _, p := range allPlayers {
		if err := r.existsInSet(plType, p); err != nil {
			return nil, err
		}
	}

	// m := tennis.Match{
	// 	Date:             args.Match.Date,
	// 	RefLink:          args.Match.RefLink,
	// 	HomePlayersLinks: args.Match.HomePlayersLinks,
	// 	AwayPlayersLinks: args.Match.AwayPlayersLinks,
	// 	StdLink:          args.Match.StdLink,
	// }

	match, err := r.tennis.CreateMatch(args.Match)

	return &MatchResolver{match: match, tennis: r.tennis}, err
}

// UpdateMatch : mutation to update a match
func (r *RootResolver) UpdateMatch(ctx context.Context, args *updateMatchArgs) (*MatchResolver, error) {
	if args.Match.ID == "" {
		return nil, errors.New("id is required")
	}

	if args.Match.RefLink == "" {
		return nil, errors.New("referee is required")
	}

	if args.Match.StdLink == "" {
		return nil, errors.New("stadium is required")
	}

	nbHomePlayers := len(args.Match.HomePlayersLinks)
	nbAwayPlayers := len(args.Match.AwayPlayersLinks)

	if nbHomePlayers < 1 || nbAwayPlayers < 1 {
		return nil, errors.New("not enough players")
	}

	if nbHomePlayers > 2 || nbAwayPlayers > 2 {
		return nil, errors.New("too many players")
	}

	if nbHomePlayers != nbAwayPlayers {
		return nil, errors.New("the teams are not balanced")
	}

	allPlayers := append(args.Match.HomePlayersLinks, args.Match.AwayPlayersLinks...)

	if !containsUniqueKeys(allPlayers) {
		return nil, errors.New("duplicate players found")
	}

	refType := (&tennis.Referee{}).GetType()
	if err := r.existsInSet(refType, args.Match.RefLink); err != nil {
		return nil, err
	}

	stdType := (&tennis.Stadium{}).GetType()
	if err := r.existsInSet(stdType, args.Match.StdLink); err != nil {
		return nil, err
	}

	plType := (&tennis.Player{}).GetType()
	for _, p := range allPlayers {
		if err := r.existsInSet(plType, p); err != nil {
			return nil, err
		}
	}

	// m := tennis.Match{
	// 	ID:               args.ID,
	// 	Date:             args.Date,
	// 	RefLink:          args.Referee,
	// 	HomePlayersLinks: args.HomePlayers,
	// 	AwayPlayersLinks: args.AwayPlayers,
	// 	StdLink:          args.Stadium,
	// }

	match, err := r.tennis.UpdateMatch(args.Match)

	return &MatchResolver{match: match, tennis: r.tennis}, err
}

// CreatePlayer creates a new Player and returns it
func (r *RootResolver) CreatePlayer(ctx context.Context, args *createPlayerArgs) (*PlayerResolver, error) {
	titles := int32(args.Player.Titles)

	var (
		weight, ranking, height int32
	)

	if args.Player.Weight != nil {
		weight = int32(*args.Player.Weight)
	}

	if args.Player.Ranking != nil {
		ranking = int32(*args.Player.Ranking)
	}

	if args.Player.Height != nil {
		height = int32(*args.Player.Height)
	}

	p := tennis.Player{
		Name:        args.Player.Name,
		Image:       args.Player.Image,
		Birth:       args.Player.Birth,
		Nationality: args.Player.Nationality,
		Weight:      &weight,
		Ranking:     &ranking,
		Titles:      &titles,
		Height:      &height,
	}

	player, err := r.tennis.CreatePlayer(p)

	return &PlayerResolver{player: player}, err
}

// UpdatePlayer updates a Player if it exists and returns it
func (r *RootResolver) UpdatePlayer(ctx context.Context, args *updatePlayerArgs) (*PlayerResolver, error) {
	titles := int32(args.Player.Titles)

	var (
		weight, ranking, height int32
	)

	if args.Player.Weight != nil {
		weight = int32(*args.Player.Weight)
	}

	if args.Player.Ranking != nil {
		ranking = int32(*args.Player.Ranking)
	}

	if args.Player.Height != nil {
		height = int32(*args.Player.Height)
	}

	p := tennis.Player{
		ID:          args.Player.ID,
		Name:        args.Player.Name,
		Image:       args.Player.Image,
		Birth:       args.Player.Birth,
		Nationality: args.Player.Nationality,
		Weight:      &weight,
		Ranking:     &ranking,
		Titles:      &titles,
		Height:      &height,
	}
	player, err := r.tennis.UpdatePlayer(p)

	return &PlayerResolver{player: player}, err
}

// CreateTennisReferee creates a new Referee and returns it
func (r *RootResolver) CreateTennisReferee(ctx context.Context, args *createTennisRefereeArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.CreateReferee(args.Referee)

	return &TennisRefereeResolver{ref: ref}, err
}

// UpdateTennisReferee updates a tennis referee if it existss
func (r *RootResolver) UpdateTennisReferee(ctx context.Context, args *updateTennisRefereeArgs) (*TennisRefereeResolver, error) {
	ref, err := r.tennis.UpdateReferee(args.Referee)

	return &TennisRefereeResolver{ref: ref}, err
}

// CreateStadium creates a new Stadium and returns it
func (r *RootResolver) CreateStadium(ctx context.Context, args *createStadiumArgs) (*StadiumResolver, error) {
	stadium, err := r.tennis.CreateStadium(args.Stadium)

	return &StadiumResolver{stadium: stadium}, err
}

// UpdateStadium updates a stadium if it exists
func (r *RootResolver) UpdateStadium(ctx context.Context, args *updateStadiumArgs) (*StadiumResolver, error) {
	stadium, err := r.tennis.UpdateStadium(args.Stadium)

	return &StadiumResolver{stadium: stadium}, err
}

func (r *RootResolver) existsInSet(set string, key graphql.ID) error {
	if ex, err := r.tennis.KeyExistsInSet(set, string(key)); err != nil {
		return err
	} else if !ex {
		return fmt.Errorf("%v: key: %s, set: %s", ErrKeyDoesNotExist, key, set)
	}

	return nil
}

// checks if the array has no duplicate
func containsUniqueKeys(keys []graphql.ID) bool {
	kmap := make(map[graphql.ID]bool)
	for _, key := range keys {
		kmap[key] = true
	}

	return len(kmap) == len(keys)
}
