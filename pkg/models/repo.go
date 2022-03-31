package models

import (
	"context"
	"log"

	"github.com/cave/pkg/database"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//errHandlerNotSet error = errors.New("handler not set properly")
	RdDB *redis.Client
	db   *mongo.Database

	Categories = []Category{
		{Title: "Systematic Theology I", Semester: 1},
		{Title: "Systematic Theology II", Semester: 1},
		{Title: "Old Testament Survey", Semester: 1},
		{Title: "New Testament Survey", Semester: 1},
		{Title: "Church History Survey", Semester: 1},
		{Title: "Hermeneutics", Semester: 1},
		{Title: "HOmiletics", Semester: 1},
		{Title: "Research & Writing", Semester: 1},
		{Title: "Apostolic Discipleship & Mentoring", Semester: 1},
		{Title: "Introduction To Philosophy", Semester: 1},
		{Title: "Christian Ethics", Semester: 2},
		{Title: "Christian Family", Semester: 2},
		{Title: "Christian Apologetics", Semester: 2},
		{Title: "Introduction To Islam", Semester: 2},
		{Title: "Cross Cultural Missions", Semester: 2},
		{Title: "Principles Of Counseling", Semester: 2},
		{Title: "Children Ministry", Semester: 2},
		{Title: "Youth Ministry", Semester: 2},
		{Title: "Leadership & Administration", Semester: 2},
		{Title: "Africa Traditional Religion (ATR) & World Religions", Semester: 2},
	}
)

// SetRepoDB global db instances
func SetRepoDB() {
	db = database.MgDB.Db
	RdDB = database.RdDB

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	if _, err := db.Collection("users").Indexes().CreateOne(context.Background(), index); err != nil {
		log.Println("Could not create index:", err)
	}
}
