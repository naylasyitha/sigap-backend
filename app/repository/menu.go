package repository

import (
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *entity.MpasiMenu) error
	FindAll() ([]entity.MpasiMenu, error)
	FindByID(id uuid.UUID) (*entity.MpasiMenu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *entity.MpasiMenu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) FindAll() ([]entity.MpasiMenu, error) {
	var menus []entity.MpasiMenu

	err := r.db.
		Preload("Ingredients").
		Preload("Steps").
		Preload("Nutritions").
		Find(&menus).Error

	return menus, err
}

func (r *menuRepository) FindByID(id uuid.UUID) (*entity.MpasiMenu, error) {
	var menu entity.MpasiMenu

	err := r.db.
		Preload("Ingredients").
		Preload("Steps").
		Preload("Nutritions").
		Where("id = ?", id).
		First(&menu).Error

	return &menu, err
}
