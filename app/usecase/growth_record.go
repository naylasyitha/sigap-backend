package usecase

import (
	"errors"

	"sigap-backend/app/repository"
	"sigap-backend/domain/dto"
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type GrowthRecordUsecase interface {
	Create(req dto.CreateGrowthRecordRequest) (*dto.GrowthRecordResponse, error)
	FindAllByChildID(childID uuid.UUID) ([]dto.GrowthRecordResponse, error)
	GetLatestByChildID(childID uuid.UUID) (*dto.GrowthRecordResponse, error)
}

type growthRecordUsecase struct {
	growthRepo repository.GrowthRecordRepository
	childRepo  repository.ChildRepository
}

func NewGrowthRecordUsecase(
	growthRepo repository.GrowthRecordRepository,
	childRepo repository.ChildRepository,
) GrowthRecordUsecase {
	return &growthRecordUsecase{
		growthRepo: growthRepo,
		childRepo:  childRepo,
	}
}

func (u *growthRecordUsecase) Create(req dto.CreateGrowthRecordRequest) (*dto.GrowthRecordResponse, error) {
	_, err := u.childRepo.FindByID(req.ChildID)
	if err != nil {
		return nil, errors.New("child not found")
	}

	record := entity.GrowthRecord{
		ChildID:           req.ChildID,
		Weight:            req.Weight,
		Height:            req.Height,
		HeadCircumference: req.HeadCircumference,
		MeasurementDate:   req.MeasurementDate,
	}

	if err := u.growthRepo.Create(&record); err != nil {
		return nil, errors.New("failed to create growth record")
	}

	return mapGrowthRecordToResponse(record), nil
}

func (u *growthRecordUsecase) FindAllByChildID(childID uuid.UUID) ([]dto.GrowthRecordResponse, error) {
	records, err := u.growthRepo.FindAllByChildID(childID)
	if err != nil {
		return nil, errors.New("failed to fetch growth records")
	}

	var responses []dto.GrowthRecordResponse
	for _, record := range records {
		responses = append(responses, *mapGrowthRecordToResponse(record))
	}

	return responses, nil
}

func (u *growthRecordUsecase) GetLatestByChildID(childID uuid.UUID) (*dto.GrowthRecordResponse, error) {
	record, err := u.growthRepo.GetLatestByChildID(childID)
	if err != nil {
		return nil, errors.New("growth record not found")
	}

	return mapGrowthRecordToResponse(*record), nil
}

func mapGrowthRecordToResponse(record entity.GrowthRecord) *dto.GrowthRecordResponse {
	return &dto.GrowthRecordResponse{
		ID:                record.ID,
		ChildID:           record.ChildID,
		Weight:            record.Weight,
		Height:            record.Height,
		HeadCircumference: record.HeadCircumference,
		MeasurementDate:   record.MeasurementDate,
	}
}
