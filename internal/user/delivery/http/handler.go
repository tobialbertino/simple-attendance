package http

import (
	"simple-attendance/internal/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserUseCase usecase.UserUseCase
}

func NewHandler(userUC usecase.UserUseCase) *Handler {
	return &Handler{
		UserUseCase: userUC,
	}
}

func (h *Handler) Route(app *fiber.App) {
	// user
	u := NewUserHandler(h.UserUseCase)
	u.Route(app)
}
