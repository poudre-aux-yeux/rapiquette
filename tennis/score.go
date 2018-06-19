package tennis

import (
	"errors"
	"fmt"
)

// ScorePoint : Mutates a score to reflect a point added
func (s *Score) ScorePoint(team bool) error {
	if mwon, _, _, _, _, _, _, _ := matchWon(s); mwon {
		return errors.New("The match is already over")
	}

	// Le service passe à l'adversaire (version simple)
	s.Service = !s.Service

	// Si le match n'a pas démarré on crée le premier jeu / premier set
	if len(s.Sets) == 0 {
		s.Sets = append(s.Sets, newSet(team))
	}

	lastSet := s.Sets[len(s.Sets)-1]
	lastGame := lastSet.Games[len(lastSet.Games)-1]
	if swon, _, _, _, _, _ := setWon(lastSet); swon {
		s.Sets = append(s.Sets, newSet(team))
		fmt.Println("Added a new set")
	} else if gwon, _ := gameWon(lastGame); gwon {
		s.Sets[len(s.Sets)-1].Games = append(lastSet.Games, newGame(team))
		fmt.Println("Added a new game")
	} else if team {
		fmt.Println("Added a point to the home team")
		s.Sets[len(s.Sets)-1].Games[len(s.Sets[len(s.Sets)-1].Games)-1].HomePoints++
	} else {
		fmt.Println("Added a point to the away team")
		s.Sets[len(s.Sets)-1].Games[len(s.Sets[len(s.Sets)-1].Games)-1].AwayPoints++
	}

	return nil
}

// DisplayableScore : a score in an easier to display way
type DisplayableScore struct {
	Home       []int32 `json:"home"`
	Away       []int32 `json:"away"`
	Winner     bool    `json:"winner"`
	HomePoints string  `json:"homePoints"`
	AwayPoints string  `json:"awayPoints"`
}

// CalculateScore : Calculate a displayable score
// exemples:
// 6/2 2/6 6/4 30/40
// 6/0 6/0 6/1
// AV/40
func (s *Score) CalculateScore() *DisplayableScore {
	won, winner, _, _, hGames, aGames, hPts, aPts := matchWon(s)

	score := DisplayableScore{
		Home: hGames,
		Away: aGames,
	}

	if !won {
		score.HomePoints, score.AwayPoints = displayablePoints(hPts, aPts)
	} else {
		score.Winner = winner
	}

	return &score
}

func displayablePoints(home, away int32) (string, string) {
	if home > 2 && away > 2 {
		if home > away {
			return "A", "40"
		} else if away > home {
			return "40", "A"
		} else {
			return "40", "40"
		}
	}

	return naivePoints(home), naivePoints(away)
}

func naivePoints(score int32) string {
	switch score {
	case 0:
		return "0"
	case 1:
		return "15"
	case 2:
		return "30"
	default:
		return "ADV"
	}
}

// is the game won, and by who ?
func gameWon(game *Game) (bool, bool) {
	if game.AwayPoints < 4 && game.HomePoints < 4 {
		return false, false
	}

	diff := game.AwayPoints - game.HomePoints
	if diff < 0 {
		diff = -diff
	}

	if diff < 2 {
		return false, false
	}

	if game.AwayPoints > 3 && game.AwayPoints > game.HomePoints {
		return true, false
	} else if game.HomePoints > 3 && game.HomePoints > game.AwayPoints {
		return true, true
	}

	return false, false
}

// is the set won, and by who ?
func setWon(set *Set) (bool, bool, int32, int32, int32, int32) {
	awayWon, homeWon := int32(0), int32(0)
	for _, game := range set.Games {
		if won, winner := gameWon(game); won {
			if winner {
				homeWon++
			} else {
				awayWon++
			}
		}
	}

	lastGameHomePoints := set.Games[len(set.Games)-1].HomePoints
	lastGameAwayPoints := set.Games[len(set.Games)-1].AwayPoints

	if homeWon < 6 && awayWon < 6 {
		return false, false, homeWon, awayWon, lastGameHomePoints, lastGameAwayPoints
	}

	diff := homeWon - awayWon

	if diff < 2 && diff > -2 {
		return false, false, homeWon, awayWon, lastGameHomePoints, lastGameAwayPoints
	}

	if homeWon > 5 || awayWon > 5 {
		return true, diff > 0, homeWon, awayWon, 0, 0
	}

	return false, false, homeWon, awayWon, lastGameHomePoints, lastGameAwayPoints
}

// is the match won, and by who ?
func matchWon(score *Score) (bool, bool, int32, int32, []int32, []int32, int32, int32) {
	homeWon, awayWon := int32(0), int32(0)
	homeSets, awaySets := make([]int32, 0), make([]int32, 0)
	homePoints, awayPoints := int32(0), int32(0)
	for _, set := range score.Sets {
		if won, winner, home, away, lghp, lgap := setWon(set); won {
			if winner {
				homeWon++
			} else {
				awayWon++
			}
			homeSets = append(homeSets, home)
			awaySets = append(awaySets, away)
			homePoints = lghp
			awayPoints = lgap
		}
	}

	if homeWon == 3 {
		return true, true, homeWon, awayWon, homeSets, awaySets, 0, 0
	} else if awayWon == 3 {
		return true, false, homeWon, awayWon, homeSets, awaySets, 0, 0
	} else {
		return false, false, homeWon, awayWon, homeSets, awaySets, homePoints, awayPoints
	}
}

// create a new set
func newSet(team bool) *Set {
	s := Set{
		Games: make([]*Game, 0),
	}

	s.Games = append(s.Games, newGame(team))
	return &s
}

// create a new game
func newGame(team bool) *Game {
	g := Game{}
	if team {
		g.HomePoints++
	} else {
		g.AwayPoints++
	}
	return &g
}
