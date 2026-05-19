package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MpasiStep struct {
	ID          uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	MenuID      uuid.UUID `gorm:"type:char(36);not null" json:"menu_id"`
	StepNumber  int       `gorm:"not null" json:"step_number"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Menu MpasiMenu `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
}

func (u *MpasiStep) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
