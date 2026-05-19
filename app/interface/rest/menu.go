package rest

import (
	"net/http"
	"sigap-backend/app/usecase"
	"sigap-backend/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuHandler struct {
	usecase  usecase.MenuUsecase
	validate *validator.Validate
}

func NewMenuHandler(routerGroup fiber.Router, menuUsecase usecase.MenuUsecase) {
	handler := MenuHandler{
		usecase:  menuUsecase,
		validate: validator.New(),
	}

	menu := routerGroup.Group("/menus")

	menu.Post("/", handler.Create)
	menu.Get("/", handler.FindAll)
	menu.Get("/:id", handler.FindByID)
}

func (h *MenuHandler) Create(ctx *fiber.Ctx) error {
	var req dto.CreateMenuRequest

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
			"message": "Create menu failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Menu created successfully",
		"data":    res,
	})
}

func (h *MenuHandler) FindAll(ctx *fiber.Ctx) error {
	res, err := h.usecase.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch menus failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Menus fetched successfully",
		"data":    res,
	})
}

func (h *MenuHandler) FindByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid menu id",
			"errors":  "Menu id is not valid",
		})
	}

	res, err := h.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Menu not found",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Menu fetched successfully",
		"data":    res,
	})
}
