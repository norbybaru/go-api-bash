package auth

import (
	"context"
	"dancing-pony/internal/platform/database"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	// Returns a User with the specified ID.
	FindById(ctx context.Context, id string) (*User, error)
	// Return a User with the specified email
	FindByEmail(ctx context.Context, email string) (*User, error)
	// Verify that email exists
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	// Verify that email exists
	CheckNicknameExist(ctx context.Context, email string) (bool, error)
	// Create a new User in the DB.
	Create(ctx context.Context, user User) error
}

type authRepository struct {
	db *database.DB
}

func NewAuthRepository(db *database.DB) Repository {
	return &authRepository{db}
}

func (r *authRepository) FindById(ctx context.Context, id string) (*User, error) {
	var user User

	err := r.db.
		With(ctx).
		Select().
		Model(id, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	err := r.db.
		With(ctx).
		Select().
		Where(dbx.HashExp{"email": email}).
		One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	var count int
	err := r.db.With(ctx).
		Select("COUNT(*)").
		From("users").
		Where(dbx.HashExp{
			"email": email,
		}).
		Row(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *authRepository) CheckNicknameExist(ctx context.Context, nickname string) (bool, error) {
	var count int
	err := r.db.With(ctx).
		Select("COUNT(*)").
		From("users").
		Where(dbx.HashExp{
			"nickname": nickname,
		}).
		Row(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *authRepository) Create(ctx context.Context, user User) error {
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.With(ctx).
		Model(&user).
		Insert()
}
