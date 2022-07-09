package controller

import (
	// Middlewares

	"github.com/cave/config"
	"github.com/cave/pkg/middlewares"
	"github.com/cave/pkg/models"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type Resp map[string]interface{}

// SetupRoutes setups r
func SetupRoutes(r *fiber.App, db *config.DB) {
	models.SetRepoDB(db)
	cfg := config.GetConfig()

	r.Static("/", cfg.Webroot)
	r.Static("/static/*", cfg.Webroot+"/index.html")

	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.Use("/docs", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		baseUrl := c.BaseURL()
		return c.JSON(fiber.Map{
			"message": "Welcome to Cave API v1",
			"url":     baseUrl,
		})
	})

	hook := v1.Group("/webhook")
	hook.Post("/mail", webhook.payment)

	// Auth Group
	auth := v1.Group("/auth")
	auth.Get("/", userAuth.token)
	auth.Post("/", userAuth.signin)
	auth.Post("/temp", userAuth.tempSignup)
	auth.Post("/:token", userAuth.verify)
	auth.Put("/", userAuth.signup)
	auth.Delete("/", userAuth.signout)

	// User Group
	u := v1.Group("/user")
	u.Post("/", middlewares.RequireLoggedIn(), user.create)
	u.Get("/", middlewares.RequireLoggedIn(), user.getAll)
	u.Get("/:id", middlewares.RequireLoggedIn(), user.getOne)
	u.Patch("/:id", middlewares.RequireLoggedIn(), user.updateOne)
	u.Delete("/:id", middlewares.RequireLoggedIn(), user.deleteOne)

	// Mail Routes
	m := v1.Group("/mail")
	m.Post("/", middlewares.RequireLoggedIn(), email.send)
	m.Put("/", email.send)
	m.Get("/", middlewares.RequireLoggedIn(), email.zohoCode)
	m.Post("/token", middlewares.RequireLoggedIn(), email.token)

	// m.Get("/", middlewares.RequireLoggedIn(), email.getAll)
	// m.Get("/:id", middlewares.RequireLoggedIn(), email.getOne)
	// m.Post("/:id", middlewares.RequireLoggedIn(), email.updateOne)
	// m.Delete("/:id", middlewares.RequireLoggedIn(), email.deleteOne)

	// Downloads Group
	downloads := v1.Group("/file")
	downloads.Get("/download/:filename", file.download)
	downloads.Post("/upload", file.upload)

	// LogBook Routes
	lb := v1.Group("/logbook")
	lb.Post("/", middlewares.RequireLoggedIn(), logBook.create)
	lb.Get("/", middlewares.RequireLoggedIn(), logBook.getAll)
	lb.Get("/:id", middlewares.RequireLoggedIn(), logBook.getOne)
	lb.Patch("/:id", middlewares.RequireLoggedIn(), logBook.updateOne)
	lb.Delete("/:id", middlewares.RequireLoggedIn(), logBook.deleteOne)
}
