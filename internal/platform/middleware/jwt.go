package middleware

import (
	"dancing-pony/internal/common/response"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(secret string, contextKey string) fiber.Handler {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(secret)},
		ContextKey:   contextKey, // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "missing or malformed JWT" {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.NewUnauthenticatedResponse())
	}

	// Return status 401 and failed authentication error.
	return c.
		Status(fiber.StatusUnauthorized).
		JSON(response.NewUnauthenticatedResponse())
}
