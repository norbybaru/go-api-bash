package dish

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

// Service encapsulates business logic for Dishes.
type Service interface {
	// List all dishes and paginate through all
	ListDishes(ctx context.Context, page int, limit int) (*[]Dish, error)
	// View a single dish
	ViewDish(ctx context.Context, id int) (*Dish, error)
	// Update a single dish
	UpdateDish(ctx context.Context, input UpdateDishRequest, id int) (*Dish, error)
	// Add a new dish
	CreateDish(ctx context.Context, input CreateDishRequest) (*Dish, error)
	// Delete an existing dish
	DeleteDish(ctx context.Context, id int) error
}

type dishService struct {
	repo Repository
}

func NewDishService(repo Repository) Service {
	return &dishService{repo}
}

var (
	errorCreateDish            = errors.New("Failed to create a new dish")
	errorBrowseDishes          = errors.New("Failed to retrieve dishes")
	errorInvalidDish           = errors.New("Failed to retrieve dish")
	errorUpdateDish            = errors.New("Failed to update dish")
	ValidationNameAlreadyExist = errors.New("Dish name is already taken")
	ErrorResourceNotFound      = errors.New("Resource not found")
)

func (s *dishService) ListDishes(ctx context.Context, page int, limit int) (*[]Dish, error) {
	dishes, err := s.repo.Query(ctx, page, limit)

	if err != nil {
		log.Error(err)
		return nil, errorBrowseDishes
	}

	return dishes, nil
}

func (s *dishService) ViewDish(ctx context.Context, id int) (*Dish, error) {
	dish, err := s.repo.GetById(ctx, id)

	if err != nil {
		log.Error(err)
		return nil, errorInvalidDish
	}

	return dish, nil
}

func (s *dishService) UpdateDish(ctx context.Context, input UpdateDishRequest, id int) (*Dish, error) {
	dish, err := s.repo.GetById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorResourceNotFound
		}

		log.Error(err)

		return nil, errorInvalidDish
	}

	updatedDish := NewDish(input.Name, input.Description, input.ImageUrl, input.Price)

	existingDish, err := s.repo.GetBySlug(ctx, dish.Slug)

	if err != nil {
		log.Error(err)
		return nil, errorCreateDish
	}

	if existingDish != nil && existingDish.Slug != updatedDish.Slug {
		return nil, ValidationNameAlreadyExist
	}

	updatedDish.Id = dish.Id
	updatedDish.CreatedAt = dish.CreatedAt

	if err := s.repo.Update(ctx, *updatedDish); err != nil {
		log.Error(err)
		return nil, errorUpdateDish
	}

	return updatedDish, nil
}

func (s *dishService) CreateDish(ctx context.Context, input CreateDishRequest) (*Dish, error) {

	dish := NewDish(input.Name, input.Description, input.ImageUrl, input.Price)

	exist, err := s.repo.DishSlugExist(ctx, dish.Slug)

	if err != nil {
		log.Error(err)
		return nil, errorCreateDish
	}

	if exist {
		return nil, ValidationNameAlreadyExist
	}

	if err := s.repo.Create(ctx, *dish); err != nil {
		log.Error(err)
		return nil, errorCreateDish
	}

	return dish, nil
}

func (s *dishService) DeleteDish(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if err == sql.ErrNoRows {
			return ErrorResourceNotFound
		}

		log.Error(err)

		return err
	}

	return nil
}
