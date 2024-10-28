package repository

import (
	"context"
	"database/sql"

	db "github.com/TiagoAmaralFerreira/go-echo/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user db.User) (db.User, error)
	GetUser(ctx context.Context, id int) (db.User, error)
	ListUsers(ctx context.Context) ([]db.User, error)
	UpdateUser(ctx context.Context, user db.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *sql.DB
}

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, user db.User) (db.User, error) {
	panic("unimplemented")
}

// DeleteUser implements UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetUser implements UserRepository.
func (u *userRepository) GetUser(ctx context.Context, id int) (db.User, error) {
	panic("unimplemented")
}

// ListUsers implements UserRepository.
func (u *userRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, user db.User) error {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
