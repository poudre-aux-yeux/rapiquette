package resolvers

import (
	graphql "github.com/poudre-aux-yeux/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
)

// AdminResolver : resolves raquette.Admin
type AdminResolver struct {
	admin *raquette.Admin
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
