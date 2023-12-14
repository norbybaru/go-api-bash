package dish

import "github.com/gofiber/fiber/v2"

type dishController struct {
	service *Service
}

func NewController(service Service) *dishController {
	return &dishController{&service}
}

// Show all dishes resource handler
func (r *dishController) ListDishes(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "list",
	})
}

// View a single dish resource handler
func (r *dishController) ShowDish(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "show",
	})
}

// Create a new dish resource handler
func (r *dishController) CreateDish(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "create",
	})
}

// Update an existing dish resource handler
func (r *dishController) UpdateDish(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "update",
	})
}

// Delete an existing dish resource handler
func (r *dishController) DestroyDish(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "destroy",
	})
}
