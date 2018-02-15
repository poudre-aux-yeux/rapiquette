package resolvers

// UserSearchResolver : result of an user search
type UserSearchResolver struct {
	Result interface{} `json:"result"`
}

// ToAdmin converts a result to an Admin
func (r *UserSearchResolver) ToAdmin() (*AdminResolver, bool) {
	res, ok := r.Result.(*AdminResolver)
	return res, ok
}

// ToRaquetteReferee converts a result to a RaquetteReferee
func (r *UserSearchResolver) ToRaquetteReferee() (*RaquetteRefereeResolver, bool) {
	res, ok := r.Result.(*RaquetteRefereeResolver)
	return res, ok
}
