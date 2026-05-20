package dto

import (
	"time"

	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type CreateScheduleRequest struct {
	ChildID      uuid.UUID           `json:"child_id" validate:"required"`
	Title        string              `json:"title" validate:"required"`
	Type         entity.ScheduleType `json:"type" validate:"required"`
	ScheduleDate time.Time           `json:"schedule_date" validate:"required"`
	Location     string              `json:"location"`
}

type UpdateScheduleRequest struct {
	ChildID      uuid.UUID             `json:"child_id" validate:"required"`
	Title        string                `json:"title" validate:"required"`
	Type         entity.ScheduleType   `json:"type" validate:"required"`
	Status       entity.ScheduleStatus `json:"status" validate:"required"`
	ScheduleDate time.Time             `json:"schedule_date" validate:"required"`
	Location     string                `json:"location"`
}

type ScheduleResponse struct {
	ID           uuid.UUID             `json:"id"`
	ChildID      uuid.UUID             `json:"child_id"`
	ChildName    string                `json:"child_name,omitempty"`
	Title        string                `json:"title"`
	Type         entity.ScheduleType   `json:"type"`
	Status       entity.ScheduleStatus `json:"status"`
	ScheduleDate time.Time             `json:"schedule_date"`
	Location     string                `json:"location"`
}
