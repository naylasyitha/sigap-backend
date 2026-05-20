package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GrowthRecord struct {
	ID                uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	ChildID           uuid.UUID `gorm:"type:char(36);not null" json:"child_id"`
	Weight            float64   `gorm:"type:decimal(5,2);not null" json:"weight"`
	Height            float64   `gorm:"type:decimal(5,2);not null" json:"height"`
	HeadCircumference float64   `gorm:"type:decimal(5,2)" json:"head_circumference"`
	AgeMonths         int       `gorm:"not null" json:"age_months"`
	MeasurementDate   time.Time `gorm:"not null" json:"measurement_date"`
	Notes             string    `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	Child Child `gorm:"foreignKey:ChildID" json:"child,omitempty"`
}

func (g *GrowthRecord) BeforeCreate(tx *gorm.DB) error {
	g.ID = uuid.New()
	return nil
}
