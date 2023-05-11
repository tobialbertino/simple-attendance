package postgres

import (
	"context"
	"simple-attendance/internal/attendance/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepository interface {
	AddAttendance(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error)
	UpdateActivityById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error)
	CheckInById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error)
	CheckOutById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error)
	DeleteById(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (int64, error)
	GetAllByUserId(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (entity.ListAttendances, error)

	// authorization
	// VerifyUser can get attendance return
	VerifyUser(ctx context.Context, db *pgxpool.Pool, data entity.Attendance) (entity.Attendance, error)
}
