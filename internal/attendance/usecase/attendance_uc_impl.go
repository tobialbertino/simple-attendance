package usecase

import (
	"context"
	"simple-attendance/exception"
	"simple-attendance/internal/attendance/models/domain"
	"simple-attendance/internal/attendance/models/entity"
	"simple-attendance/internal/attendance/repository/postgres"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceUseCaseImpl struct {
	AttendanceRepository postgres.AttendanceRepository
	DB                   *pgxpool.Pool
	Validate             *validator.Validate
}

func NewAttendanceUseCase(attendanceRepo postgres.AttendanceRepository, DB *pgxpool.Pool, validate *validator.Validate) AttendanceUseCase {
	return &AttendanceUseCaseImpl{
		AttendanceRepository: attendanceRepo,
		DB:                   DB,
		Validate:             validate,
	}
}

// GetAttendanceById implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) GetAttendanceById(req domain.ReqId) (domain.ResAttendance, error) {
	var response domain.ResAttendance

	err := useCase.Validate.Struct(req)
	if err != nil {
		return response, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		Id: req.Id,
	}

	responseRepo, err := useCase.AttendanceRepository.VerifyUser(context.Background(), useCase.DB, request)
	if err != nil {
		return response, err
	}

	response = responseRepo.ToDomain()

	return response, nil
}

// AddAttendance implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) AddAttendance(req domain.ReqAddAttendance) (domain.IdAttendance, error) {
	var response domain.IdAttendance

	err := useCase.Validate.Struct(req)
	if err != nil {
		return response, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		Id:     uuid.New().String(),
		UserId: req.UserId,
	}

	_, err = useCase.AttendanceRepository.AddAttendance(context.Background(), useCase.DB, request)
	if err != nil {
		return response, err
	}

	response.ID = request.Id

	return response, nil
}

// CheckInById implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) CheckInById(req domain.ReqCheckIn) (domain.RowsAffected, error) {
	var response domain.RowsAffected

	err := useCase.Validate.Struct(req)
	if err != nil {
		return domain.RowsAffected{}, exception.NewClientError(err.Error(), 400)
	}

	timeNow := time.Now().Unix()
	request := entity.Attendance{
		CheckIn: &timeNow,
		Id:      req.Id,
	}

	i, err := useCase.AttendanceRepository.CheckInById(context.Background(), useCase.DB, request)
	if err != nil {
		return domain.RowsAffected{}, err
	}

	response.RowsAffected = i

	return response, nil
}

// CheckOutById implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) CheckOutById(req domain.ReqCheckOut) (domain.RowsAffected, error) {
	var response domain.RowsAffected

	err := useCase.Validate.Struct(req)
	if err != nil {
		return domain.RowsAffected{}, exception.NewClientError(err.Error(), 400)
	}

	timeNow := time.Now().Unix()
	request := entity.Attendance{
		CheckOut: &timeNow,
		Id:       req.Id,
	}

	i, err := useCase.AttendanceRepository.CheckOutById(context.Background(), useCase.DB, request)
	if err != nil {
		return domain.RowsAffected{}, err
	}

	response.RowsAffected = i

	return response, nil
}

// DeleteById implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) DeleteById(req domain.ReqId) (domain.RowsAffected, error) {
	var response domain.RowsAffected

	err := useCase.Validate.Struct(req)
	if err != nil {
		return domain.RowsAffected{}, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		Id: req.Id,
	}

	i, err := useCase.AttendanceRepository.DeleteById(context.Background(), useCase.DB, request)
	if err != nil {
		return domain.RowsAffected{}, err
	}

	response.RowsAffected = i

	return response, nil
}

// GetAllByUserId implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) GetAllByUserId(req domain.ReqGetAllByUserId) ([]domain.ResAttendance, error) {
	var response []domain.ResAttendance

	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		UserId: req.UserId,
	}

	res, err := useCase.AttendanceRepository.GetAllByUserId(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response = res.ToDomain()

	return response, nil
}

// UpdateActivityById implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) UpdateActivityById(req domain.ReqUpdate) (domain.RowsAffected, error) {
	var response domain.RowsAffected

	err := useCase.Validate.Struct(req)
	if err != nil {
		return domain.RowsAffected{}, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		Activity: req.Activity,
		Id:       req.Id,
		Location: req.Location,
	}

	i, err := useCase.AttendanceRepository.UpdateActivityById(context.Background(), useCase.DB, request)
	if err != nil {
		return domain.RowsAffected{}, err
	}

	response.RowsAffected = i

	return response, nil
}

// VerifyUser implements AttendanceUseCase
func (useCase *AttendanceUseCaseImpl) VerifyUser(req domain.ReqAddAttendance) (bool, error) {
	err := useCase.Validate.Struct(req)
	if err != nil {
		return false, exception.NewClientError(err.Error(), 400)
	}

	request := entity.Attendance{
		UserId: req.UserId,
		Id:     req.Id,
	}

	result, err := useCase.AttendanceRepository.VerifyUser(context.Background(), useCase.DB, request)
	if err != nil {
		return false, err
	}
	if result.UserId != req.UserId {
		return false, exception.NewClientError("Forbidden", 403)
	}

	return true, nil
}
