package auth

import (
	"context"
	"dancing-pony/internal/common/bcrypt"
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/session"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
)

type Service interface {
	// Register new user
	Register(ctx context.Context, input RegisterRequest) (*User, error)
	// Authenticate user
	Login(ctx context.Context, input LoginRequest) (*jwt.Tokens, error)
}

type authService struct {
	repo    Repository
	session session.Storage
}

func NewAuthService(repo Repository, session session.Storage) Service {
	return &authService{repo, session}
}

var (
	errorRequestFailed      = errors.New("Failed to process request")
	errorRegistrationFailed = errors.New("Failed to register user")
	validationEmailExist    = errors.New("Email already taken by another user")
	validationNicknameExist = errors.New("Nickname already taken by another user")
	errorInvalidCredentials = errors.New("Invalid credentials")
)

func (s *authService) validateUniqueFields(ctx context.Context, input RegisterRequest) error {
	exist, err := s.repo.CheckEmailExist(ctx, input.Email)

	if err != nil {
		log.Error(err)
		return errorRequestFailed
	}

	if exist {
		return validationEmailExist
	}

	exist, err = s.repo.CheckNicknameExist(ctx, input.Nickname)

	if err != nil {
		log.Error(err)
		return errorRequestFailed
	}

	if exist {
		return validationNicknameExist
	}

	return nil
}

func (s *authService) Register(ctx context.Context, input RegisterRequest) (*User, error) {
	err := s.validateUniqueFields(ctx, input)

	if err != nil {
		return nil, err
	}

	password, err := bcrypt.HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	user := NewUser(input.Name, input.Email, password, input.Nickname)

	if err := s.repo.Create(ctx, *user); err != nil {
		log.Error(err)
		return nil, errorRegistrationFailed
	}

	user, err = s.repo.FindByEmail(ctx, input.Email)

	if err != nil {
		log.Error(err)
		return nil, errorRegistrationFailed
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, input LoginRequest) (*jwt.Tokens, error) {
	user, err := s.repo.FindByEmail(ctx, input.Email)

	if err != nil {
		log.Error(err)
		return nil, errorInvalidCredentials
	}

	if !bcrypt.ComparePasswords(user.Password, input.Password) {
		return nil, errorInvalidCredentials
	}

	jwt := jwt.Init(config.JWT.Secret, config.JWT.ExpireMinutes)

	fmt.Printf("%+v", user)
	token, err := jwt.GenerateNewTokens(user.Id)

	if err != nil {
		log.Error(err)
		return nil, errorRequestFailed
	}

	key := fmt.Sprintf("%s_%v", config.JWT.ContextKey, user.Id)

	if err := s.session.Set(key, token.Access, 0); err != nil {
		log.Error(err)
		return nil, errorRequestFailed
	}

	return token, nil
}
