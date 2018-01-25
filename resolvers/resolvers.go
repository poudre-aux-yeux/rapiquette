// Package resolvers contains the data resolution layer
package resolvers

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
)

// Resolver : default resolver
type Resolver struct{}

type adminArgs struct {
	ID graphql.ID
}

// Admin : returns an *AdminResolver
func (r *Resolver) Admin(args adminArgs) *AdminResolver {
	admin := raquette.Admin{}
	admin.ID = "abcd"
	admin.Email = "a@b.fr"
	admin.PasswordHash = "2130192UIUFDISFU"
	return &AdminResolver{admin: admin}
}

// AdminResolver : resolves raquette.Admin
type AdminResolver struct {
	admin raquette.Admin
}

// ID : resolves the ID
func (r *AdminResolver) ID() graphql.ID {
	return r.admin.ID
}

// Username : resolves the Username
func (r *AdminResolver) Username() string {
	return r.admin.Username
}

// Email : resolves the Email
func (r *AdminResolver) Email() string {
	return r.admin.Email
}

// PasswordHash : resolves the PasswordHash
func (r *AdminResolver) PasswordHash() string {
	return r.admin.PasswordHash
}
