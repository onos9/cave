package main

import (
	"log"
	"os"
	"time"

	"github.com/cave/cmd/graphql/resolvers"
	"github.com/cave/configs"
	"github.com/cave/migrations"
	"github.com/cave/pkg/auth"
	"github.com/cave/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("go-lms-graphql" + " : ")
	log := log.New(os.Stdout, log.Prefix(), log.Flags())

	if err := envconfig.Process("go-lms-api", &configs.CFG); err != nil {
		log.Fatalf("main : Error Parsing Config file: %+v", err)
	}

	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database %+v", err)
	}
	log.Printf("%+v", dbConfig)
	db, err := database.Initialize(dbConfig.Storage)
	defer db.Close()

	migrations.Migrate(db)

	authenticator, _ := auth.NewAuthenticatorFile("", time.Now().UTC(), configs.CFG.Auth.KeyExpiration)

	r := gin.Default()
	resolvers.ApplyResolvers(r, db, authenticator)

	log.Println("Running @http://" + configs.CFG.Server.Graphql + "/graphql")
	log.Fatalln(r.Run(configs.CFG.Server.Graphql))

}
