package rating

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddRating(ctx context.Context, rating Rating) error {
	args := m.Called(rating)
	return args.Error(0)
}

func (m *MockRepository) FindUserDishRating(ctx context.Context, userId int, dishId int) (*Rating, error) {
	args := m.Called(userId, dishId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Rating), args.Error(1)
}

func (m *MockRepository) FindDishRatingsById(ctx context.Context, dishId int) (*[]Rating, error) {
	args := m.Called(dishId)

	return args.Get(0).(*[]Rating), args.Error(1)
}

var ctx = context.Background()

func missingDish(t *testing.T) {
	mockObj := new(MockRepository)
	service := NewService(mockObj)

	input := CreateRatingRequest{
		Rate:   5,
		DishId: 1,
		UserId: 2,
	}

	// set up expectations with a placeholder in the argument list
	mockObj.On("FindUserDishRating", input.UserId, input.DishId).Return(nil, sql.ErrNoRows)

	// call the code we are testing
	err := service.AddRating(ctx, input)

	// assert that the expectations were met
	mockObj.AssertExpectations(t)

	if err != errorResourceNotFound {
		t.Errorf("Expected err: %v, but got: %v", errorResourceNotFound, err)
	}
}

func addNewDishRating(t *testing.T) {
	mockObj := new(MockRepository)
	service := NewService(mockObj)

	input := CreateRatingRequest{
		Rate:   5,
		DishId: 1,
		UserId: 2,
	}

	// set up expectations with a placeholder in the argument list
	mockObj.On("FindUserDishRating", input.UserId, input.DishId).Return(nil, nil)

	newRating := NewRating(input.UserId, input.DishId, input.Rate)
	// set up expectations with a placeholder in the argument list
	mockObj.On("AddRating", *newRating).Return(nil)

	// call the code we are testing
	err := service.AddRating(ctx, input)

	// assert that the expectations were met
	mockObj.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func alreadyRatedDish(t *testing.T) {
	mockObj := new(MockRepository)
	service := NewService(mockObj)

	input := CreateRatingRequest{
		Rate:   5,
		DishId: 1,
		UserId: 2,
	}

	rating := NewRating(input.UserId, input.DishId, input.Rate)

	// set up expectations with a placeholder in the argument list
	mockObj.On("FindUserDishRating", input.UserId, input.DishId).Return(rating, nil)
	// call the code we are testing
	err := service.AddRating(ctx, input)

	// assert that the expectations were met
	mockObj.AssertExpectations(t)

	if err != validationRatingExist {
		t.Errorf("Expected: %v, Got: %v", validationRatingExist, err)
	}
}

func TestAddRating(t *testing.T) {
	t.Run("Add new dish rating", addNewDishRating)

	t.Run("Dish Already rated", alreadyRatedDish)

	t.Run("Cannot rate invalid dish", missingDish)
}
