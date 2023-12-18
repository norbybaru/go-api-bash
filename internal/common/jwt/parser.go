package jwt

import (
	"dancing-pony/internal/platform/config"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenMetadata struct {
	Identifier interface{}
	Expires    int64
	Email      string
	NickName   string
	Name       string
}

func (t *TokenMetadata) GetIdentifier() int {
	return int(t.Identifier.(float64))
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		identifier := claims["id"]
		email := fmt.Sprint(claims["email"])
		nickname := fmt.Sprint(claims["nickname"])
		name := fmt.Sprint(claims["name"])
		expires := int64(claims["expires"].(float64))

		return &TokenMetadata{
			Identifier: identifier,
			Expires:    expires,
			Email:      email,
			NickName:   nickname,
			Name:       name,
		}, nil
	}

	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.JWT.Secret), nil
}

// Extract access_token from request HTTP header
func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")

	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}
