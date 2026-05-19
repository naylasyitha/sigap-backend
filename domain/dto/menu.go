package dto

import (
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Name       string                 `json:"name" validate:"required"`
	ImageURL   string                 `json:"image_url" validate:"required"`
	AgeRange   entity.MpasiAgeRange   `json:"age_range" validate:"required"`
	MealType   entity.MealType        `json:"meal_type" validate:"required"`
	Difficulty entity.DifficultyLevel `json:"difficulty" validate:"required"`
	Duration   string                 `json:"duration" validate:"required"`
	Portion    int                    `json:"portion" validate:"required"`
	Calories   int                    `json:"calories" validate:"required"`

	Ingredients []CreateIngredientRequest `json:"ingredients" validate:"required,dive"`
	Steps       []CreateStepRequest       `json:"steps" validate:"required,dive"`
	Nutritions  []CreateNutritionRequest  `json:"nutritions" validate:"required,dive"`
}

type CreateIngredientRequest struct {
	Name   string `json:"name" validate:"required"`
	Amount string `json:"amount" validate:"required"`
}

type CreateStepRequest struct {
	StepNumber  int    `json:"step_number" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateNutritionRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type MenuResponse struct {
	ID         uuid.UUID              `json:"id"`
	Name       string                 `json:"name"`
	ImageURL   string                 `json:"image_url"`
	AgeRange   entity.MpasiAgeRange   `json:"age_range"`
	MealType   entity.MealType        `json:"meal_type"`
	Difficulty entity.DifficultyLevel `json:"difficulty"`
	Duration   string                 `json:"duration"`
	Portion    int                    `json:"portion"`
	Calories   int                    `json:"calories"`

	Ingredients []IngredientResponse `json:"ingredients,omitempty"`
	Steps       []StepResponse       `json:"steps,omitempty"`
	Nutritions  []NutritionResponse  `json:"nutritions,omitempty"`
}

type IngredientResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount string    `json:"amount"`
}

type StepResponse struct {
	ID          uuid.UUID `json:"id"`
	StepNumber  int       `json:"step_number"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type NutritionResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}
