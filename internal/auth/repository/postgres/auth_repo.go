package postgres

import (
	"context"
	"simple-attendance/internal/auth/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	AddRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (int64, error)
	VerifyRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (string, error)
	DeleteRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (int64, error)
}
