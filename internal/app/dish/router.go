package dish

import (
	"dancing-pony/internal/platform/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(fiber *fiber.App, db *database.DB) {
	c := NewController(NewDishService(NewDishRepository(db)))

	router := fiber.Group("/api")

	group := router.Group("/v1/dishes")
	group.Get("/", c.Browse).Name("dish.index")
	group.Get("/:id", c.Read).Name("dish.show")
	group.Put("/:id", c.Edit).Name("dish.update")
	group.Post("/", c.Add).Name("dish.store")
	group.Delete("/:id", c.Delete).Name("dish.destroy")
}
