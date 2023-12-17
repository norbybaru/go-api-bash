package jwt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	Type    string `json:"token_type"`
	Expiry  int64  `json:"expires_in"`
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type jsonWebToken struct {
	Secret string
	Expiry time.Time
}

func Init(secret string, expiryMinutes int) *jsonWebToken {
	return &jsonWebToken{
		Secret: secret,
		Expiry: time.Now().Add(time.Minute * time.Duration(expiryMinutes)),
	}
}

// Generate a new Access & Refresh tokens.
func (j *jsonWebToken) GenerateNewTokens(identifier interface{}) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := j.generateNewAccessToken(identifier)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := j.generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Type:    "Bearer",
		Expiry:  j.Expiry.Unix(),
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (j *jsonWebToken) generateNewAccessToken(identifier interface{}) (string, error) {
	// Set secret key from .env file.
	secret := j.Secret
	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = identifier
	claims["expires"] = j.Expiry.Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *jsonWebToken) generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := j.Secret + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	day := time.Hour * 24

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(day * 1).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}
