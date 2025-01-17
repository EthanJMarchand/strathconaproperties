package sp

import (
	"context"
)

type User struct {
	ID           uint
	FirstName    string
	LastName     string
	PasswordHash string
	Phone        string
	Email        string
}

type UserService struct {
	DB UserStore
}

type UserStore interface {
	Create(ctx context.Context, firstName, lastname, email, phone, password string) (*User, error)
	Authenticate(ctx context.Context, email, password string) (*User, error)
	UpdateEmail(ctx context.Context, oldEmail, newEmail string) error
	UpdatePhone(ctx context.Context, oldPhone, newPhone string) error
	GetUsers(ctx context.Context) ([]*User, error)
}

func NewUserService(u UserStore) UserService {
	return UserService{
		DB: u,
	}
}
