package http

import (
	"simple-attendance/internal/attendance/models/domain"
	"simple-attendance/internal/attendance/usecase"
	"simple-attendance/pkg/helper"
	"simple-attendance/pkg/middleware"
	"simple-attendance/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	AttendanceUseCase usecase.AttendanceUseCase
}

func NewAttendanceHandler(attendanceoUseCase usecase.AttendanceUseCase) *AttendanceHandler {
	return &AttendanceHandler{
		AttendanceUseCase: attendanceoUseCase,
	}
}

func (h *AttendanceHandler) Route(app *fiber.App) {
	// attendance
	g := app.Group("/attend", middleware.ProtectedJWT())

	g.Post("", h.Add)
	g.Put("/activity/:id", h.UpdateActivityById)
	g.Put("/check-in/:id", h.CheckInById)
	g.Put("/check-out/:id", h.CheckOutById)
	g.Delete(":id", h.DeleteById)
	g.Get("", h.GetAllByUserId)
	g.Get(":id", h.GetAttendanceById)
}

func (h *AttendanceHandler) Add(c *fiber.Ctx) error {
	var request domain.ReqAddAttendance

	userId := helper.GetIDUserFromToken(c)
	request.UserId = userId

	result, err := h.AttendanceUseCase.AddAttendance(request)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "absen berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

// Untuk authorization attendance user id
func (h *AttendanceHandler) VerifyUser(c *fiber.Ctx) error {
	var request domain.ReqAddAttendance

	id := c.Params("id")

	userId := helper.GetIDUserFromToken(c)
	request.UserId = userId
	request.Id = id

	_, err := h.AttendanceUseCase.VerifyUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (h *AttendanceHandler) UpdateActivityById(c *fiber.Ctx) error {
	// verify user
	err := h.VerifyUser(c)
	if err != nil {
		return err
	}

	var request domain.ReqUpdate

	id := c.Params("id")
	request.Id = id

	err = c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.UpdateActivityById(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "aktivitas berhasul di update",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AttendanceHandler) CheckInById(c *fiber.Ctx) error {
	// verify user
	err := h.VerifyUser(c)
	if err != nil {
		return err
	}

	var request domain.ReqCheckIn

	id := c.Params("id")
	request.Id = id

	err = c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.CheckInById(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "check-in berhasul di update",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AttendanceHandler) CheckOutById(c *fiber.Ctx) error {
	// verify user
	err := h.VerifyUser(c)
	if err != nil {
		return err
	}

	var request domain.ReqCheckOut

	id := c.Params("id")
	request.Id = id

	err = c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.CheckOutById(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "check-out berhasul di update",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AttendanceHandler) DeleteById(c *fiber.Ctx) error {
	// verify user
	err := h.VerifyUser(c)
	if err != nil {
		return err
	}

	var request domain.ReqId

	id := c.Params("id")
	request.Id = id

	err = c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.DeleteById(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "attendance berhasul di hapus",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AttendanceHandler) GetAllByUserId(c *fiber.Ctx) error {
	var request domain.ReqGetAllByUserId

	userId := helper.GetIDUserFromToken(c)
	request.UserId = userId

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.GetAllByUserId(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "semua aktivitas berhasul di dapat",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AttendanceHandler) GetAttendanceById(c *fiber.Ctx) error {
	var request domain.ReqId

	id := c.Params("id")
	request.Id = id

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.AttendanceUseCase.GetAttendanceById(request)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Succes Get Attendance By Id",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}
