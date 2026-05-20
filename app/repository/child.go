package repository

import (
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChildRepository interface {
	Create(child *entity.Child) error
	FindAllByUserID(userID uuid.UUID) ([]entity.Child, error)
	FindByID(id uuid.UUID) (*entity.Child, error)
	Update(child *entity.Child) error
	Delete(id uuid.UUID) error
}

type childRepository struct {
	db *gorm.DB
}

func NewChildRepository(db *gorm.DB) ChildRepository {
	return &childRepository{db: db}
}

func (r *childRepository) Create(child *entity.Child) error {
	return r.db.Create(child).Error
}

func (r *childRepository) FindAllByUserID(userID uuid.UUID) ([]entity.Child, error) {
	var children []entity.Child
	err := r.db.Where("user_id = ?", userID).Find(&children).Error
	return children, err
}

func (r *childRepository) FindByID(id uuid.UUID) (*entity.Child, error) {
	var child entity.Child
	err := r.db.Where("id = ?", id).First(&child).Error
	return &child, err
}

func (r *childRepository) Update(child *entity.Child) error {
	return r.db.Save(child).Error
}

func (r *childRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Child{}, "id = ?", id).Error
}
