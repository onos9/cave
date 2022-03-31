package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	targetVersionCol = "target_versions"
)

// TargetVersion defines a model for a specific version of a target
type TargetVersion struct {
	utils.Base
	TargetID      string           `bson:"targetId,omitempty" json:"targetId"`
	VersionName   string           `bson:"versionName,omitempty" json:"versionName"`
	Target        Target           `bson:"target,omitempty" json:"target"`
	ContentBlocks ContentBlockList `bson:"contentBlocks,omitempty" json:"contentBlocks"`
}

// TargetVersionList defines array of target version objects
type TargetVersionList []*TargetVersion

/**
CRUD functions
*/

// Create creates a new TargetVersion record
func (m *TargetVersion) Create() error {
	_, err := db.Collection(targetVersionCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches TargetVersion by id
func (m *TargetVersion) FetchByID() error {
	err := db.Collection(targetVersionCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *TargetVersion) FetchAll(cl *TargetVersionList) error {
	cursor, err := db.Collection(targetVersionCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given TargetVersion
func (m *TargetVersion) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(targetVersionCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes TargetVersion by id
func (m *TargetVersion) Delete() error {
	_, err := db.Collection(targetVersionCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *TargetVersion) DeleteMany() error {
	_, err := db.Collection(targetVersionCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
