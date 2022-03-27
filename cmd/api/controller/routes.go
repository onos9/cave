package controller

import (
	// Middlewares
	"github.com/cave/cmd/api/mods"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {
	mods.SetRepoDB()

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")

	// app.Get("/web/*", func(ctx *fiber.Ctx) error {
	// 	return ctx.SendFile("./dist/index.html")
	// })

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Use("/docs", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API Version One Home",
		})
	})

	// v1.Get("/home", ctl.HomeController)

	// Auth Group
	auth := v1.Group("/auth")
	auth.Post("/register", candidate.register)
	auth.Post("/login", candidate.login)

	// Candidates
	cand := v1.Group("/candidate")
	cand.Post("/", candidate.create)
	cand.Get("/", candidate.getAll)
	cand.Get("/:_id", candidate.google)

	// Candidates
	mail := v1.Group("/mail")
	mail.Post("/", mailer.sendMail)
	mail.Get("/", mailer.credential)
	mail.Post("/code", mailer.code)
	mail.Get("/token", mailer.token)

	// Authenticated Routes
	//mail.Post("/", middlewares.RequireLoggedIn(), candidate.create)
}
