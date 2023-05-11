package postgres

import (
	"context"
	"simple-attendance/internal/user/models/entity"
	"simple-attendance/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// CheckUsername implements UserRepository
func (repo *UserRepositoryImpl) CheckUsername(ctx context.Context, db *pgxpool.Pool, user entity.User) (int, error) {
	SQL := `SELECT username FROM users WHERE username = $1`
	varArgs := []interface{}{
		user.Username,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return 2, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL, varArgs...)
	if err != nil {
		return 2, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		counter++
	}

	return counter, nil
}

func (repo *UserRepositoryImpl) AddUser(ctx context.Context, db *pgxpool.Pool, user entity.User) (string, error) {
	var id string

	SQL := `INSERT INTO users VALUES ($1, $2, $3, $4) RETURNING id`
	varArgs := []interface{}{
		user.Id,
		user.Username,
		user.Passwword,
		user.FullName,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result := tx.QueryRow(ctx, SQL, varArgs...)

	result.Scan(&id)

	return id, nil
}

func (repo *UserRepositoryImpl) GetUserById(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.User, error) {
	var res entity.User

	SQL := `SELECT id, username, fullname FROM users WHERE id = $1`
	varArgs := []interface{}{
		user.Id,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.User{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&res.Id, &res.Username, &res.FullName)

	return res, nil
}

func (repo *UserRepositoryImpl) VerifyUserCredential(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.User, error) {
	var res entity.User

	SQL := `SELECT id, password FROM users WHERE username = $1`
	varArgs := []interface{}{
		user.Username,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.User{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result := tx.QueryRow(ctx, SQL, varArgs...)
	if err != nil {
		return entity.User{}, err
	}

	result.Scan(&res.Id, &res.Passwword)

	return res, nil
}

func (repo *UserRepositoryImpl) GetUsersByUsername(ctx context.Context, db *pgxpool.Pool, user entity.User) (entity.ListUser, error) {
	var (
		res      entity.User
		listUser entity.ListUser = make(entity.ListUser, 0)
	)

	SQL := `SELECT id, username, fullname FROM users WHERE username LIKE $1`
	varArgs := []interface{}{
		"%" + user.Username + "%",
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&res.Id, &res.Username, &res.FullName)
		if err != nil {
			return nil, err
		}
		listUser = append(listUser, res)
	}

	return listUser, nil
}
