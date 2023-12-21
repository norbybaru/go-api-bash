package rating

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

type Service interface {
	FindDishRating(ctx context.Context, dishId int) (*[]Rating, error)
	AddRating(ctx context.Context, input CreateRatingRequest) error
}

type ratingService struct {
	repo Repository
}

var (
	errorFailedProcessRequest = errors.New("Failed to process request")
	errorResourceNotFound     = errors.New("Resource not found")
	validationRatingExist     = errors.New("Rating already exist")
)

func NewService(repo Repository) Service {
	return &ratingService{repo}
}

// Retrieve ratings of a dish
func (s *ratingService) FindDishRating(ctx context.Context, dishId int) (*[]Rating, error) {
	ratings, err := s.repo.FindDishRatingsById(ctx, dishId)

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		return nil, errorFailedProcessRequest
	}

	return ratings, nil
}

// Add new dish rating
func (s *ratingService) AddRating(ctx context.Context, input CreateRatingRequest) error {
	rating, err := s.repo.FindUserDishRating(ctx, input.UserId, input.DishId)

	if err != nil {

		if err == sql.ErrNoRows {
			return errorResourceNotFound
		}

		log.Error(err)
		return err
	}

	if rating != nil {
		return validationRatingExist
	}

	newRating := NewRating(input.UserId, input.DishId, input.Rate)

	if err := s.repo.AddRating(ctx, *newRating); err != nil {
		log.Error(err)
		return errorFailedProcessRequest
	}

	return nil
}
