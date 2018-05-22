// Package raquette contains the models for the Raquette application
package raquette

import (
	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// User : an application user
type User interface {
	ID() graphql.ID
	PasswordHash() string
	Username() string
	Email() string
}

// SimpleUser : an user without any particular right
type SimpleUser struct {
	ID           graphql.ID `json:"id"`
	PasswordHash string     `json:"hash"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
}

// Admin : manages the application
type Admin struct {
	ID           graphql.ID `json:"id"`
	PasswordHash string     `json:"hash"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
}

// Referee : can interact with the Referee front-end
type Referee struct {
	ID           graphql.ID `json:"id"`
	PasswordHash string     `json:"hash"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	Ref          *tennis.Referee
	RefLink      graphql.ID `json:"referee"`
}

// GetType returns the type of the struct
func (s *Admin) GetType() string {
	return "Admin"
}

// GetType returns the type of the struct
func (s *Referee) GetType() string {
	return "Referee"
}

// GetType returns the type of the struct
func (s *SimpleUser) GetType() string {
	return "SimpleUser"
}
