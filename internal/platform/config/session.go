package config

import "os"

type sessionConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database int
}

var Session *sessionConfig

func loadSessionEnvConfig() {
	Session = &sessionConfig{
		Driver:   "redis",
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
	}
}
