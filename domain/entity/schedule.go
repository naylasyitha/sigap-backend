package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleType string

const (
	WeightCheck  ScheduleType = "WEIGHT_CHECK"
	Immunization ScheduleType = "IMMUNIZATION"
	Vitamin      ScheduleType = "VITAMIN"
	Consultation ScheduleType = "CONSULTATION"
)

type ScheduleStatus string

const (
	Pending   ScheduleStatus = "PENDING"
	Completed ScheduleStatus = "COMPLETED"
)

type Schedule struct {
	ID           uuid.UUID      `gorm:"primaryKey;type:char(36)" json:"id"`
	ChildID      uuid.UUID      `gorm:"type:char(36);not null" json:"child_id"`
	Title        string         `gorm:"type:varchar(100);not null" json:"title"`
	Type         ScheduleType   `gorm:"type:varchar(50);not null" json:"type"`
	Status       ScheduleStatus `gorm:"type:varchar(50);not null" json:"status"`
	ScheduleDate time.Time      `gorm:"not null" json:"schedule_date"`
	Location     string         `gorm:"type:varchar(150)" json:"location"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`

	Child Child `gorm:"foreignKey:ChildID" json:"child,omitempty"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
