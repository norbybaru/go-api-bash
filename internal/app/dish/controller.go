package dish

import (
	"dancing-pony/internal/platform/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type dishController struct {
	service Service
}

func NewController(service Service) *dishController {
	return &dishController{service}
}

// Show all dishes resource handler
func (r *dishController) Browse(c *fiber.Ctx) error {
	dishes, err := r.service.ListDishes(c.Context(), 0, -1)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    dishes,
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

	return c.JSON(fiber.Map{
		"success": true,
		"data":    dish,
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

	if err := r.service.DeleteDish(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})
}
