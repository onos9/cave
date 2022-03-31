package lms

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	targetCol = "targets"
)

// Target defines a model for target groups single target (equivalent to lesson)
type Target struct {
	utils.Base
	Role                   string            `bson:"role,omitempty" json:"role"`
	Title                  string            `bson:"title,omitempty" json:"title"`
	Description            string            `bson:"description,omitempty" json:"description"`
	CompletionInstructions string            `bson:"completionInstructions,omitempty" json:"completionINstructions"`
	ResourceURL            string            `bson:"resourceURL,omitempty" json:"resourceURL"`
	TargetGroupID          string            `bson:"targetGroupId,omitempty" json:"targetGroupId"`
	SortIndex              int               `bson:"sortIndex,omitempty" json:"sortIndex"`
	SessionAt              *time.Time        `bson:"sessionAt,omitempty" json:"sessionAt"`
	LinkToComplete         string            `bson:"linkToComplete,omitempty" json:"linkToComplete"`
	Resubmittable          bool              `bson:"resubmittable,omitempty" json:"resubmittable"`
	CheckList              interface{}       `bson:"checkList,omitempty" json:"checkList"`
	ReviewChecklist        interface{}       `bson:"reviewCheckList,omitempty" json:"reviewCheckList"`
	TargetGroup            TargetGroup       `bson:"targetGroup,omitempty" json:"targetGroup"`
	TargetVersions         TargetVersionList `bson:"targetVersions,omitempty" json:"targetVersions"`
	Quizzes                QuizList          `bson:"quizzes,omitempty" json:"quizzes"`
}

// TargetList defines array of target objects
type TargetList []*Target

/**
CRUD. functions
*/

// Create creates a new Target record
func (m *Target) Create() error {
	_, err := db.Collection(targetCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Target by id
func (m *Target) FetchByID() error {
	err := db.Collection(targetCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Target) FetchAll(cl *TargetList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given Target
func (m *Target) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(targetCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes Target by id
func (m *Target) Delete() error {
	_, err := db.Collection(targetCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Target) DeleteMany() error {
	_, err := db.Collection(targetCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
