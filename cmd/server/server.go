package server

import (
	"dancing-pony/internal/config"
	"dancing-pony/internal/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber *fiber.App
	db    *database.DB
}

func NewApp() *App {
	config.BootstrapConfig()
	app := fiber.New()
	dbConnection := database.Init(config.Database.Source, config.Database.Driver)

	return &App{
		fiber: app,
		db:    dbConnection,
	}
}

func (app *App) Start() {
	app.fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.db.Ping()

	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}

	PORT := os.Getenv("PORT")

	log.Fatal(app.fiber.Listen(":" + PORT))
}
