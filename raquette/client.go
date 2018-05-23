package raquette

import (
	"encoding/json"
	"errors"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/kvs"
	"github.com/segmentio/ksuid"
)

// ErrNotFound : a key was not found in the kvs
var ErrNotFound = errors.New("not found")

// Client is a client to fetch raquette data
type Client struct {
	redis *kvs.Redis
}

func generateID() string {
	return ksuid.New().String()
}

// New creates a new raquette client
func New(redis *kvs.Redis) (*Client, error) {
	if redis == nil {
		return nil, errors.New("redis can't be nil")
	}

	return &Client{redis: redis}, nil
}

// GetAllReferees : Return every referee
func (c *Client) GetAllReferees() ([]*Referee, error) {
	var referee Referee
	keys, err := c.redis.GetSetMembers(referee.GetType())

	if err != nil {
		return nil, err
	}

	referees := make([]*Referee, 0)

	for _, key := range keys {
		r, err := c.GetRefereeByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		referees = append(referees, r)
	}

	return referees, nil
}

// GetAllAdmins : Return every admin
func (c *Client) GetAllAdmins() ([]*Admin, error) {
	var admin Admin
	keys, err := c.redis.GetSetMembers(admin.GetType())

	if err != nil {
		return nil, err
	}

	admins := make([]*Admin, 0)

	for _, key := range keys {
		a, err := c.GetAdminByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		admins = append(admins, a)
	}

	return admins, nil
}

// GetAdminByUsername finds and returns an administrator by their username
func (c *Client) GetAdminByUsername(username string) (*Admin, error) {
	var admin Admin
	keys, err := c.redis.GetSetMembers(admin.GetType())

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		a, err := c.GetAdminByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		if a.Username == username {
			return a, nil
		}
	}

	return nil, ErrNotFound
}

// GetRefereeByUsername finds and returns a referee by their username
func (c *Client) GetRefereeByUsername(username string) (*Referee, error) {
	var ref Referee
	keys, err := c.redis.GetSetMembers(ref.GetType())

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		r, err := c.GetRefereeByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		if r.Username == username {
			return r, nil
		}
	}

	return nil, ErrNotFound
}

// GetAdminByID : Finds a Admin in the key-value store
func (c *Client) GetAdminByID(id graphql.ID) (*Admin, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return nil, fmt.Errorf("unable to resolve: %v", err)
	}

	var a Admin

	if err = json.Unmarshal(data, &a); err != nil {
		return nil, fmt.Errorf("the data is malformed: %v", err)
	}

	return &a, nil
}

// GetRefereeByID : Finds a Referee in the key-value store
func (c *Client) GetRefereeByID(id graphql.ID) (*Referee, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return nil, fmt.Errorf("unable to resolve: %v", err)
	}

	var r Referee

	if err = json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("the data is malformed: %v", err)
	}

	return &r, nil
}
