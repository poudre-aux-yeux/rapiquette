package tennis

import (
	"encoding/json"
	"errors"
	"fmt"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/kvs"
)

// Client is a client to fetch tennis data
type Client struct {
	redis *kvs.Redis
}

// New creates a new tennis client
func New(redis *kvs.Redis) (*Client, error) {
	if redis == nil {
		return nil, errors.New("redis can't be nil")
	}

	return &Client{redis: redis}, nil
}

// GetMatch : Finds a match in the key-value store
func (c Client) GetMatch(id graphql.ID) (Match, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return Match{}, fmt.Errorf("unable to resolve: ", err)
	}

	var m Match

	if err = json.Unmarshal(data, &m); err != nil {
		return Match{}, fmt.Errorf("the data is malformed: ", err)
	}

	return m, nil
}

// CreateMatch : Creates a match in the key-value store
func (c Client) CreateMatch(m Match) (Match, error) {
	id := graphql.ID("TODO: newGraphqlID")
	m.ID = id
	data, err := json.Marshal(m)

	if err != nil {
		return Match{}, fmt.Errorf("impossible to marshal the match: ", err)
	}

	if err = c.redis.Set(string(id), data); err != nil {
		return Match{}, fmt.Errorf("could not set the data: ", err)
	}

	return m, nil
}
