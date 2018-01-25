package raquette

type Admin struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"hash"`
	Username     string `json:"username"`
}
