package http

import (
	"simple-attendance/internal/auth/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AuthUseCase usecase.AuthUseCase
}

func NewHandler(authUC usecase.AuthUseCase) *Handler {
	return &Handler{
		AuthUseCase: authUC,
	}
}

func (h *Handler) Route(app *fiber.App) {
	// auth
	u := NewAuthHandler(h.AuthUseCase)
	u.Route(app)
}
