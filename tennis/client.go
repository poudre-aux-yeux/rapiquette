package tennis

import (
	"encoding/json"
	"errors"
	"fmt"

	graphql "github.com/neelance/graphql-go"
	"github.com/poudre-aux-yeux/rapiquette/kvs"
	"github.com/segmentio/ksuid"
)

// Client is a client to fetch tennis data
type Client struct {
	redis *kvs.Redis
}

func generateID() string {
	return ksuid.New().String()
}

// New creates a new tennis client
func New(redis *kvs.Redis) (*Client, error) {
	if redis == nil {
		return nil, errors.New("redis can't be nil")
	}

	return &Client{redis: redis}, nil
}

// GetAllMatches : Return every match
func (c Client) GetAllMatches() ([]Match, error) {
	keys, err := c.redis.GetSetMembers(matchType())

	if err != nil {
		return nil, err
	}

	matches := make([]Match, 0)

	for _, key := range keys {
		m, err := c.GetMatchByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		matches = append(matches, m)
	}

	return matches, nil
}

// GetAllPlayers : Return every player
func (c Client) GetAllPlayers() ([]Player, error) {
	keys, err := c.redis.GetSetMembers(playerType())

	if err != nil {
		return nil, err
	}

	players := make([]Player, 0)

	for _, key := range keys {
		m, err := c.GetPlayerByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		players = append(players, m)
	}

	return players, nil
}

// GetMatchByID : Finds a match in the key-value store
func (c Client) GetMatchByID(id graphql.ID) (Match, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return Match{}, fmt.Errorf("unable to resolve: %v", err)
	}

	var m Match

	if err = json.Unmarshal(data, &m); err != nil {
		return Match{}, fmt.Errorf("the data is malformed: %v", err)
	}

	return m, nil
}

// GetPlayerByID : Finds a match in the key-value store
func (c Client) GetPlayerByID(id graphql.ID) (Player, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return Player{}, fmt.Errorf("unable to resolve: %v", err)
	}

	var p Player

	if err = json.Unmarshal(data, &p); err != nil {
		return Player{}, fmt.Errorf("the data is malformed: %v", err)
	}

	return p, nil
}

// CreateMatch : Creates a match in the key-value store
func (c Client) CreateMatch(m Match) (Match, error) {
	id := generateID()
	m.ID = graphql.ID(id)

	return m, c.Create(m, id, matchType())
}

// CreatePlayer adds a player to the key-value store
func (c Client) CreatePlayer(p Player) (Player, error) {
	id := generateID()
	p.ID = graphql.ID(id)

	return p, c.Create(p, id, playerType())
}

// Create adds an item to the key-value store
// util
func (c Client) Create(item interface{}, id, set string) error {
	data, err := json.Marshal(item)

	if err != nil {
		return fmt.Errorf("impossible to marshal: %v", err)
	}

	if err = c.redis.Set(id, data); err != nil {
		return fmt.Errorf("could not set the data: %v", err)
	}

	if err = c.redis.AddToSet(set, id); err != nil {
		return fmt.Errorf("couldn't add the key %v to set %v: %v", id, set, err)
	}

	return nil
}
