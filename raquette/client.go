package raquette

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

// Client is a client to fetch raquette data
type Client struct {
	pool *redis.Pool
}

// New creates a new raquette client
func New(pool *redis.Pool) (*Client, error) {
	if pool == nil {
		return nil, errors.New("nil Redis pool")
	}

	return &Client{pool: pool}, nil
}
