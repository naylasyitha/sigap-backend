package dto

import (
	"time"

	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type CreateChildRequest struct {
	Name        string        `json:"name" validate:"required"`
	Gender      entity.Gender `json:"gender" validate:"required"`
	BirthDate   time.Time     `json:"birth_date" validate:"required"`
	BirthWeight float64       `json:"birth_weight" validate:"required"`
	BirthHeight float64       `json:"birth_height" validate:"required"`
}

type UpdateChildRequest struct {
	Name        string        `json:"name" validate:"required"`
	Gender      entity.Gender `json:"gender" validate:"required"`
	BirthDate   time.Time     `json:"birth_date" validate:"required"`
	BirthWeight float64       `json:"birth_weight" validate:"required"`
	BirthHeight float64       `json:"birth_height" validate:"required"`
}

type ChildResponse struct {
	ID          uuid.UUID     `json:"id"`
	UserID      uuid.UUID     `json:"user_id"`
	Name        string        `json:"name"`
	Gender      entity.Gender `json:"gender"`
	BirthDate   time.Time     `json:"birth_date"`
	BirthWeight float64       `json:"birth_weight"`
	BirthHeight float64       `json:"birth_height"`
}
