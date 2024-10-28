package service

import (
	"context"

	db "github.com/TiagoAmaralFerreira/go-echo/db/sqlc"
	"github.com/TiagoAmaralFerreira/go-echo/internal/repository"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user db.User) error
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user db.User) error {
	err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	return err
}
