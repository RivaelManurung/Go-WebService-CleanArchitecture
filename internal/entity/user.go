package entity

// User merepresentasikan data pengguna
type User struct {
	ID       int64
	Email    string
	Password string // Ini adalah hashed password
}