package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type DB struct {
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
}

func NewDBConnection() *DB {
	mongo, mongoClient := connectMongo()
	return &DB{
		MongoClient: mongoClient,
		MongoDB:     mongo,
	}
}

// ConnectMongo Returns the Mongo DB Instance
func connectMongo() (*mongo.Database, *mongo.Client) {
	opts := options.Client().ApplyURI(GetConfig().Mongo.URI)
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
	mongo := client.Database(GetConfig().Mongo.MongoDBName)
	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Establish Mongo DB Connection")
		fmt.Println(strings.Repeat("!", 40))

		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("😀 Connected To Mongo DB")
	fmt.Println(strings.Repeat("-", 40))

	return mongo, client
}

// ConnectRedis returns the Redis Instance
func RedisClient(dbn int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", GetConfig().Redis.HOST, GetConfig().Redis.PORT),
		Password: "",  // no password set
		DB:       dbn, // database number
	})
	return client
}
