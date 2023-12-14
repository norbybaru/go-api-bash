package dish

import (
	"dancing-pony/internal/platform/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(fiber *fiber.App, db *database.DB) {
	c := NewController(NewDishService(NewDishRepository(db)))

	router := fiber.Group("/api")

	group := router.Group("/v1/dishes")
	group.Get("/", c.ListDishes).Name("dish.index")
	group.Post("/", c.CreateDish).Name("dish.store")
	group.Put("/:id", c.UpdateDish).Name("dish.update")
	group.Delete("/:id", c.DestroyDish).Name("dish.destroy")
}
