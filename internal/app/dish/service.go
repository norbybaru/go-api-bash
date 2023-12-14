package dish

import "context"

// Service encapsulates business logic for Dishes.
type Service interface {
	// List all dishes and paginate through all
	Browse(ctx context.Context, page int, limit int) error
	// View a single dish
	Read(ctx context.Context, slug string) (*Dish, error)
	// Update a single dish
	Edit(ctx context.Context, input UpdateDishRequest, id string) (*Dish, error)
	// Add a new dish
	Add(ctx context.Context, input CreateDishRequest) (*Dish, error)
	// Delete an existing dish
	Delete(ctx context.Context, id string) error
}

type dishService struct {
	repo *Repository
}

func NewDishService(repo Repository) Service {
	return &dishService{&repo}
}

func (s *dishService) Browse(ctx context.Context, page int, limit int) error {
	return nil
}

func (s *dishService) Read(ctx context.Context, slug string) (*Dish, error) {
	return nil, nil
}

func (s *dishService) Edit(ctx context.Context, input UpdateDishRequest, id string) (*Dish, error) {
	return nil, nil
}

func (s *dishService) Add(ctx context.Context, input CreateDishRequest) (*Dish, error) {
	return nil, nil
}

func (s *dishService) Delete(ctx context.Context, id string) error {
	return nil
}
