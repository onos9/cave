package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/cave/cmd/api/handlers"
	"github.com/cave/configs"
	"github.com/cave/migrations"
	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/database"
	"github.com/cave/pkg/flag"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("go-lms-api" + " : ")
	log := log.New(os.Stdout, log.Prefix(), log.Flags())

	if err := envconfig.Process("go-lms-api", &configs.CFG); err != nil {
		log.Fatalf("main : Error Parsing Config file: %+v", err)
	}

	if err := flag.Process(&configs.CFG); err != nil {
		if err != flag.ErrHelp {
			log.Fatalf("main : Error Parsing Command Line : %+v", err)
		}
		// else provide help here
		return
	}

	// Print the config
	{
		cfgJSON, err := json.MarshalIndent(configs.CFG, "", "")
		if err != nil {
			log.Fatalf("main : Error marshaling config to JSON : %+v", err)
		}
		log.Printf("main : Config : %v\n", string(cfgJSON))
	}

	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database %+v", err)
	}
	log.Printf("%+v", dbConfig)
	db, err := database.Initialize(dbConfig.Storage)
	if err != nil {
		log.Fatalf("main: Error initializing database %+v", err)
	}

	defer db.DB.Close()
	defer db.Redis.Close()

	authenticator, _ := auth.NewAuthenticatorFile("", time.Now().UTC(), configs.CFG.Auth.KeyExpiration)

	migrations.Migrate(db.DB)

	app := gin.Default()
	handlers.ApplyRoutes(app, authenticator, db)
	app.Use(database.InjectDB(db))
	app.Run(configs.CFG.Server.Host)
}
