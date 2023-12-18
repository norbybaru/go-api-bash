package server

import (
	"dancing-pony/internal/app/auth"
	"dancing-pony/internal/app/dish"
	"dancing-pony/internal/app/rating"
	"dancing-pony/internal/platform/config"
	"dancing-pony/internal/platform/database"
	"dancing-pony/internal/platform/migration"
	"dancing-pony/internal/platform/session"
	"database/sql"
	"errors"

	_ "dancing-pony/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type App struct {
	Fiber   *fiber.App
	Store   *database.DB
	Session session.Storage
}

// Instantiate new application
func NewApp() *App {
	config.BootstrapConfig()
	fiber := fiber.New(fiber.Config{
		ErrorHandler: fiberErrorHandler,
	})
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
	app.registerApiRoutes()
	app.LoadDefaultMiddleware()

	RunDbMigrations()

	if config.App.Env == "local" {
		StartServer(app.Fiber)
	} else {
		StartServerWithGracefulShutdown(app.Fiber)
	}
}

// Register application routes
func (app *App) registerApiRoutes() {
	router := app.Fiber.Group("/api")
	auth.RegisterApiRoutes(router, app.Store, app.Session)
	dish.RegisterApiRoutes(router, app.Store)
	rating.RegisterApiRoutes(router, app.Store)
}

// Register default public routes
func (app *App) registerDefaultRoutes() {
	app.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "ok",
		})
	})

	app.Fiber.Get("/docs/*", swagger.HandlerDefault) // default

	app.Fiber.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}

// Run Database migration on application start
func RunDbMigrations() {
	migration.RunMigrations(config.Database.Source)
}

// Load application default middleware
func (app *App) LoadDefaultMiddleware() {
	app.Fiber.Use(
		cors.New(),
		logger.New(),
		recover.New(),
	)
}

// Global error handler
func fiberErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   errors.New("Resource not found").Error(),
		})
	}
	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Return from controller
	return nil
}
