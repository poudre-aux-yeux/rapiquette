package resolvers

import (
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// PointScoredEvent : dispatched when a point is scored
type PointScoredEvent struct {
	team  bool
	match *tennis.Match
}

// MatchCreatedEvent : a match was created
type MatchCreatedEvent struct {
	match *tennis.Match
}

// Team : resolves the team that scored the point
func (r *PointScoredEvent) Team() bool {
	return r.team
}

// Match : resolves the match where the point was score
func (r *PointScoredEvent) Match() (*MatchResolver, error) {
	return &MatchResolver{match: r.match}, nil
}

// Match : resolves the match where the point was score
func (r *MatchCreatedEvent) Match() (*MatchResolver, error) {
	return &MatchResolver{match: r.match}, nil
}
