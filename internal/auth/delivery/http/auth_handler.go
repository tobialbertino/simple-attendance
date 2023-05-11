package http

import (
	"simple-attendance/internal/auth/models/domain"
	"simple-attendance/internal/auth/usecase"
	"simple-attendance/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUC,
	}
}

func (h *AuthHandler) Route(app *fiber.App) {
	// notes
	auth := app.Group("/auth")

	// auth user
	auth.Post("", h.postAuthenticationHandler)
	auth.Put("", h.putAuthenticationHandler)
	auth.Delete("", h.deleteAuthenticationHandler)
}

func (h *AuthHandler) postAuthenticationHandler(c *fiber.Ctx) error {
	var req domain.ReqLoginUser

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.AddRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Authentication berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AuthHandler) putAuthenticationHandler(c *fiber.Ctx) error {
	var req domain.ReqRefreshToken

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.VerifyRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Access Token berhasil diperbarui",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AuthHandler) deleteAuthenticationHandler(c *fiber.Ctx) error {
	var req domain.ReqRefreshToken

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.DeleteRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Refresh token berhasil dihapus",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}
