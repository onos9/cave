package app

import (
	"fmt"

	// Configs
	cfg "github.com/cave/configs"
	"github.com/cave/pkg/api/routes"
	"github.com/swaggo/swag/example/basic/docs"

	// routes

	// database
	db "github.com/cave/pkg/database"

	// models

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run starts the app
// @title Gofiber Boilerplate API
// @version 1.0
// @description This is my gofiber boilerplate api server.
// @termsOfService http://swagger.io/terms/
// @contact.name Cozy
// @contact.url https://github.com/ItsCosmas
// @contact.email devcosmas@gmail.com
// @license.name MIT
// @license.url https://github.com/ItsCosmas/github.com/cave/blob/master/LICENSE
// @BasePath /api/v1
func Run() {
	// Setup Configs
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// connect to the database
	dbError := db.Connect()
	if dbError != nil {
		log.Fatal(dbError)
		return
	}

	app := fiber.New()

	// middlewares
	app.Use(compress.New())

	// Default Log Middleware
	app.Use(logger.New())

	// Recovery Middleware
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// routes setup
	routes.SetupRoutes(app)

	
	// Setup Swagger
	// Todo: FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
