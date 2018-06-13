// Package resolvers contains the data resolution layer
package resolvers

import (
	"context"
	"strings"

	graphql "github.com/poudre-aux-yeux/graphql-go"
)

type queryArgs struct {
	ID graphql.ID
}

type searchArgs struct {
	Text string
}

// Admins : resolves the Admin query
func (r *RootResolver) Admins(ctx context.Context) ([]*AdminResolver, error) {
	admins, err := r.raquette.GetAllAdmins()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*AdminResolver, len(admins))

	for i, admin := range admins {
		resolvers[i] = &AdminResolver{admin: admin}
	}

	return resolvers, nil
}

// RaquetteReferees : resolves the RaquetteReferee query
func (r *RootResolver) RaquetteReferees(ctx context.Context) ([]*RaquetteRefereeResolver, error) {
	refs, err := r.raquette.GetAllReferees()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*RaquetteRefereeResolver, len(refs))

	for i, ref := range refs {
		resolvers[i] = &RaquetteRefereeResolver{ref: ref, tennis: r.tennis}
	}

	return resolvers, nil
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
	panic("not implemented")
	// set := tennis.Set{}
	// return &SetResolver{set: set}, ErrNotImplemented
}

// Game : resolves the Game query
func (r *RootResolver) Game(ctx context.Context, args *queryArgs) (*GameResolver, error) {
	panic("not implemented")
	// game := tennis.Game{}
	// return &GameResolver{game: game}, ErrNotImplemented
}

// Admin : resolves the Admin query
func (r *RootResolver) Admin(ctx context.Context, args *queryArgs) (*AdminResolver, error) {
	admin, err := r.raquette.GetAdminByID(args.ID)
	return &AdminResolver{admin: admin}, err
}

// RaquetteReferee : resolves the RaquetteReferee query
func (r *RootResolver) RaquetteReferee(ctx context.Context, args *queryArgs) (*RaquetteRefereeResolver, error) {
	ref, err := r.raquette.GetRefereeByID(args.ID)
	return &RaquetteRefereeResolver{ref: ref}, err
}

// SearchUsers : search Admins and RaquetteReferees
func (r *RootResolver) SearchUsers(ctx context.Context, args *searchArgs) ([]*UserSearchResolver, error) {
	// TODO : Filter with Redis capabilities instead of retrieving everything and filtering
	var l []*UserSearchResolver
	admins, err := r.raquette.GetAllAdmins()

	if err != nil {
		return nil, err
	}

	for _, admin := range admins {
		if strings.Contains(admin.Username, args.Text) {
			l = append(l, &UserSearchResolver{&AdminResolver{admin: admin}})
		}
	}

	refs, err := r.raquette.GetAllReferees()

	if err != nil {
		return nil, err
	}

	for _, ref := range refs {
		if strings.Contains(ref.Username, args.Text) {
			l = append(l, &UserSearchResolver{&RaquetteRefereeResolver{ref: ref}})
		}
	}
	return l, nil
}

// TennisSearch : search Stadiums, tennis Refs, Players and Matches
func (r *RootResolver) TennisSearch(ctx context.Context, args *searchArgs) ([]*SearchResolver, error) {
	// TODO : search stadiums, refs and players with Redis
	// TODO : add all matches with stadiums, refs and players found
	panic("not implemented")
	// return nil, ErrNotImplemented
}
