package auth

import (
	"dancing-pony/internal/platform/validator"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	service Service
}

func NewAuthController(service Service) *authController {
	return &authController{service}
}

// Authenticate User
func (r *authController) Login(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	validator := validator.NewValidator()

	if err := validator.Validate(request); err != nil {
		return validator.JsonResponse(c, err)
	}

	token, err := r.service.Login(c.Context(), request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    token,
	})
}

// Register new user
func (r *authController) Register(c *fiber.Ctx) error {
	var request RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
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

		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
