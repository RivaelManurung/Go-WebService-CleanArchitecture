package user

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"` // Sembunyikan password dari JSON secara default
	Email    string `json:"email" gorm:"unique"`
}

type UserRepository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}

type UserService interface {
	Register(user *User) (*User, string, error)
	Login(username, password string) (string, error)
}