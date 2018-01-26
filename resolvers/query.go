// Package resolvers contains the data resolution layer
package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// Matches : resolves the Matches query
func (r *RootResolver) Matches() []*MatchResolver {
	return make([]*MatchResolver, 0)
}

// Match : resolves the Match query
func (r *RootResolver) Match() *MatchResolver {
	match := tennis.Match{}
	return &MatchResolver{match: match}
}

// Stadium : resolves the Stadium query
func (r *RootResolver) Stadium() *StadiumResolver {
	stadium := tennis.Stadium{}
	return &StadiumResolver{stadium: stadium}
}

type adminArgs struct {
	ID graphql.ID
}

// Admin : resolves the Admin query
func (r *RootResolver) Admin(args adminArgs) *AdminResolver {
	admin := raquette.Admin{}
	admin.ID = "abcd"
	admin.Email = "a@b.fr"
	admin.PasswordHash = "2130192UIUFDISFU"
	return &AdminResolver{admin: admin}
}
