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

	app := fiber.New()

	/*
		========== Setup Configs ============
	*/

	cfg.LoadConfig()
	config := cfg.GetConfig()

	/*
		========== Setup DB ============
	*/

	// Connect to Postgres
	// db.ConnectPostgres()

	// Drop on serve restarts in dev
	// db.PgDB.Migrator().DropTable(&user.User{})

	// Migration
	// db.PgDB.AutoMigrate(&user.User{})
	//migrations.Migrate(db.DB)

	// Connect to Mongo
	db.ConnectMongo()

	// Connect to Redis
	db.ConnectRedis()

	zoho.NewMailer(config)

	/*
		============ Set Up Middlewares ============
	*/

	// Default Log Middleware
	app.Use(logger.New())

	// Recovery Middleware
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Accept, Origin, Content-Type",
	}))

	/*
		============ Set Up Routes ============
	*/
	controller.SetupRoutes(app)

	/*
		============ Setup Swagger ===============
	*/

	// FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
