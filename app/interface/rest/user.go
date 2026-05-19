package rest

import (
	"net/http"
	"sigap-backend/app/usecase"
	"sigap-backend/domain/dto"
	"sigap-backend/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase  usecase.UserUsecase
	validate *validator.Validate
}

func NewUserHandler(routerGroup fiber.Router, userUsecase usecase.UserUsecase) {
	handler := UserHandler{
		usecase:  userUsecase,
		validate: validator.New(),
	}

	auth := routerGroup.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Put("/profile", middleware.AuthMiddleware, handler.UpdateProfile)
	auth.Post("/logout", middleware.AuthMiddleware, handler.Logout)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest

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

	res, err := h.usecase.Register(req)
	if err != nil {
		status := fiber.StatusBadRequest
		if err.Error() == "email already in use" {
			status = fiber.StatusConflict
		}
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Registration failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": res.Message,
		"data": fiber.Map{
			"user": fiber.Map{
				"id":    res.User.ID,
				"name":  res.User.Name,
				"email": res.User.Email,
			},
		},
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest

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

	res, err := h.usecase.Login(req)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Login failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": res.Message,
		"data": fiber.Map{
			"token": res.Token,
			"user": fiber.Map{
				"id":    res.User.ID,
				"name":  res.User.Name,
				"email": res.User.Email,
			},
		},
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	var req dto.UpdateProfileRequest

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

	res, err := h.usecase.UpdateProfile(userID, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Update profile failed",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Profile updated successfully",
		"data": fiber.Map{
			"user": res,
		},
	})
}

func (h *UserHandler) Logout(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Logout successful",
	})
}
