package store

import (
	"context"
	"fmt"
	"github.com/ethanjmarchand/StrathconaProperties/internal/sp"
	"github.com/ethanjmarchand/StrathconaProperties/internal/startup"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Postgres struct {
	DB *pgxpool.Pool
}

// NewPostgresStore callers need to remember to call Close() on the returned *Postgres.
func NewPostgresStore(s *startup.Config) (*Postgres, error) {
	ctx := context.Background()
	cp, err := pgxpool.New(ctx, "postgresql://"+s.DBUsername+":"+s.DBPassword+"@localhost:5432/"+s.DBName+"?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("problem with the database connection string: %w", err)
	}
	err = cp.Ping(ctx)
	cp.Reset()
	if err != nil {
		return nil, fmt.Errorf("failed database ping: %w", err)
	}
	fmt.Println("database connected successfully")
	return &Postgres{
		DB: cp,
	}, nil
}

func (p *Postgres) Create(ctx context.Context, firstName, lastname, email, phone, password string) (*sp.User, error) {
	email = strings.ToLower(email)
	pwdhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user := sp.User{
		FirstName: firstName,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
	}
	query := `INSERT INTO users (first_name, last_name, email, phone, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	rows := p.DB.QueryRow(ctx, query, firstName, lastname, email, phone, string(pwdhash))
	err = rows.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return &user, nil
}

// Authenticate is
func (p *Postgres) Authenticate(ctx context.Context, email, password string) (*sp.User, error) {
	user := sp.User{
		Email: strings.ToLower(email),
	}
	row := p.DB.QueryRow(ctx, "SELECT password_hash FROM users WHERE email = $1", user.Email)
	err := row.Scan(&user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed authenticating user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("failed authenticating user: %w", err)
	}
	return &user, nil
}

func (p *Postgres) UpdateEmail(ctx context.Context, oldEmail, newEmail string) error {
	return nil
}
func (p *Postgres) UpdatePhone(ctx context.Context, oldPhone, newPhone string) error {
	return nil
}
func (p *Postgres) GetUsers(ctx context.Context) ([]*sp.User, error) {
	rows, err := p.DB.Query(ctx, `SELECT id, first_name, last_name, email, phone FROM users;`)
	if err != nil {
		return nil, fmt.Errorf("failed getting users before for loop: %w", err)
	}
	defer rows.Close()
	var users []*sp.User
	for rows.Next() {
		var user sp.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed getting users inside for: %w", err)
		}
		users = append(users, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed getting users after for loop: %w", err)
	}
	return users, nil
}
