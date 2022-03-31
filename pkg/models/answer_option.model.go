package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	answerOptionCol = "answer_options"
)

// AnswerOption is a model for options of Answers for a target
type AnswerOption struct {
	utils.Base
	QuizQuestionID string `bson:"quizQuestionId,omitempty" json:"quizQuestionId"`
	Value          string `bson:"value,omitempty" json:"value"`
	Hint           string `bson:"hint,omitempty" json:"hint"`

	QuizQuestion QuizQuestion `bson:"quizQuestion,omitempty" json:"quizQuestion"`
}

// AnswerOptionList defines array of answer options objects
type AnswerOptionList []*AnswerOption

/**
CRUD functions
*/

// Create creates a new answerOption record
func (m *AnswerOption) Create() error {
	_, err := db.Collection(answerOptionCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches AnswerOption by id
func (m *AnswerOption) FetchByID() error {
	err := db.Collection(answerOptionCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *AnswerOption) FetchAll(cl *AnswerOptionList) error {
	cursor, err := db.Collection(answerOptionCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given answerOption
func (m *AnswerOption) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(answerOptionCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes answerOption by id
func (m *AnswerOption) Delete() error {
	_, err := db.Collection(answerOptionCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *AnswerOption) DeleteMany() error {
	_, err := db.Collection(answerOptionCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
