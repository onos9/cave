package models

import (
	"context"

	"github.com/cave/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// SetRepoDB global db handler
func SetRepoDB(dbi *database.DB) {

	db = dbi.MongoDB

	SetIndex("users", "email")
}

func SetIndex(doc, field string) error {

	coll := db.Collection(doc)
	count, err := coll.CountDocuments(context.Background(), bson.M{})
	if err != nil && count == 0 {
		return err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{field: 1}, Options: opt}
	if _, err := coll.Indexes().CreateOne(context.Background(), index); err != nil {
		return err
	}

	return nil
}
