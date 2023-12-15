package dish

import (
	"context"
	"dancing-pony/internal/platform/database"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access dishes from the data source.
type Repository interface {
	// Get returns the dish with the specified dish ID.
	GetById(ctx context.Context, id int) (*Dish, error)
	// Get returns the dish with the specified dish ID.
	GetBySlug(ctx context.Context, slug string) (*Dish, error)
	// Count returns the number of dishes.
	Count(ctx context.Context) (int, error)
	// Query returns the list of dishes with the given offset and limit.
	Query(ctx context.Context, offset, limit int) (*[]Dish, error)
	// Create saves a new dish in the storage.
	Create(ctx context.Context, dish Dish) error
	// Update updates the dish with given ID in the storage.
	Update(ctx context.Context, dish Dish) error
	// Delete removes the dish with given ID from the storage.
	Delete(ctx context.Context, id int) error
}

type dishRepository struct {
	db *database.DB
}

func NewDishRepository(db *database.DB) Repository {
	return &dishRepository{db}
}

func (r *dishRepository) GetById(ctx context.Context, id int) (*Dish, error) {
	var dish Dish
	err := r.db.With(ctx).
		Select().
		Model(id, &dish)

	return &dish, err
}

func (r *dishRepository) GetBySlug(ctx context.Context, slug string) (*Dish, error) {
	var dish Dish
	err := r.db.
		With(ctx).
		Select().
		Where(dbx.HashExp{"slug": slug}).
		One(&dish)

	return &dish, err
}

func (r *dishRepository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).
		Select("COUNT(*)").
		From("dishes").
		Row(&count)

	return count, err
}

func (r *dishRepository) Query(ctx context.Context, offset, limit int) (*[]Dish, error) {
	var dishes *[]Dish
	err := r.db.With(ctx).
		Select().
		From("dishes").
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&dishes)

	return dishes, err
}

func (r *dishRepository) Create(ctx context.Context, dish Dish) error {
	now := time.Now().UTC()
	dish.CreatedAt = now
	dish.UpdatedAt = now

	return r.db.With(ctx).
		Model(&dish).
		Insert()
}

func (r *dishRepository) Update(ctx context.Context, dish Dish) error {
	dish.UpdatedAt = time.Now().UTC()

	return r.db.With(ctx).
		Model(&dish).
		Update()
}

func (r *dishRepository) Delete(ctx context.Context, id int) error {
	dish, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}

	return r.db.With(ctx).
		Model(dish).
		Delete()
}
