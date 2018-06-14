package resolvers

import (
	"context"
	"math/rand"
	"time"

	graphql "github.com/poudre-aux-yeux/graphql-go"
)

// PointScored : subscribe to scored points for every match
func (r *RootResolver) PointScored(ctx context.Context, args struct{ MatchID graphql.ID }) <-chan *PointScoredEvent {
	c := make(chan *PointScoredEvent)
	r.pointScoredSubscriber <- &PointScoredSubscriber{events: c, stop: ctx.Done(), matchID: args.MatchID}
	return c
}

// MatchCreated : subscribe to match creations
func (r *RootResolver) MatchCreated(ctx context.Context) <-chan *MatchCreatedEvent {
	c := make(chan *MatchCreatedEvent)
	r.matchCreatedSubscriber <- &MatchCreatedSubscriber{events: c, stop: ctx.Done()}
	return c
}

func (r *RootResolver) broadcastPointScored() {
	subscribers := map[string]*PointScoredSubscriber{}
	unsubscribe := make(chan string)
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.pointScoredSubscriber:
			subscribers[randomID()] = s
		case e := <-r.pointScoredEvents:
			for id, s := range subscribers {
				if s.matchID != e.match.ID {
					continue
				}
				go func(id string, s *PointScoredSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(1 * time.Second):
					}
				}(id, s)
			}
		}
	}
}

func (r *RootResolver) broadcastMatchCreated() {
	subscribers := map[string]*MatchCreatedSubscriber{}
	unsubscribe := make(chan string)
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.matchCreatedSubscriber:
			subscribers[randomID()] = s
		case e := <-r.matchCreatedEvents:
			for id, s := range subscribers {
				go func(id string, s *MatchCreatedSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(1 * time.Second):
					}
				}(id, s)
			}
		}
	}
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
