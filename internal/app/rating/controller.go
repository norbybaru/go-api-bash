package rating

import (
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/common/response"
	"dancing-pony/internal/platform/validator"

	"github.com/gofiber/fiber/v2"
)

type ratingController struct {
	service Service
}

// Instantiate Rating controller
func NewController(service Service) *ratingController {
	return &ratingController{service}
}

// @Description Add a dish rating by user.
// @Summary create a new dish rating
// @Tags Rating
// @Accept json
// @Produce json
// @Param request body CreateRatingRequest  true "query params"
// @Success 204 {} status "No Content"
// @Success 400 {object} response.ErrorResponse{}
// @Success 401  {object} response.UnauthenticatedResponse{}
// @Success 422 {object} validator.ValidationErrorResponse{}
// @Success 500 {object} response.ErrorResponse{}
// @Security ApiKeyAuth
// @Router /v1/ratings [post]
func (r *ratingController) AddRating(c *fiber.Ctx) error {
	var request CreateRatingRequest

	if err := c.BodyParser(&request); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(response.NewErrorResponse(err))
	}

	validator := validator.NewValidator()

	if err := validator.Validate(request); err != nil {
		return validator.JsonResponse(c, err)
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.NewUnauthenticatedResponse())
	}

	request.UserId = token.GetIdentifier()

	if err := r.service.AddRating(c.Context(), request); err != nil {
		code := fiber.StatusInternalServerError

		if err == validationRatingExist {
			code = fiber.StatusUnprocessableEntity
		}

		return c.
			Status(code).
			JSON(response.NewErrorResponse(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
