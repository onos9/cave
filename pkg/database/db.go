package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cfg "github.com/cave/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var Instance MongoInstance

// Create database connection
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetConfig().Mongo.URI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(cfg.GetConfig().Mongo.MongoDBName)

	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Establish Mongo DB Connection")
		fmt.Println(strings.Repeat("!", 40))

		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("😀 Connected To Mongo DB")
	fmt.Println(strings.Repeat("-", 40))

	Instance = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
