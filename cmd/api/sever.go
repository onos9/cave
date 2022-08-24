package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cave/cmd/api/controller"
	cfg "github.com/cave/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/swag/example/basic/docs"
)

type ApiServer struct {
	Server *fiber.App
}

const idleTimeout = 5 * time.Second

func Init(db *cfg.DB) *ApiServer {
	config := cfg.GetConfig()

	// Setup fiber api
	s := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		BodyLimit:   1000 * 1024 * 1024, // limit to 500MB
	})

	s.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir(config.Webroot),
		Browse: false, 
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))

	// Set Up Middlewares
	s.Use(logger.New())  // Default Log Middleware
	s.Use(recover.New()) // Recovery Middleware

	// cors
	s.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Range, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,DELETE,PUT,PATCH",
		ExposeHeaders:    "X-Total-Count, Content-Range",
	}))

	controller.SetupRoutes(s, db)

	//Setup Swagger
	// FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	return &ApiServer{s}
}

func (s *ApiServer) Run() {
	config := cfg.GetConfig()

	port := fmt.Sprintf(":%s", config.Port)
	go func() {
		if err := s.Server.Listen(port); err != nil {
			log.Panic(err)
		}
	}()
}

func (s *ApiServer) Shutdown() {
	fmt.Println("\n\nShutting down API service...")
	_ = s.Server.Shutdown()
}
