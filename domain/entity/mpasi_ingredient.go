package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MpasiIngredient struct {
	ID        uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	MenuID    uuid.UUID `gorm:"type:char(36);not null" json:"menu_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Amount    string    `gorm:"type:varchar(50)" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Menu MpasiMenu `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
}

func (u *MpasiIngredient) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
