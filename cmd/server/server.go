package server

import (
	"dancing-pony/internal/app/auth"
	"dancing-pony/internal/app/dish"
	"dancing-pony/internal/app/rating"
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/migration"
	"dancing-pony/internal/platform/session"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber   *fiber.App
	Store   *database.DB
	Session session.Storage
}

var Application *App

// Initialize application and load necessary configs
func NewApp() *App {
	config.BootstrapConfig()
	fiber := fiber.New()
	dbConnection := database.Init(config.Database.Source, config.Database.Driver)
	session := session.Init(config.Session.Host, config.Session.Port, config.Session.Password)

	return &App{
		Fiber:   fiber,
		Store:   dbConnection,
		Session: session,
	}
}

// Start server
func (app *App) Start() {
	app.registerDefaultRoutes()
	app.registerDomainRoutes()
	RunDbMigrations()

	PORT := os.Getenv("PORT")

	log.Fatal(app.Fiber.Listen(":" + PORT))
}

// Register application routes
func (app *App) registerDomainRoutes() {
	auth.RegisterRoutes(app.Fiber, app.Store, app.Session)
	dish.RegisterRoutes(app.Fiber, app.Store)
	rating.RegisterRoutes(app.Fiber, app.Store)
}

// Register default public routes
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
