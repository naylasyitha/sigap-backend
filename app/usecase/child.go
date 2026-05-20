package usecase

import (
	"errors"

	"sigap-backend/app/repository"
	"sigap-backend/domain/dto"
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type ChildUsecase interface {
	Create(userID uuid.UUID, req dto.CreateChildRequest) (*dto.ChildResponse, error)
	FindAllByUserID(userID uuid.UUID) ([]dto.ChildResponse, error)
	Update(userID uuid.UUID, childID uuid.UUID, req dto.UpdateChildRequest) (*dto.ChildResponse, error)
	Delete(userID uuid.UUID, childID uuid.UUID) error
}

type childUsecase struct {
	childRepo repository.ChildRepository
}

func NewChildUsecase(childRepo repository.ChildRepository) ChildUsecase {
	return &childUsecase{childRepo: childRepo}
}

func (u *childUsecase) Create(userID uuid.UUID, req dto.CreateChildRequest) (*dto.ChildResponse, error) {
	child := entity.Child{
		UserID:      userID,
		Name:        req.Name,
		Gender:      req.Gender,
		BirthDate:   req.BirthDate,
		BirthWeight: req.BirthWeight,
		BirthHeight: req.BirthHeight,
	}

	if err := u.childRepo.Create(&child); err != nil {
		return nil, errors.New("failed to create child profile")
	}

	return mapChildToResponse(child), nil
}

func (u *childUsecase) FindAllByUserID(userID uuid.UUID) ([]dto.ChildResponse, error) {
	children, err := u.childRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to fetch child profiles")
	}

	var responses []dto.ChildResponse
	for _, child := range children {
		responses = append(responses, *mapChildToResponse(child))
	}

	return responses, nil
}

func (u *childUsecase) Update(userID uuid.UUID, childID uuid.UUID, req dto.UpdateChildRequest) (*dto.ChildResponse, error) {
	child, err := u.childRepo.FindByID(childID)
	if err != nil {
		return nil, errors.New("child profile not found")
	}

	if child.UserID != userID {
		return nil, errors.New("unauthorized access to child profile")
	}

	child.Name = req.Name
	child.Gender = req.Gender
	child.BirthDate = req.BirthDate
	child.BirthWeight = req.BirthWeight
	child.BirthHeight = req.BirthHeight

	if err := u.childRepo.Update(child); err != nil {
		return nil, errors.New("failed to update child profile")
	}

	return mapChildToResponse(*child), nil
}

func (u *childUsecase) Delete(userID uuid.UUID, childID uuid.UUID) error {
	child, err := u.childRepo.FindByID(childID)
	if err != nil {
		return errors.New("child profile not found")
	}

	if child.UserID != userID {
		return errors.New("unauthorized access to child profile")
	}

	if err := u.childRepo.Delete(childID); err != nil {
		return errors.New("failed to delete child profile")
	}

	return nil
}

func mapChildToResponse(child entity.Child) *dto.ChildResponse {
	return &dto.ChildResponse{
		ID:          child.ID,
		UserID:      child.UserID,
		Name:        child.Name,
		Gender:      child.Gender,
		BirthDate:   child.BirthDate,
		BirthWeight: child.BirthWeight,
		BirthHeight: child.BirthHeight,
	}
}
