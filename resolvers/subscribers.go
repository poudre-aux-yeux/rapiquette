package resolvers

import graphql "github.com/poudre-aux-yeux/graphql-go"

// PointScoredSubscriber : subscribers to pointScored
type PointScoredSubscriber struct {
	stop    <-chan struct{}
	events  chan<- *PointScoredEvent
	matchID graphql.ID
}

// MatchCreatedSubscriber : subscribers to matchCreated
type MatchCreatedSubscriber struct {
	stop   <-chan struct{}
	events chan<- *MatchCreatedEvent
}
