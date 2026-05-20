package rest

import (
	"net/http"

	"sigap-backend/app/usecase"
	"sigap-backend/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ScheduleHandler struct {
	usecase  usecase.ScheduleUsecase
	validate *validator.Validate
}

func NewScheduleHandler(routerGroup fiber.Router, scheduleUsecase usecase.ScheduleUsecase) {
	handler := ScheduleHandler{
		usecase:  scheduleUsecase,
		validate: validator.New(),
	}

	schedule := routerGroup.Group("/schedules")
	schedule.Post("/", handler.Create)
	schedule.Get("/", handler.FindAll)
	schedule.Get("/child/:child_id", handler.FindAllByChildID)
	schedule.Get("/:id", handler.FindByID)
	schedule.Put("/:id", handler.Update)
	schedule.Delete("/:id", handler.Delete)
}

func getUserID(ctx *fiber.Ctx) (uuid.UUID, error) {
	claims := ctx.Locals("user").(jwt.MapClaims)
	return uuid.Parse(claims["user_id"].(string))
}

func (h *ScheduleHandler) Create(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"errors":  "Invalid token",
		})
	}

	var req dto.CreateScheduleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
			"errors":  "Request format is not valid",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	res, err := h.usecase.Create(userID, req)
	if err != nil {
		status := fiber.StatusBadRequest
		if err.Error() == "unauthorized access to child profile" {
			status = fiber.StatusForbidden
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Create schedule failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Schedule created successfully",
		"data":    res,
	})
}

func (h *ScheduleHandler) FindAll(ctx *fiber.Ctx) error {
	res, err := h.usecase.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch schedules failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Schedules fetched successfully",
		"data":    res,
	})
}

func (h *ScheduleHandler) FindAllByChildID(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"errors":  "Invalid token",
		})
	}

	childID, err := uuid.Parse(ctx.Params("child_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid child id",
			"errors":  "Child id is not valid",
		})
	}

	res, err := h.usecase.FindAllByChildID(userID, childID)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "unauthorized access to child profile" {
			status = fiber.StatusForbidden
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Fetch schedules failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Schedules fetched successfully",
		"data":    res,
	})
}

func (h *ScheduleHandler) FindByID(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"errors":  "Invalid token",
		})
	}

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid schedule id",
			"errors":  "Schedule id is not valid",
		})
	}

	res, err := h.usecase.FindByID(userID, id)
	if err != nil {
		status := fiber.StatusNotFound
		if err.Error() == "unauthorized access to schedule" {
			status = fiber.StatusForbidden
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Schedule not found",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Schedule fetched successfully",
		"data":    res,
	})
}

func (h *ScheduleHandler) Update(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"errors":  "Invalid token",
		})
	}

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid schedule id",
			"errors":  "Schedule id is not valid",
		})
	}

	var req dto.UpdateScheduleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
			"errors":  "Request format is not valid",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	res, err := h.usecase.Update(userID, id, req)
	if err != nil {
		status := fiber.StatusBadRequest
		if err.Error() == "unauthorized access to schedule" {
			status = fiber.StatusForbidden
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Update schedule failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Schedule updated successfully",
		"data":    res,
	})
}

func (h *ScheduleHandler) Delete(ctx *fiber.Ctx) error {
	userID, err := getUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"errors":  "Invalid token",
		})
	}

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid schedule id",
			"errors":  "Schedule id is not valid",
		})
	}

	if err := h.usecase.Delete(userID, id); err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "unauthorized access to schedule" {
			status = fiber.StatusForbidden
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Delete schedule failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Schedule deleted successfully",
		"data":    fiber.Map{},
	})
}
