package models

import (
	"context"
	"log"
	"os"

	"github.com/cave/config"
	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// SetRepoDB global db handler
func SetRepoDB(dbi *config.DB) {
	db = dbi.MongoDB

	err := SetIndex("users", "email")
	if err != nil {
		log.Panic(err.Error())
	}

	// Hash Password
	password := os.Getenv("ADMIN_PASS")
	hashedPass, _ := utils.EncryptPassword(password)
	user := User{
		Email:        os.Getenv("ADMIN_EMAIL"),
		PasswordHash: []byte(hashedPass),
		Role:         "admin",
		IsVerified:   false,
	}

	//Save User To DB
	if err := user.Create(); err != nil {
		e := err.(mongo.WriteException)
		if c := e.WriteErrors[0].Code; c != 11000 {
			log.Panic(err.Error())
		}
	}
}

func SetIndex(doc, field string) error {

	coll := db.Collection(doc)
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{field: 1}, Options: opt}
	if _, err := coll.Indexes().CreateOne(context.Background(), index); err != nil {
		return err
	}

	return nil
}
