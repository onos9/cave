package models

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	levelCol = "levels"
)

// Level defines a model for course levels
type Level struct {
	utils.Base
	Name         string          `bson:"name,omitempty" json:"name"`
	CourseID     string          `bson:"courseId,omitempty" json:"courseId"`
	Description  string          `bson:"description,omitempty" json:"description"`
	Number       int             `bson:"number,omitempty" json:"number"`
	UnlockOn     *time.Time      `bson:"unlockOn,omitempty" json:"unlockOn"`
	Course       Course          `bson:"course,omitempty" json:"course"`
	TargetGroups TargetGroupList `bson:"targetGroups,omitempty" json:"targetGroups"`
}

// LevelList defines array of level objects
type LevelList []*Level

/**
CRUD functions
*/

// Create creates a new level record
func (m *Level) Create() error {
	_, err := db.Collection(levelCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Level by id
func (m *Level) FetchByID() error {
	err := db.Collection(levelCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Level) FetchAll(cl *LevelList) error {
	cursor, err := db.Collection(levelCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given level
func (m *Level) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(levelCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes level by id
func (m *Level) Delete() error {
	_, err := db.Collection(levelCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Level) DeleteMany() error {
	_, err := db.Collection(levelCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
