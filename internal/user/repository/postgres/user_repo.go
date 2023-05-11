package postgres

import (
	"context"
	"simple-attendance/internal/user/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CheckUsername(ctx context.Context, db *pgxpool.Pool, user entity.User) (int, error)
	AddUser(ctx context.Context, db *pgxpool.Pool, user entity.User) (string, error)
	GetUserById(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.User, error)
	GetUsersByUsername(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.ListUser, error)

	VerifyUserCredential(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.User, error)
}
