package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

const idleTimeout = 5 * time.Second

func main() {

	// Setup configs
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// Setup Adapters
	db.ConnectMongo()
	db.ConnectRedis()
	zoho.NewMailer(config)

	// Setup fiber api
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	// Set Up Middlewares
	app.Use(logger.New())  // Default Log Middleware
	app.Use(recover.New()) // Recovery Middleware

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Range, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,DELETE,PUT",
		ExposeHeaders:    "X-Total-Count, Content-Range",
	}))

	// Setup Routes
	controller.SetupRoutes(app)

	//Setup Swagger
	// FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)

	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	
	fmt.Println("\n\nShutting down server...")
	_ = app.Shutdown()
}
