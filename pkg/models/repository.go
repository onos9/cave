package models

import (
	"github.com/cave/pkg/database"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//errHandlerNotSet error = errors.New("handler not set properly")
	RdDB *redis.Client
	db   *mongo.Database
)

// SetRepoDB global db handler
func SetRepoDB() {
	db = database.MgDB.Db
	RdDB = database.RdDB

	opt := options.Index()
	opt.SetUnique(true)

	// index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	// if _, err := db.Collection("users").Indexes().CreateOne(context.Background(), index); err != nil {
	// 	log.Println("Could not create index:", err)
	// }
}
