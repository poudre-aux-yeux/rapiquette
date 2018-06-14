package resolvers

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	graphql "github.com/poudre-aux-yeux/graphql-go"
)

type scorePointArgs struct {
	MatchID graphql.ID
	Team    bool
}

// ScorePoint : score a point for a team in a match
func (r *RootResolver) ScorePoint(args struct{ Point struct{ scorePointArgs } }) (*PointScoredEvent, error) {
	m, err := r.tennis.GetMatchByID(args.Point.MatchID)

	if err != nil {
		return nil, fmt.Errorf("couldn't get the match %s: %v", args.Point.MatchID, err)
	}

	e := &PointScoredEvent{match: m, team: args.Point.Team, tennis: r.tennis}
	go func() {
		select {
		case r.pointScoredEvents <- e:
		case <-time.After(1 * time.Second):
		}
	}()
	return e, nil
}

// PointScored : A new point was scored in a match
func (r *RootResolver) PointScored(ctx context.Context) <-chan *PointScoredEvent {
	c := make(chan *PointScoredEvent)
	// NOTE: this could take a while
	r.pointScoredSubscriber <- &PointScoredSubscriber{events: c, stop: ctx.Done()}

	return c
}

func (r *RootResolver) broadcastPointScored() {
	subscribers := map[string]*PointScoredSubscriber{}
	unsubscribe := make(chan string)

	// NOTE: subscribing and sending events are at odds.
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.pointScoredSubscriber:
			subscribers[randomID()] = s
		case e := <-r.pointScoredEvents:
			for id, s := range subscribers {
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
					case <-time.After(time.Second):
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
