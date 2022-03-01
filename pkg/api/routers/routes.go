package routers

import (
	"github.com/cave/pkg/api/controllers"
	"github.com/cave/pkg/api/controllers/employe"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Route struct {
	Name    string
	Group   string
	Handler func(*fiber.Ctx)
	// Get    func(*fiber.Ctx)
	// Update func(*fiber.Ctx)
	// Delete func(*fiber.Ctx)
	// List   func(*fiber.Ctx)
}


type Routes []Route

var c = controllers.Controller{}
var routes = Routes{
	Route{Name: c.Employe.Name, Group: c.Employe.Group Handler: },
}

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Index)

	for _, route := range routes {
		router := api.Group(route)
		registerSubRoutes(router)
	}
}

func registerSubRoutes(route fiber.Router) {
	route.Get("/", employe.GetAll)
	route.Post("/", employe.CreateNew)
	route.Delete("/:id", employe.DeleteSingle)
	route.Delete("/", employe.DeleteAll)
	route.Get("/:id", employe.GetSingle)
	route.Put("/:id", employe.UpdateSingle)
}
