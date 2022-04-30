package controller

import (
	// Middlewares
	"github.com/cave/pkg/database"
	"github.com/cave/pkg/middlewares"
	"github.com/cave/pkg/models"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type Resp map[string]interface{}

// SetupRoutes setups router
func SetupRoutes(app *fiber.App, db *database.DB) {
	models.SetRepoDB(db)

	// serve Single Page application on "/web" route
	// assume static file at dist folder
	app.Static("/web", "web/dist")

	// serve the 'index.html', if there ids no matching routes.
	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./web/dist/index.html")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Use("/docs", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API Version One Home",
		})
	})

	// Auth Group
	auth := app.Group("/auth")
	auth.Get("/", userAuth.token)
	auth.Post("/", userAuth.signin)
	auth.Post("/:token", userAuth.verify)
	auth.Put("/", userAuth.signup)
	auth.Delete("/", userAuth.signout)

	// User Group
	u := app.Group("/user")
	u.Post("/", middlewares.RequireLoggedIn(), user.create)
	u.Get("/", middlewares.RequireLoggedIn(), user.getAll)
	u.Get("/:id", middlewares.RequireLoggedIn(), user.getOne)
	u.Patch("/:id", middlewares.RequireLoggedIn(), user.updateOne)
	u.Delete("/:id", middlewares.RequireLoggedIn(), user.deleteOne)

	// Mail Routes
	m := app.Group("/mail")
	m.Post("/", middlewares.RequireLoggedIn(), mailer.send)
	m.Get("/", middlewares.RequireLoggedIn(), mailer.zohoCode)
	m.Post("/token", middlewares.RequireLoggedIn(), mailer.token)
	m.Get("/", middlewares.RequireLoggedIn(), mailer.zohoCode)
	
	// m.Get("/", middlewares.RequireLoggedIn(), mail.getAll)
	// m.Get("/:id", middlewares.RequireLoggedIn(), mail.getOne)
	// m.Post("/:id", middlewares.RequireLoggedIn(), mail.updateOne)
	// m.Delete("/:id", middlewares.RequireLoggedIn(), mail.deleteOne)
}
