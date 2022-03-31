package lms

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	evaluationCriteriaCol = "evaluation_criterias"
)

// EvaluationCriteria defines a model for course evaluation criteria
type EvaluationCriteria struct {
	utils.Base
	Name        string        `bson:"name,omitempty" json:"name"`
	CourseID    string        `bson:"courseId,omitempty" json:"courseId"`
	MaxGrade    uint          `bson:"maxGrade,omitempty" json:"maxGrade"`
	PassGrade   uint          `bson:"passGrade,omitempty" json:"passGrade"`
	GradeLabels []interface{} `bson:"gradeLabels,omitempty" json:"gradeLabels"`
	Course      Course        `bson:"course,omitempty" json:"course"`
}

// EvaluationCriteriaList defines array of evaluation criteria objects
type EvaluationCriteriaList []*EvaluationCriteria

/**
CRUD functions
*/

// Create creates a new evaluationCriteria record
func (m *EvaluationCriteria) Create() error {
	_, err := db.Collection(evaluationCriteriaCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches EvaluationCriteria by id
func (m *EvaluationCriteria) FetchByID() error {
	err := db.Collection(evaluationCriteriaCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *EvaluationCriteria) FetchAll(cl *EvaluationCriteriaList) error {
	cursor, err := db.Collection(evaluationCriteriaCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given evaluationCriteria
func (m *EvaluationCriteria) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(evaluationCriteriaCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes evaluationCriteria by id
func (m *EvaluationCriteria) Delete() error {
	_, err := db.Collection(evaluationCriteriaCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *EvaluationCriteria) DeleteMany() error {
	_, err := db.Collection(evaluationCriteriaCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
