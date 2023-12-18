package dish

import (
	"dancing-pony/internal/app/rating"
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(fiber *fiber.App, db *database.DB) {
	c := NewController(
		NewDishService(NewDishRepository(db)),
		rating.NewService(rating.NewRepository(db)),
	)

	router := fiber.Group("/api")

	jwtMiddleware := middleware.JWTMiddleware(config.JWT.Secret, config.JWT.ContextKey)

	group := router.Group("/v1/dishes", jwtMiddleware)
	group.Get("/", middleware.RateLimiter(), c.Browse).Name("dish.index")
	group.Get("/:id", c.Read).Name("dish.show")
	group.Put("/:id", c.Edit).Name("dish.update")
	group.Post("/", c.Add).Name("dish.store")
	group.Delete("/:id", c.Delete).Name("dish.destroy")
}
