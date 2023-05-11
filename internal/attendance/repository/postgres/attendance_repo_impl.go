package postgres

import (
	"context"
	"simple-attendance/internal/attendance/models/entity"
	"simple-attendance/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepositoryImpl struct {
}

func NewAttendanceRepository() AttendanceRepository {
	return &AttendanceRepositoryImpl{}
}

// AddAttendancer implements AttendanceRepository
func (repo *AttendanceRepositoryImpl) AddAttendance(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error) {
	SQL := `INSERT INTO attendances(id, user_id) VALUES($1, $2) RETURNING id`
	varArgs := []interface{}{
		data.Id,
		data.UserId,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repo *AttendanceRepositoryImpl) UpdateActivityById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error) {
	SQL := `UPDATE attendances SET activity = $1, location = $2 WHERE id = $3`
	varArgs := []interface{}{
		data.Activity,
		data.Location,
		data.Id,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repo *AttendanceRepositoryImpl) CheckInById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error) {
	SQL := `UPDATE attendances SET check_in = $1 WHERE id = $2`
	varArgs := []interface{}{
		data.CheckIn,
		data.Id,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repo *AttendanceRepositoryImpl) CheckOutById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error) {
	SQL := `UPDATE attendances SET check_out = $1 WHERE id = $2`
	varArgs := []interface{}{
		data.CheckOut,
		data.Id,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repo *AttendanceRepositoryImpl) DeleteById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error) {
	SQL := `DELETE FROM attendances WHERE id = $1`
	varArgs := []interface{}{
		data.Id,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repo *AttendanceRepositoryImpl) GetAllByUserId(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (entity.ListAttendances, error) {
	var listResult entity.ListAttendances
	var result entity.Attendance

	SQL := `SELECT attendances.* FROM attendances WHERE attendances.user_id = $1`
	varArgs := []interface{}{
		data.UserId,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return listResult, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.Id, &result.UserId, &result.Activity, &result.Location, &result.CheckIn, &result.CheckOut)
		if err != nil {
			return nil, err
		}
		listResult = append(listResult, result)
	}

	return listResult, nil
}

func (repo *AttendanceRepositoryImpl) VerifyUser(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (entity.Attendance, error) {
	var result entity.Attendance

	SQL := `SELECT * FROM attendances WHERE id = $1`
	varArgs := []interface{}{
		data.Id,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return result, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&result.Id, &result.UserId, &result.Activity, &result.Location, &result.CheckIn, &result.CheckOut)

	return result, nil
}
