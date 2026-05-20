package usecase

import (
	"errors"

	"sigap-backend/app/repository"
	"sigap-backend/domain/dto"
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type ScheduleUsecase interface {
	Create(userID uuid.UUID, req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error)
	FindAll() ([]dto.ScheduleResponse, error)
	FindAllByChildID(userID uuid.UUID, childID uuid.UUID) ([]dto.ScheduleResponse, error)
	FindByID(userID uuid.UUID, scheduleID uuid.UUID) (*dto.ScheduleResponse, error)
	Update(userID uuid.UUID, scheduleID uuid.UUID, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error)
	Delete(userID uuid.UUID, scheduleID uuid.UUID) error
}

type scheduleUsecase struct {
	scheduleRepo repository.ScheduleRepository
	childRepo    repository.ChildRepository
}

func NewScheduleUsecase(
	scheduleRepo repository.ScheduleRepository,
	childRepo repository.ChildRepository,
) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepo: scheduleRepo,
		childRepo:    childRepo,
	}
}

func (u *scheduleUsecase) Create(userID uuid.UUID, req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error) {
	child, err := u.childRepo.FindByID(req.ChildID)
	if err != nil {
		return nil, errors.New("child profile not found")
	}

	if child.UserID != userID {
		return nil, errors.New("unauthorized access to child profile")
	}

	schedule := entity.Schedule{
		ChildID:      req.ChildID,
		Title:        req.Title,
		Type:         req.Type,
		Status:       entity.Pending,
		ScheduleDate: req.ScheduleDate,
		Location:     req.Location,
	}

	if err := u.scheduleRepo.Create(&schedule); err != nil {
		return nil, errors.New("failed to create schedule")
	}

	schedule.Child = *child

	return mapScheduleToResponse(schedule), nil
}

func (u *scheduleUsecase) FindAll() ([]dto.ScheduleResponse, error) {
	schedules, err := u.scheduleRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch schedules")
	}

	var responses []dto.ScheduleResponse
	for _, schedule := range schedules {
		responses = append(responses, *mapScheduleToResponse(schedule))
	}

	return responses, nil
}

func (u *scheduleUsecase) FindAllByChildID(userID uuid.UUID, childID uuid.UUID) ([]dto.ScheduleResponse, error) {
	child, err := u.childRepo.FindByID(childID)
	if err != nil {
		return nil, errors.New("child profile not found")
	}

	if child.UserID != userID {
		return nil, errors.New("unauthorized access to child profile")
	}

	schedules, err := u.scheduleRepo.FindAllByChildID(childID)
	if err != nil {
		return nil, errors.New("failed to fetch schedules")
	}

	var responses []dto.ScheduleResponse
	for _, schedule := range schedules {
		responses = append(responses, *mapScheduleToResponse(schedule))
	}

	return responses, nil
}

func (u *scheduleUsecase) FindByID(userID uuid.UUID, scheduleID uuid.UUID) (*dto.ScheduleResponse, error) {
	schedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}

	if schedule.Child.UserID != userID {
		return nil, errors.New("unauthorized access to schedule")
	}

	return mapScheduleToResponse(*schedule), nil
}

func (u *scheduleUsecase) Update(userID uuid.UUID, scheduleID uuid.UUID, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error) {
	schedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}

	if schedule.Child.UserID != userID {
		return nil, errors.New("unauthorized access to schedule")
	}

	child, err := u.childRepo.FindByID(req.ChildID)
	if err != nil {
		return nil, errors.New("child profile not found")
	}

	if child.UserID != userID {
		return nil, errors.New("unauthorized access to child profile")
	}

	schedule.ChildID = req.ChildID
	schedule.Title = req.Title
	schedule.Type = req.Type
	schedule.Status = req.Status
	schedule.ScheduleDate = req.ScheduleDate
	schedule.Location = req.Location

	if err := u.scheduleRepo.Update(schedule); err != nil {
		return nil, errors.New("failed to update schedule")
	}

	schedule.Child = *child

	return mapScheduleToResponse(*schedule), nil
}

func (u *scheduleUsecase) Delete(userID uuid.UUID, scheduleID uuid.UUID) error {
	schedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	if schedule.Child.UserID != userID {
		return errors.New("unauthorized access to schedule")
	}

	if err := u.scheduleRepo.Delete(scheduleID); err != nil {
		return errors.New("failed to delete schedule")
	}

	return nil
}

func mapScheduleToResponse(schedule entity.Schedule) *dto.ScheduleResponse {
	childName := ""

	if schedule.Child.ID != uuid.Nil {
		childName = schedule.Child.Name
	}

	return &dto.ScheduleResponse{
		ID:           schedule.ID,
		ChildID:      schedule.ChildID,
		ChildName:    childName,
		Title:        schedule.Title,
		Type:         schedule.Type,
		Status:       schedule.Status,
		ScheduleDate: schedule.ScheduleDate,
		Location:     schedule.Location,
	}
}
