package auth

import (
	"github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/raquette"
)

var secret = "secretodelavega"

// Claims : JWT Claims
type Claims struct {
	UserID    graphql.ID `json:"ID"`
	IsAdmin   bool       `json:"isAdmin"`
	IsReferee bool       `json:"isReferee"`
	jwt.StandardClaims
}

// LoginPayload : information needed to login
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConfirmLogin : validate the login / password combo
func ConfirmLogin(p *LoginPayload, c *raquette.Client) (*raquette.SimpleUser, error) {
	a, err := c.GetAdminByUsername(p.Username)
	if err == nil {
		user := raquette.SimpleUser{
			ID:       a.ID,
			Username: a.Username,
			Email:    a.Email,
		}
		return &user, nil
	}

	r, err := c.GetRefereeByUsername(p.Username)
	if err != nil {
		return nil, err
	}

	user := raquette.SimpleUser{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
	}

	return &user, nil
}
