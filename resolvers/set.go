package resolvers

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// SetResolver : resolves tennis.Set
type SetResolver struct {
	set tennis.Set
}

// ID : resolves the ID
func (r *SetResolver) ID() graphql.ID {
	return r.set.ID
}

// Games : resolves the Games
func (r *SetResolver) Games() []*GameResolver {
	return make([]*GameResolver, 0)
}
