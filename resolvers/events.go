package resolvers

import (
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// PointScoredEvent : dispatched when a point is scored
type PointScoredEvent struct {
	team   bool
	match  *tennis.Match
	tennis *tennis.Client
}

// Team : resolves the team that scored the point
func (r *PointScoredEvent) Team() bool {
	return r.team
}

// Match : resolves the match where the point was score
func (r *PointScoredEvent) Match() (*MatchResolver, error) {
	match, err := r.tennis.GetMatchByID(r.match.ID)
	return &MatchResolver{match: match}, err
}