package server

import (
	"dancing-pony/internal/app/dish"
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/migration"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber *fiber.App
	Store *database.DB
}

func NewApp() *App {
	config.BootstrapConfig()
	fiber := fiber.New()
	dbConnection := database.Init(config.Database.Source, config.Database.Driver)

	app := &App{
		Fiber: fiber,
		Store: dbConnection,
	}

	return app
}

func (app *App) Start() {
	app.registerDefaultRoutes()
	app.registerDomainRoutes()
	RunDbMigrations()

	PORT := os.Getenv("PORT")

	log.Fatal(app.Fiber.Listen(":" + PORT))
}

func (app *App) registerDomainRoutes() {
	dish.RegisterRoutes(app.Fiber, app.Store)
}

func (app *App) registerDefaultRoutes() {
	app.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "ok",
		})
	})
}

func RunDbMigrations() {
	migration.RunMigrations(config.Database.Source)
}
