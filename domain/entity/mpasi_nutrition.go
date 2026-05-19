package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MpasiNutrition struct {
	ID        uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	MenuID    uuid.UUID `gorm:"type:char(36);not null" json:"menu_id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	Value     string    `gorm:"type:varchar(50);not null" json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Menu MpasiMenu `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
}

func (u *MpasiNutrition) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
