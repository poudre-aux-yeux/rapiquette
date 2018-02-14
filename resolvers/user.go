package resolvers

import (
	"github.com/poudre-aux-yeux/rapiquette/raquette"
)

// UserResolver : resolves raquette.User
type UserResolver struct {
	raquette.User
}

// ToAdmin : converts to an admin
func (r *UserResolver) ToAdmin() (*AdminResolver, bool) {
	u, ok := r.User.(*AdminResolver)
	return u, ok
}
