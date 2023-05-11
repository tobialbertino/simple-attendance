package usecase

import "simple-attendance/internal/attendance/models/domain"

type AttendanceUseCase interface {
	AddAttendance(req domain.ReqAddAttendance) (domain.IdAttendance, error)
	UpdateActivityById(req domain.ReqUpdate) (domain.RowsAffected, error)
	CheckInById(req domain.ReqCheckIn) (domain.RowsAffected, error)
	CheckOutById(req domain.ReqCheckOut) (domain.RowsAffected, error)
	DeleteById(req domain.ReqId) (domain.RowsAffected, error)
	GetAllByUserId(req domain.ReqGetAllByUserId) ([]domain.ResAttendance, error)
	GetAttendanceById(req domain.ReqId) (domain.ResAttendance, error)

	VerifyUser(req domain.ReqAddAttendance) (bool, error)
}
