package usecase

import (
	"errors"
	"sigap-backend/app/repository"
	"sigap-backend/domain/dto"
	"sigap-backend/domain/entity"

	"github.com/google/uuid"
)

type MenuUsecase interface {
	Create(req dto.CreateMenuRequest) (*dto.MenuResponse, error)
	FindAll() ([]dto.MenuResponse, error)
	FindByID(id uuid.UUID) (*dto.MenuResponse, error)
}

type menuUsecase struct {
	menuRepo repository.MenuRepository
}

func NewMenuUsecase(menuRepo repository.MenuRepository) MenuUsecase {
	return &menuUsecase{menuRepo: menuRepo}
}

func (u *menuUsecase) Create(req dto.CreateMenuRequest) (*dto.MenuResponse, error) {
	menu := entity.MpasiMenu{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		AgeRange:   req.AgeRange,
		MealType:   req.MealType,
		Difficulty: req.Difficulty,
		Duration:   req.Duration,
		Portion:    req.Portion,
		Calories:   req.Calories,
	}

	for _, ingredient := range req.Ingredients {
		menu.Ingredients = append(menu.Ingredients, entity.MpasiIngredient{
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
		})
	}

	for _, step := range req.Steps {
		menu.Steps = append(menu.Steps, entity.MpasiStep{
			StepNumber:  step.StepNumber,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, nutrition := range req.Nutritions {
		menu.Nutritions = append(menu.Nutritions, entity.MpasiNutrition{
			Name:  nutrition.Name,
			Value: nutrition.Value,
		})
	}

	if err := u.menuRepo.Create(&menu); err != nil {
		return nil, errors.New("failed to create menu")
	}

	return mapMenuToResponse(menu), nil
}

func (u *menuUsecase) FindAll() ([]dto.MenuResponse, error) {
	menus, err := u.menuRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch menus")
	}

	var responses []dto.MenuResponse
	for _, menu := range menus {
		responses = append(responses, *mapMenuToResponse(menu))
	}

	return responses, nil
}

func (u *menuUsecase) FindByID(id uuid.UUID) (*dto.MenuResponse, error) {
	menu, err := u.menuRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	return mapMenuToResponse(*menu), nil
}

func mapMenuToResponse(menu entity.MpasiMenu) *dto.MenuResponse {
	res := &dto.MenuResponse{
		ID:         menu.ID,
		Name:       menu.Name,
		ImageURL:   menu.ImageURL,
		AgeRange:   menu.AgeRange,
		MealType:   menu.MealType,
		Difficulty: menu.Difficulty,
		Duration:   menu.Duration,
		Portion:    menu.Portion,
		Calories:   menu.Calories,
	}

	for _, ingredient := range menu.Ingredients {
		res.Ingredients = append(res.Ingredients, dto.IngredientResponse{
			ID:     ingredient.ID,
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
		})
	}

	for _, step := range menu.Steps {
		res.Steps = append(res.Steps, dto.StepResponse{
			ID:          step.ID,
			StepNumber:  step.StepNumber,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, nutrition := range menu.Nutritions {
		res.Nutritions = append(res.Nutritions, dto.NutritionResponse{
			ID:    nutrition.ID,
			Name:  nutrition.Name,
			Value: nutrition.Value,
		})
	}

	return res
}
