package auth

import (
	"dancing-pony/internal/common/response"
	"dancing-pony/internal/platform/validator"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	service Service
}

// Instantiate controller
func NewAuthController(service Service) *authController {
	return &authController{service}
}

// Login method to auth user and return access and refresh tokens.
// @Description Authenticate user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest  true "query params"
// @Success 200 {object} response.JsonResponse{data=jwt.Tokens}
// @Success 400 {object} response.ErrorResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Router /v1/auth/login [post]
func (r *authController) Login(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	validator := validator.NewValidator()

	if err := validator.Validate(request); err != nil {
		return validator.JsonResponse(c, err)
	}

	token, err := r.service.Login(c.Context(), request)

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	return c.JSON(response.NewJsonResponse(token))
}

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest  true "query params"
// @Success 200 {object} response.JsonResponse{data=User}
// @Success 400 {object} response.ErrorResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Router /v1/auth/register [post]
func (r *authController) Register(c *fiber.Ctx) error {
	var request RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	validator := validator.NewValidator()

	if err := validator.Validate(request); err != nil {
		return validator.JsonResponse(c, err)
	}

	user, err := r.service.Register(c.Context(), request)

	if err != nil {
		code := fiber.StatusInternalServerError

		if err == validationEmailExist || err == validationNicknameExist {
			code = fiber.StatusUnprocessableEntity
		}

		return c.
			Status(code).
			JSON(response.NewErrorResponse(err))
	}

	return c.JSON(response.NewJsonResponse(user))
}
