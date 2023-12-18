package middleware

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	// 3 requests per 10 seconds max
	return limiter.New(limiter.Config{
		Max:          3,
		Expiration:   10 * time.Second,
		LimitReached: LimitReached,
	})
}

func LimitReached(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"success": false,
		"error":   errors.New("Rate limit reached. Please try again later").Error(),
	})
}
