package service

import (
	"context"

	db "github.com/TiagoAmaralFerreira/go-echo/db/sqlc"
	"github.com/TiagoAmaralFerreira/go-echo/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user db.User) (db.User, error)
	GetUser(ctx context.Context, id int) (db.User, error)
	ListUsers(ctx context.Context) ([]db.User, error)
	UpdateUser(ctx context.Context, user db.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repository.UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser(ctx context.Context, user db.User) (db.User, error) {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetUser implements UserService.
func (u *userService) GetUser(ctx context.Context, id int) (db.User, error) {
	panic("unimplemented")
}

// ListUsers implements UserService.
func (u *userService) ListUsers(ctx context.Context) ([]db.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(ctx context.Context, user db.User) error {
	panic("unimplemented")
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
