package config

import "os"

type databaseConfig struct {
	Name     string
	Driver   string
	Host     string
	Password string
	Port     string
	Source   string
	Username string
}

var Database *databaseConfig

func loadDatabaseEnv() {
	Database = &databaseConfig{
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Source:   os.Getenv("DB_SOURCE"),
		Username: os.Getenv("DB_USERNAME"),
	}
}
