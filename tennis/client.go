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
	var match Match
	keys, err := c.redis.GetSetMembers(match.GetType())

	if err != nil {
		return nil, err
	}

	matches := make([]Match, 0)

	for _, key := range keys {
		match, err = c.GetMatchByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		matches = append(matches, match)
	}

	return matches, nil
}

// GetAllPlayers : Return every player
func (c Client) GetAllPlayers() ([]Player, error) {
	var player Player
	keys, err := c.redis.GetSetMembers(player.GetType())

	if err != nil {
		return nil, err
	}

	players := make([]Player, 0)

	for _, key := range keys {
		player, err = c.GetPlayerByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		players = append(players, player)
	}

	return players, nil
}

// GetAllReferees : Return every player
func (c Client) GetAllReferees() ([]*Referee, error) {
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

// GetAllStadiums : Return every stadium
func (c Client) GetAllStadiums() ([]Stadium, error) {
	var stadium Stadium
	keys, err := c.redis.GetSetMembers(stadium.GetType())

	if err != nil {
		return nil, err
	}

	stadiums := make([]Stadium, 0)

	for _, key := range keys {
		stadium, err = c.GetStadiumByID(graphql.ID(key))

		if err != nil {
			fmt.Printf("error getting item %v: %v\n", key, err)
		}

		stadiums = append(stadiums, stadium)
	}

	return stadiums, nil
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

// GetRefereeByID : Finds a match in the key-value store
func (c Client) GetRefereeByID(id graphql.ID) (*Referee, error) {
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

// GetStadiumByID : Finds a match in the key-value store
func (c Client) GetStadiumByID(id graphql.ID) (Stadium, error) {
	data, err := c.redis.Get(string(id))

	if err != nil {
		return Stadium{}, fmt.Errorf("unable to resolve: %v", err)
	}

	var s Stadium

	if err = json.Unmarshal(data, &s); err != nil {
		return Stadium{}, fmt.Errorf("the data is malformed: %v", err)
	}

	return s, nil
}

// CreateMatch : Creates a match in the key-value store
func (c Client) CreateMatch(m Match) (Match, error) {
	id := generateID()
	m.ID = graphql.ID(id)

	return m, c.Create(m, id, m.GetType())
}

// CreatePlayer adds a player to the key-value store
func (c Client) CreatePlayer(p Player) (Player, error) {
	id := generateID()
	p.ID = graphql.ID(id)

	return p, c.Create(p, id, p.GetType())
}

// CreateReferee adds a tennis referee to the key-value store
func (c Client) CreateReferee(r Referee) (Referee, error) {
	id := generateID()
	r.ID = graphql.ID(id)

	return r, c.Create(r, id, r.GetType())
}

// CreateStadium adds a stadium to the key-value store
func (c Client) CreateStadium(r Stadium) (Stadium, error) {
	id := generateID()
	r.ID = graphql.ID(id)

	return r, c.Create(r, id, r.GetType())
}
