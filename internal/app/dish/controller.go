package dish

import (
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
func (r *dishController) ListDishes(c *fiber.Ctx) error {
	dishes, err := r.service.Browse(c.Context(), 0, -1)

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
func (r *dishController) ShowDish(c *fiber.Ctx) error {
	slug := c.Params("slug")
	dish, err := r.service.Read(c.Context(), slug)

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
func (r *dishController) CreateDish(c *fiber.Ctx) error {
	var request CreateDishRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	dish, err := r.service.Add(c.Context(), request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
func (r *dishController) UpdateDish(c *fiber.Ctx) error {
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

	dish, err := r.service.Edit(c.Context(), request, id)

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
func (r *dishController) DestroyDish(c *fiber.Ctx) error {

	dishId := c.Params("id")
	id, err := strconv.Atoi(dishId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := r.service.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})
}
