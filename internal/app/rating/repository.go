package rating

import (
	"context"
	"dancing-pony/internal/platform/database"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	AddRating(ctx context.Context, rating Rating) error
	FindUserDishRating(ctx context.Context, userId int, dishId int) (*Rating, error)
	FindDishRatingsById(ctx context.Context, dishId int) (*[]Rating, error)
}

type ratingRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &ratingRepository{db}
}

// Add new dish rating
func (r *ratingRepository) AddRating(ctx context.Context, rating Rating) error {
	now := time.Now().UTC()
	rating.CreatedAt = now
	rating.UpdatedAt = now

	return r.db.With(ctx).
		Model(&rating).
		Insert()
}

// Retrieve dish rating by user
func (r *ratingRepository) FindUserDishRating(ctx context.Context, userId int, dishId int) (*Rating, error) {
	var rating Rating
	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{
			"user_id": userId,
			"dish_id": dishId,
		}).
		One(&rating)

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

// Find dish ratings
func (r *ratingRepository) FindDishRatingsById(ctx context.Context, dishId int) (*[]Rating, error) {
	var ratings *[]Rating

	err := r.db.With(ctx).
		Select().
		From("ratings").
		Where(dbx.HashExp{
			"dish_id": dishId,
		}).
		All(&ratings)

	if err != nil {
		return nil, err
	}

	return ratings, nil
}
