package rating

import (
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/platform/validator"

	"github.com/gofiber/fiber/v2"
)

type ratingController struct {
	service Service
}

func NewController(service Service) *ratingController {
	return &ratingController{service}
}

func (r *ratingController) AddRating(c *fiber.Ctx) error {
	var request CreateRatingRequest

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

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   jwt.ErrorUnAuthenticated,
		})
	}

	request.UserId = token.GetIdentifier()

	if err := r.service.AddRating(c.Context(), request); err != nil {
		code := fiber.StatusInternalServerError

		if err == validationRatingExist {
			code = fiber.StatusUnprocessableEntity
		}

		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
