package repository

import (
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GrowthRecordRepository interface {
	Create(record *entity.GrowthRecord) error
	FindAllByChildID(childID uuid.UUID) ([]entity.GrowthRecord, error)
	GetLatestByChildID(childID uuid.UUID) (*entity.GrowthRecord, error)
}

type growthRecordRepository struct {
	db *gorm.DB
}

func NewGrowthRecordRepository(db *gorm.DB) GrowthRecordRepository {
	return &growthRecordRepository{db: db}
}

func (r *growthRecordRepository) Create(record *entity.GrowthRecord) error {
	return r.db.Create(record).Error
}

func (r *growthRecordRepository) FindAllByChildID(childID uuid.UUID) ([]entity.GrowthRecord, error) {
	var records []entity.GrowthRecord

	err := r.db.
		Where("child_id = ?", childID).
		Order("measurement_date ASC").
		Find(&records).Error

	return records, err
}

func (r *growthRecordRepository) GetLatestByChildID(childID uuid.UUID) (*entity.GrowthRecord, error) {
	var record entity.GrowthRecord

	err := r.db.
		Where("child_id = ?", childID).
		Order("measurement_date DESC").
		First(&record).Error

	return &record, err
}
