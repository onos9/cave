package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	// Configs
	cfg "github.com/cave/config"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MgDB MongoInstance
	RdDB *redis.Client
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// ConnectMongo Returns the Mongo DB Instance
func ConnectMongo() {
	opts := options.Client().ApplyURI(cfg.GetConfig().Mongo.URI)
	client, err := mongo.NewClient(opts)
	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Create Mongo DB Client")
		fmt.Println(strings.Repeat("!", 40))

		log.Fatal(err)
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

	MgDB = MongoInstance{
		Client: client,
		Db:     db,
	}
}

// ConnectRedis returns the Redis Instance
func ConnectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetConfig().Redis.HOST, cfg.GetConfig().Redis.PORT),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping(client.Context()).Result()

	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Establish Redis Connection")
		fmt.Println(strings.Repeat("!", 40))
		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("😀 Connected To Redis: %s\n", pong)
	fmt.Println(strings.Repeat("-", 40))

	RdDB = client
}
