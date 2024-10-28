package repository

import (
	"context"
	"database/sql"

	db "github.com/TiagoAmaralFerreira/go-echo/db/sqlc"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) error
}

type UserRepository struct {
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewUserRepository(sqlDB *sql.DB) *UserRepository {
	q := db.New(sqlDB)
	return &UserRepository{
		DBtx:    sqlDB,
		Queries: q,
		SqlConn: sqlDB,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) error {
	return r.Queries.CreateUser(ctx, arg)
}
