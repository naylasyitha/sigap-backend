package rest

import (
	"net/http"

	"sigap-backend/app/usecase"
	"sigap-backend/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GrowthRecordHandler struct {
	usecase  usecase.GrowthRecordUsecase
	validate *validator.Validate
}

func NewGrowthRecordHandler(routerGroup fiber.Router, growthUsecase usecase.GrowthRecordUsecase, mid fiber.Handler) {
	handler := GrowthRecordHandler{
		usecase:  growthUsecase,
		validate: validator.New(),
	}

	growth := routerGroup.Group("/growth-records")

	growth.Post("/", mid, handler.Create)
	growth.Get("/child/:child_id", mid, handler.FindAllByChildID)
	growth.Get("/child/:child_id/latest", mid, handler.GetLatestByChildID)
}

func (h *GrowthRecordHandler) Create(ctx *fiber.Ctx) error {
	var req dto.CreateGrowthRecordRequest

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

	res, err := h.usecase.Create(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Create growth record failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Growth record created successfully",
		"data":    res,
	})
}

func (h *GrowthRecordHandler) FindAllByChildID(ctx *fiber.Ctx) error {
	childID, err := uuid.Parse(ctx.Params("child_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid child id",
			"errors":  "Child id is not valid",
		})
	}

	res, err := h.usecase.FindAllByChildID(childID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch growth records failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Growth records fetched successfully",
		"data":    res,
	})
}

func (h *GrowthRecordHandler) GetLatestByChildID(ctx *fiber.Ctx) error {
	childID, err := uuid.Parse(ctx.Params("child_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid child id",
			"errors":  "Child id is not valid",
		})
	}

	res, err := h.usecase.GetLatestByChildID(childID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Growth record not found",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Latest growth record fetched successfully",
		"data":    res,
	})
}
