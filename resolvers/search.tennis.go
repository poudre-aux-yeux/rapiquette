package resolvers

// SearchResolver : result of a tennis search
type SearchResolver struct {
	Result interface{} `json:"result"`
}

// ToPlayer converts a result to an Player
func (r *SearchResolver) ToPlayer() (*PlayerResolver, bool) {
	res, ok := r.Result.(*PlayerResolver)
	return res, ok
}

// ToTennisReferee converts a result to a TennisReferee
func (r *SearchResolver) ToTennisReferee() (*TennisRefereeResolver, bool) {
	res, ok := r.Result.(*TennisRefereeResolver)
	return res, ok
}

// ToMatch converts a result to an Match
func (r *SearchResolver) ToMatch() (*MatchResolver, bool) {
	res, ok := r.Result.(*MatchResolver)
	return res, ok
}

// ToStadium converts a result to a Stadium
func (r *SearchResolver) ToStadium() (*StadiumResolver, bool) {
	res, ok := r.Result.(*StadiumResolver)
	return res, ok
}
