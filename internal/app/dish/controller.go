package dish

import (
	"dancing-pony/internal/app/rating"
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/common/utils"
	"dancing-pony/internal/platform/paginator"
	"dancing-pony/internal/platform/validator"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type dishController struct {
	service       Service
	ratingService rating.Service
}

func NewController(service Service, ratingService rating.Service) *dishController {
	return &dishController{service, ratingService}
}

// Show all dishes resource handler
func (r *dishController) Browse(c *fiber.Ctx) error {
	limit := utils.ParseInt(c.Query(paginator.PerPageVar), paginator.DefaultPageSize)
	page := utils.ParseInt(c.Query(paginator.PageVar), 0)

	pagination, err := r.service.ListDishes(c.Context(), page, limit)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL(), c.Path())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    pagination.Records,
		"meta":    pagination.Paginator,
		"links":   pagination.Paginator.BuildLinks(fullURL, limit),
	})
}

// View a single dish resource handler
func (r *dishController) Read(c *fiber.Ctx) error {
	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	dish, err := r.service.ViewDish(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	ratings, err := r.ratingService.FindDishRating(c.Context(), dish.Id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	var ratingSlice = []RatingResource{}
	for _, v := range *ratings {

		data := RatingResource{
			UserId: v.UserId,
			Rating: int(v.Rate),
		}

		ratingSlice = append(ratingSlice, data)
	}

	dishResource := DishResourceResponse{Dish: *dish, Ratings: ratingSlice}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    dishResource,
	})
}

// Create a new dish resource handler
func (r *dishController) Add(c *fiber.Ctx) error {
	var request CreateDishRequest

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

	request.UserId = int(token.Identifier.(float64))

	dish, err := r.service.CreateDish(c.Context(), request)

	if err != nil {
		code := fiber.StatusInternalServerError

		if err == ValidationNameAlreadyExist {
			code = fiber.StatusBadRequest
		}

		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    &dish,
	})
}

// Update an existing dish resource handler
func (r *dishController) Edit(c *fiber.Ctx) error {
	var request UpdateDishRequest

	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   jwt.ErrorUnAuthenticated,
		})
	}

	request.UserId = int(token.Identifier.(float64))

	dish, err := r.service.UpdateDish(c.Context(), request, id)

	if err != nil {
		code := fiber.StatusInternalServerError

		if err == ErrorResourceNotFound {
			code = fiber.StatusNotFound
		}

		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    dish,
	})
}

// Delete an existing dish resource handler
func (r *dishController) Delete(c *fiber.Ctx) error {
	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   jwt.ErrorUnAuthenticated,
		})
	}

	userId := int(token.Identifier.(float64))

	if err := r.service.DeleteDish(c.Context(), id, userId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})
}
