package main

import (
	"fmt"

	"github.com/cave/cmd/api/controller"
	cfg "github.com/cave/config"
	db "github.com/cave/pkg/database"
	"github.com/cave/pkg/zoho"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {

	// Setup configs
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// Setup Adapters
	db.ConnectMongo()
	db.ConnectRedis()
	zoho.NewMailer(config)

	// Setup fiber api
	app := fiber.New()

	// Set Up Middlewares
	app.Use(logger.New())   // Default Log Middleware
	app.Use(recover.New())  // Recovery Middleware

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Accept, Origin, Content-Type, Authorization",
	}))

	// Setup Routes
	controller.SetupRoutes(app)

	//Setup Swagger 
	// FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
