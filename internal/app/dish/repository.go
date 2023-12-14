package dish

import (
	"context"
	"dancing-pony/internal/platform/database"
)

// Repository encapsulates the logic to access dishes from the data source.
type Repository interface {
	// Get returns the dish with the specified dish ID.
	Get(ctx context.Context, id string) (*Dish, error)
	// Count returns the number of dishes.
	Count(ctx context.Context) (int, error)
	// Query returns the list of dishes with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]*Dish, error)
	// Create saves a new dish in the storage.
	Create(ctx context.Context, dish Dish) error
	// Update updates the dish with given ID in the storage.
	Update(ctx context.Context, dish Dish) error
	// Delete removes the dish with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

type dishRepository struct {
	db *database.DB
}

func NewDishRepository(db *database.DB) Repository {
	return &dishRepository{db}
}

func (r *dishRepository) Get(ctx context.Context, id string) (*Dish, error) {
	return nil, nil
}

func (r *dishRepository) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func (r *dishRepository) Query(ctx context.Context, offset, limit int) ([]*Dish, error) {
	return nil, nil
}

func (r *dishRepository) Create(ctx context.Context, dish Dish) error {
	return nil
}

func (r *dishRepository) Update(ctx context.Context, dish Dish) error {
	return nil
}

func (r *dishRepository) Delete(ctx context.Context, id string) error {
	return nil
}
