package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateGrowthRecordRequest struct {
	ChildID           uuid.UUID `json:"child_id" validate:"required"`
	Weight            float64   `json:"weight" validate:"required"`
	Height            float64   `json:"height" validate:"required"`
	HeadCircumference float64   `json:"head_circumference"`
	MeasurementDate   time.Time `json:"measurement_date" validate:"required"`
}

type GrowthRecordResponse struct {
	ID                uuid.UUID `json:"id"`
	ChildID           uuid.UUID `json:"child_id"`
	Weight            float64   `json:"weight"`
	Height            float64   `json:"height"`
	HeadCircumference float64   `json:"head_circumference"`
	MeasurementDate   time.Time `json:"measurement_date"`
}
