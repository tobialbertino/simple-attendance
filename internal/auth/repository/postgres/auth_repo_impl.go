package postgres

import (
	"context"
	"simple-attendance/internal/auth/models/entity"
	"simple-attendance/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

// AddRefreshToken implements AuthRepository
func (repo *AuthRepositoryImpl) AddRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (int64, error) {
	SQL := `INSERT INTO authentications VALUES ($1)`
	varArgs := []interface{}{
		token.Token,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	return i, nil
}

func (repo *AuthRepositoryImpl) VerifyRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (string, error) {
	var resultToken string
	SQL := `SELECT token FROM authentications WHERE token = $1`
	varArgs := []interface{}{
		token.Token,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result := tx.QueryRow(ctx, SQL, varArgs...)
	if err != nil {
		return "", err
	}

	result.Scan(&resultToken)
	return resultToken, nil
}

func (repo *AuthRepositoryImpl) DeleteRefreshToken(ctx context.Context, db *pgxpool.Pool, token entity.Token) (int64, error) {
	SQL := `DELETE FROM authentications WHERE token = $1`
	varArgs := []interface{}{
		token.Token,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	return i, nil
}
