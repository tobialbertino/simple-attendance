package http

import (
	"simple-attendance/internal/attendance/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AttendanceUseCase usecase.AttendanceUseCase
}

func NewHandler(attendUC usecase.AttendanceUseCase) *Handler {
	return &Handler{
		AttendanceUseCase: attendUC,
	}
}

func (h *Handler) Route(app *fiber.App) {
	// attend
	u := NewAttendanceHandler(h.AttendanceUseCase)
	u.Route(app)
}
