package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type Child struct {
	ID          uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	UserID      uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Gender      Gender    `gorm:"type:varchar(20);not null" json:"gender"`
	BirthDate   time.Time `gorm:"not null" json:"birth_date"`
	BirthWeight float64   `gorm:"type:decimal(5,2)" json:"birth_weight"`
	BirthHeight float64   `gorm:"type:decimal(5,2)" json:"birth_height"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User          User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	GrowthRecords []GrowthRecord `gorm:"foreignKey:ChildID" json:"growth_records,omitempty"`
	Schedules     []Schedule     `gorm:"foreignKey:ChildID" json:"schedules,omitempty"`
}

func (c *Child) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
