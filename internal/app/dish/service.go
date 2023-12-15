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
	Browse(ctx context.Context, page int, limit int) (*[]Dish, error)
	// View a single dish
	Read(ctx context.Context, slug string) (*Dish, error)
	// Update a single dish
	Edit(ctx context.Context, input UpdateDishRequest, id int) (*Dish, error)
	// Add a new dish
	Add(ctx context.Context, input CreateDishRequest) (*Dish, error)
	// Delete an existing dish
	Delete(ctx context.Context, id int) error
}

type dishService struct {
	repo Repository
}

func NewDishService(repo Repository) Service {
	return &dishService{repo}
}

var (
	errorCreateDish       = errors.New("Failed to create a new dish")
	errorBrowseDishes     = errors.New("Failed to retrieve dishes")
	errorInvalidDish      = errors.New("Failed to retrieve dish")
	errorUpdateDish       = errors.New("Failed to update dish")
	ErrorResourceNotFound = errors.New("Resource not found")
)

func (s *dishService) Browse(ctx context.Context, page int, limit int) (*[]Dish, error) {
	dishes, err := s.repo.Query(ctx, page, limit)

	if err != nil {
		log.Error(err)
		return nil, errorBrowseDishes
	}

	return dishes, nil
}

func (s *dishService) Read(ctx context.Context, slug string) (*Dish, error) {
	dish, err := s.repo.GetBySlug(ctx, slug)

	if err != nil {
		log.Error(err)
		return nil, errorInvalidDish
	}

	return dish, nil
}

func (s *dishService) Edit(ctx context.Context, input UpdateDishRequest, id int) (*Dish, error) {
	dish, err := s.repo.GetById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorResourceNotFound
		}

		log.Error(err)

		return nil, errorInvalidDish
	}

	dishData := NewDish(input.Name, input.Description, input.ImageUrl, input.Price)
	dishData.Id = dish.Id
	dishData.CreatedAt = dish.CreatedAt

	if err := s.repo.Update(ctx, *dishData); err != nil {
		log.Error(err)
		return nil, errorUpdateDish
	}

	return dishData, nil
}

func (s *dishService) Add(ctx context.Context, input CreateDishRequest) (*Dish, error) {

	dish := NewDish(input.Name, input.Description, input.ImageUrl, input.Price)

	if err := s.repo.Create(ctx, *dish); err != nil {
		log.Error(err)
		return nil, errorCreateDish
	}

	return dish, nil
}

func (s *dishService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if err == sql.ErrNoRows {
			return ErrorResourceNotFound
		}

		log.Error(err)

		return err
	}

	return nil
}
