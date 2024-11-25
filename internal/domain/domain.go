package domain

import "time"

type Active struct {
	ID        string `json:"id"`
	Bedrooms  string `json:"bedrooms"`
	Bathrooms string `json:"bathrooms"`
}

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Favourites   []string
}

type UserService interface {
}

type UserStore interface {
	Create(NewUser) (*User, RememberToken, error)
	Authenticate(email, pw string) (*User, RememberToken, error)
	ByToken(RememberToken) (*User, error)
	ResetToken(email string) (ResetToken, error)
	UpdatePw(pw string, tok ResetToken) (RememberToken, error)
}
