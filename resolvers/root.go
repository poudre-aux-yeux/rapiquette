package resolvers

import (
	"errors"

	"github.com/poudre-aux-yeux/rapiquette/raquette"
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

var (
	// ErrUnableToResolve indicates that the resource couldn't be resolved
	ErrUnableToResolve = errors.New("unable to resolve")
	// ErrKeyDoesNotExist indicates that the key was not found in the key value store
	ErrKeyDoesNotExist = errors.New("the key does not exist")
	// ErrNotImplemented indicates that the resolver isn't ready yet
	ErrNotImplemented = errors.New("not implemented")
)

// NewRoot : Create the root resolver with a Redis Pool
func NewRoot(tennis *tennis.Client, raquette *raquette.Client) (*RootResolver, error) {
	if tennis == nil || raquette == nil {
		return nil, ErrUnableToResolve
	}
	r := &RootResolver{
		tennis:                 tennis,
		raquette:               raquette,
		pointScoredEvents:      make(chan *PointScoredEvent),
		pointScoredSubscriber:  make(chan *PointScoredSubscriber),
		matchCreatedEvents:     make(chan *MatchCreatedEvent),
		matchCreatedSubscriber: make(chan *MatchCreatedSubscriber),
	}

	go r.broadcastPointScored()
	go r.broadcastMatchCreated()

	return r, nil
}

// RootResolver : default resolver
type RootResolver struct {
	tennis                 *tennis.Client
	raquette               *raquette.Client
	pointScoredEvents      chan *PointScoredEvent
	pointScoredSubscriber  chan *PointScoredSubscriber
	matchCreatedEvents     chan *MatchCreatedEvent
	matchCreatedSubscriber chan *MatchCreatedSubscriber
}
