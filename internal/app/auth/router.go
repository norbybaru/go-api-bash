package auth

import (
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/session"

	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(router fiber.Router, db *database.DB, session session.Storage) {
	c := NewAuthController(NewAuthService(NewAuthRepository(db), session))

	group := router.Group("/v1/auth")

	group.Post("/login", c.Login).Name("auth.login")
	group.Post("/register", c.Register).Name("auth.register")
}
