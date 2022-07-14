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
	Id primitive.ObjectID `json:"id" bson:"_id"`

	Email          string        `bson:"email,omitempty" json:"email,omitempty"`
	FullName       string        `bson:"fullName,omitempty" json:"fullName,omitempty"`
	MatricNumber   string        `bson:"matricNumber,omitempty" json:"matricNumber,omitempty"`
	Converts       string        `bson:"converts,omitempty" json:"converts"`
	Location       string        `bson:"location,omitempty" json:"location,omitempty"`
	PrayerLocation string        `bson:"prayerLocation,omitempty" json:"prayerLocation,omitempty"`
	PrayerWalk     []interface{} `bson:"prayerWalk,omitempty" json:"prayerWalk,omitempty"`
	ConvertInfo    []interface{} `bson:"convertInfo,omitempty" json:"convertInfo"`
	BibleRead      []interface{} `bson:"bibleRead,omitempty" json:"bibleRead"`
	Exercise       []interface{} `bson:"exercise,omitempty" json:"exercise"`
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

// FetchByEmail fetches User by email
func (m *LogBook) FetchByEmail() error {
	err := db.Collection(logBookCol).FindOne(context.TODO(), bson.M{"email": m.Email}).Decode(&m)
	if err != nil {
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
