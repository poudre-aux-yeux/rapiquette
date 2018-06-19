package resolvers

import (
	"github.com/poudre-aux-yeux/rapiquette/tennis"
)

// DisplayableScoreResolver : resolves the displayable scores
type DisplayableScoreResolver struct {
	score *tennis.DisplayableScore
}

// Home : resolves the Home games
func (r *DisplayableScoreResolver) Home() *[]int32 {
	return &r.score.Home
}

// Away : resolves the Away games
func (r *DisplayableScoreResolver) Away() *[]int32 {
	return &r.score.Away
}

// Winner : resolves the winner if exists
func (r *DisplayableScoreResolver) Winner() bool {
	return r.score.Winner
}

// HomePoints : points for the #1 team as a string
func (r *DisplayableScoreResolver) HomePoints() string {
	return r.score.HomePoints
}

// AwayPoints : points for the #2 team as a string
func (r *DisplayableScoreResolver) AwayPoints() string {
	return r.score.AwayPoints
}
