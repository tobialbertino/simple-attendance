package app

import (
	attendanceHandler "simple-attendance/internal/attendance/delivery/http"
	attendanceRepository "simple-attendance/internal/attendance/repository/postgres"
	attendanceUseCase "simple-attendance/internal/attendance/usecase"
	authHandler "simple-attendance/internal/auth/delivery/http"
	authRepository "simple-attendance/internal/auth/repository/postgres"
	authUseCase "simple-attendance/internal/auth/usecase"
	userHandler "simple-attendance/internal/user/delivery/http"
	userRepository "simple-attendance/internal/user/repository/postgres"
	userUseCase "simple-attendance/internal/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(app *fiber.App, DB *pgxpool.Pool, validate *validator.Validate) {
	// user setup
	userRepo := userRepository.NewUserRepository()
	userUc := userUseCase.NewUserUseCase(userRepo, DB, validate)
	userHandler := userHandler.NewHandler(userUc)
	userHandler.Route(app)

	// Auth setup
	authRepo := authRepository.NewAuthRepository()
	authUc := authUseCase.NewAuthUseCase(userUc, authRepo, DB, validate)
	authHandler := authHandler.NewHandler(authUc)
	authHandler.Route(app)

	//attend setup
	attendanceRepo := attendanceRepository.NewAttendanceRepository()
	attendanceUc := attendanceUseCase.NewAttendanceUseCase(attendanceRepo, DB, validate)
	attendanceHandler := attendanceHandler.NewHandler(attendanceUc)
	attendanceHandler.Route(app)

}
