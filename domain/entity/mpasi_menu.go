package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MpasiAgeRange string

const (
	Age6To8   MpasiAgeRange = "6_8_MONTHS"
	Age9To11  MpasiAgeRange = "9_11_MONTHS"
	Age12To24 MpasiAgeRange = "12_24_MONTHS"
	Age24Plus MpasiAgeRange = "24_PLUS_MONTHS"
)

type MealType string

const (
	Snack      MealType = "SNACK"
	Breakfast  MealType = "BREAKFAST"
	Lunch      MealType = "LUNCH"
	Dinner     MealType = "DINNER"
	FingerFood MealType = "FINGER_FOOD"
)

type DifficultyLevel string

const (
	Easy   DifficultyLevel = "EASY"
	Medium DifficultyLevel = "MEDIUM"
	Hard   DifficultyLevel = "HARD"
)

type MpasiMenu struct {
	ID         uuid.UUID       `gorm:"primaryKey;type:char(36)" json:"id"`
	Name       string          `gorm:"type:varchar(100);not null" json:"name"`
	ImageURL   string          `gorm:"type:text" json:"image_url"`
	AgeRange   MpasiAgeRange   `gorm:"type:varchar(50);not null" json:"age_range"`
	MealType   MealType        `gorm:"type:varchar(50);not null" json:"meal_type"`
	Difficulty DifficultyLevel `gorm:"type:varchar(50);not null" json:"difficulty"`
	Duration   string          `gorm:"type:varchar(50)" json:"duration"`
	Portion    int             `gorm:"not null" json:"portion"`
	Calories   int             `gorm:"not null" json:"calories"`
	// IsRecommended bool `gorm:"default:false" json:"is_recommended"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Ingredients []MpasiIngredient `gorm:"foreignKey:MenuID" json:"ingredients,omitempty"`
	Steps       []MpasiStep       `gorm:"foreignKey:MenuID" json:"steps,omitempty"`
	Nutritions  []MpasiNutrition  `gorm:"foreignKey:MenuID" json:"nutritions,omitempty"`
}

func (u *MpasiMenu) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
