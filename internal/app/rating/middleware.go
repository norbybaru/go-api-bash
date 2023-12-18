package rating

import (
	"dancing-pony/internal/common/jwt"
	"errors"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func RestrictUser(c *fiber.Ctx) error {
	token, err := jwt.ExtractTokenMetadata(c)

	if err != nil {
		return err
	}

	banned := []string{"Sméagol"}

	if slices.Contains(banned, token.NickName) || slices.Contains(banned, token.Name) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   errors.New("Action Not Authorised").Error(),
		})
	}

	return c.Next()
}
