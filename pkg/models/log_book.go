package models

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Database table for LogBook
	logBookCol = "logbooks"
)

// LogBook struct for users table
type LogBook struct {
	utils.Base
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	CourseCode string             `bson:"courseCode,omitempty" json:"courseCode,omitempty"`
	CourseName string             `bson:"courseName,omitempty" json:"courseName,omitempty"`
	Group      string             `bson:"group,omitempty" json:"group,omitempty"`
	UserID     string             `bson:"userID,omitempty" json:"userID,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	Prayer     []*Prayer          `bson:"prayer,omitempty" json:"prayer,omitempty"`
	Evangelism []*Evangelism      `bson:"evangelism,omitempty" json:"evangelism,omitempty"`
	Exercise   []*Exercise        `bson:"exercise,omitempty" json:"exercise,omitempty"`
}

type Evangelism struct {
	Converts    string                   `bson:"converts,omitempty" json:"converts"`
	Location    string                   `bson:"location,omitempty" json:"location,omitempty"`
	Date        string                   `bson:"date,omitempty" json:"date,omitempty"`
	Testimonies string                   `bson:"testimonies,omitempty" json:"testimonies,omitempty"`
	ConvertInfo []map[string]string `bson:"convertInfo,omitempty" json:"convertInfo"`
}

type Prayer struct {
	Location    string `bson:"location,omitempty" json:"location,omitempty"`
	Date        string `bson:"date,omitempty" json:"date,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type Exercise struct {
	Chapters     string `bson:"chapters,omitempty" json:"chapters,omitempty"`
	EndChapter   string `bson:"endChapter,omitempty" json:"endChapter,omitempty"`
	StartChapter string `bson:"startChapter,omitempty" json:"startChapter,omitempty"`
	Day          string `bson:"day,omitempty" json:"day,omitempty"`
	Author       string `bson:"author,omitempty" json:"author,omitempty"`
	BookTitle    string `bson:"bookTitle,omitempty" json:"bookTitle,omitempty"`
	PrayerTime   string `bson:"prayerTime,omitempty" json:"prayerTime"`
	NoPages      string `bson:"noPages,omitempty" json:"noPages"`
}

// LogBookList defines array of logBook objects
type LogBookList []*LogBook

/**
CRUD functions
*/

// Create creates a new logBook record
func (m *LogBook) Create() error {
	t := time.Now()
	m.CreatedAt = &t
	m.Id = primitive.NewObjectID()

	result, err := db.Collection(logBookCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	m.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

// FetchByID fetches LogBook by id
func (m *LogBook) FetchByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = db.Collection(logBookCol).FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByEmail fetches LogBook by userID
func (m *LogBook) FetchByUserId(ul *LogBookList) error {
	cursor, err := db.Collection(logBookCol).Find(context.TODO(), bson.M{"userID": m.UserID})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), ul); err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all LogBook
func (m *LogBook) FetchAll(ul *LogBookList) error {
	cursor, err := db.Collection(logBookCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), ul); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given logBook
func (m *LogBook) UpdateOne() error {
	t := time.Now()
	m.UpdatedAt = &t

	bm, err := bson.Marshal(m)
	if err != nil {
		return err
	}

	var val bson.M
	err = bson.Unmarshal(bm, &val)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": m.Id}
	update := bson.D{{Key: "$set", Value: val}}
	_, err = db.Collection(logBookCol).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes logBook by id
func (m *LogBook) Delete() error {
	t := time.Now()
	m.DeletedAt = &t
	_, err := db.Collection(logBookCol).DeleteOne(context.TODO(), bson.M{"_id": m.Id})
	if err != nil {
		return err
	}
	return nil
}

func (m *LogBook) DeleteMany() error {
	_, err := db.Collection(logBookCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
