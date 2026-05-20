package repository

import (
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *entity.Schedule) error
	FindAll() ([]entity.Schedule, error)
	FindAllByChildID(childID uuid.UUID) ([]entity.Schedule, error)
	FindByID(id uuid.UUID) (*entity.Schedule, error)
	Update(schedule *entity.Schedule) error
	Delete(id uuid.UUID) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *entity.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	err := r.db.
		Preload("Child").
		Order("schedule_date ASC").
		Find(&schedules).Error

	return schedules, err
}

func (r *scheduleRepository) FindAllByChildID(childID uuid.UUID) ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	err := r.db.
		Preload("Child").
		Where("child_id = ?", childID).
		Order("schedule_date ASC").
		Find(&schedules).Error

	return schedules, err
}

func (r *scheduleRepository) FindByID(id uuid.UUID) (*entity.Schedule, error) {
	var schedule entity.Schedule

	err := r.db.
		Preload("Child").
		Where("id = ?", id).
		First(&schedule).Error

	return &schedule, err
}

func (r *scheduleRepository) Update(schedule *entity.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Schedule{}, "id = ?", id).Error
}
