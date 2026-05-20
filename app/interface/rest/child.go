package rest

import (
	"net/http"

	"sigap-backend/app/usecase"
	"sigap-backend/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ChildHandler struct {
	usecase  usecase.ChildUsecase
	validate *validator.Validate
}

func NewChildHandler(routerGroup fiber.Router, childUsecase usecase.ChildUsecase, mid fiber.Handler) {
	handler := ChildHandler{
		usecase:  childUsecase,
		validate: validator.New(),
	}

	child := routerGroup.Group("/children")

	child.Post("/", mid, handler.Create)
	child.Get("/", mid, handler.FindAll)
	child.Patch("/:id", mid, handler.Update)
	child.Delete("/:id", mid, handler.Delete)
}

func (h *ChildHandler) Create(ctx *fiber.Ctx) error {
	var req dto.CreateChildRequest
	userID := ctx.Locals("user_id").(uuid.UUID)

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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Create child profile failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Child profile created successfully",
		"data":    res,
	})
}

func (h *ChildHandler) FindAll(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uuid.UUID)

	res, err := h.usecase.FindAllByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch child profiles failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Child profiles fetched successfully",
		"data":    res,
	})
}

func (h *ChildHandler) Update(ctx *fiber.Ctx) error {
	var req dto.UpdateChildRequest
	userID := ctx.Locals("user_id").(uuid.UUID)

	childID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid child id",
			"errors":  "Child id is not valid",
		})
	}

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

	res, err := h.usecase.Update(userID, childID, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Update child profile failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Child profile updated successfully",
		"data":    res,
	})
}

func (h *ChildHandler) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uuid.UUID)

	childID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid child id",
			"errors":  "Child id is not valid",
		})
	}

	if err := h.usecase.Delete(userID, childID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Delete child profile failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Child profile deleted successfully",
	})
}
