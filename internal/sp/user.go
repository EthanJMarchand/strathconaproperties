package sp

import "time"

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Favourites   string
}

type RememberToken string

type UserService interface {
	Create(User) (*User, RememberToken, error)
	Authenticate(email, pw string) (*User, RememberToken, error)
	ByToken(RememberToken) (*User, error)
	UpdatePw(pw string) (RememberToken, error)
}
