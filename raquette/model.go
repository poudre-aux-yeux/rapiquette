// Package raquette contains the models for the Raquette application
package raquette

import (
	graphql "github.com/neelance/graphql-go"
)

// User : an application user
type User struct {
	ID           graphql.ID `json:"id"`
	PasswordHash string     `json:"hash"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
}

// Admin : manages the application
type Admin struct {
	User
}

// Referee : can interact with the Referee front-end
type Referee struct {
	User
}
