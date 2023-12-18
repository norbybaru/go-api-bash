package rating

import (
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(router fiber.Router, db *database.DB) {
	c := NewController(NewService(NewRepository(db)))

	jwtMiddleware := middleware.JWTMiddleware(config.JWT.Secret, config.JWT.ContextKey)

	group := router.Group("/v1/ratings", jwtMiddleware)
	group.Post("/", c.AddRating).Name("rating.store")
}
