package config

import (
	"dancing-pony/internal/common/utils"
	"os"
)

type jwtConfig struct {
	Secret        string
	ContextKey    string
	ExpireMinutes int
}

var JWT *jwtConfig

func loadJWTConfig() {
	JWT = &jwtConfig{
		Secret:        os.Getenv("JWT_SECRET_KEY"),
		ContextKey:    os.Getenv("JWT_CONTEXT_KEY"),
		ExpireMinutes: utils.ParseInt(os.Getenv("JWT_EXPIRE_MINUTES"), 60),
	}
}
