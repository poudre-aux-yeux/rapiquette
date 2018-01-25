package raquette

// User : an application user
type User struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"hash"`
	Username     string `json:"username"`
	Email        string `json:"email"`
}

// Admin : manages the application
type Admin struct {
	User
}

// Referee : can interact with the Referee front-end
type Referee struct {
	User
}
