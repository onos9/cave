package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	contentBlockCol = "content_blocks"
)

// ContentBlock defines model for blocks of a target (like markdown, image, embed, link...)
type ContentBlock struct {
	utils.Base
	BlockType       string        `bson:"blockType,omitempty" json:"blockType"`
	Content         []interface{} `bson:"content,omitempty" json:"content"`
	SortIndex       int           `bson:"sortIndex,omitempty" json:"sortIndex"`
	TargetVersionID string        `bson:"targetVersionId,omitempty" json:"targetVersionId"`
	TargetVersion   TargetVersion `bson:"targetVersion,omitempty" json:"targetVersion"`
}

// ContentBlockList defines array of content block objects
type ContentBlockList []*ContentBlock

/**
CRUD functions
*/

// Create creates a new contentBlock record
func (m *ContentBlock) Create() error {
	_, err := db.Collection(contentBlockCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches ContentBlock by id
func (m *ContentBlock) FetchByID() error {
	err := db.Collection(contentBlockCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *ContentBlock) FetchAll(cl *ContentBlockList) error {
	cursor, err := db.Collection(contentBlockCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given contentBlock
func (m *ContentBlock) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(contentBlockCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes contentBlock by id
func (m *ContentBlock) Delete() error {
	_, err := db.Collection(contentBlockCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *ContentBlock) DeleteMany() error {
	_, err := db.Collection(contentBlockCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
