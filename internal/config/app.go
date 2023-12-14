package config

import (
	"os"
	"strconv"
)

type appConfig struct {
	Env   string
	Debug bool
}

var App *appConfig

func loadAppEnv() {
	App = &appConfig{
		Env:   os.Getenv("APP_ENV"),
		Debug: getEnvDebug(),
	}
}

func getEnvDebug() bool {
	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	if err != nil {
		return false
	}

	return debug
}
