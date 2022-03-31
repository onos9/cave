package lms

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	targetGroupCol = "target_groups"
)

// TargetGroup defines a model for a group of targets in a level
type TargetGroup struct {
	utils.Base
	Name        string `bson:"name,omitempty" json:"name"`
	Description string  `bson:"description,omitempty" json:"description"`
	SortIndex   int `bson:"sortIndex,omitempty" json:"sortIndex"`
	Milestone   bool `bson:"milestone,omitempty" json:"milestone"`
	LevelID     string     `bson:"levelId,omitempty" json:"levelId"`
	Level       Level     `bson:"level,omitempty" json:"level"`
	Targets     TargetList `bson:"targets,omitempty" json:"targets"`
}



// TargetGroupList defines array of target group objects
type TargetGroupList []*TargetGroup





/**
CRUD functions
*/

// Create creates a new TargetGroup record
func (m *TargetGroup) Create() error {
	_, err := db.Collection(targetGroupCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches TargetGroup by id
func (m *TargetGroup) FetchByID() error {
	err := db.Collection(targetGroupCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *TargetGroup) FetchAll(cl *TargetGroupList) error {
	cursor, err := db.Collection(targetGroupCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given TargetGroup
func (m *TargetGroup) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(targetGroupCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes TargetGroup by id
func (m *TargetGroup) Delete() error {
	_, err := db.Collection(targetGroupCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *TargetGroup) DeleteMany() error {
	_, err := db.Collection(targetGroupCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
