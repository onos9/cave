package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	quizCol = "quizzes"
)

// Quiz defines a model for target quizes
type Quiz struct {
	utils.Base
	Title         string           `bson:"title,omitempty" json:"title"`
	TargetID      string           `bson:"targetId,omitempty" json:"targetId"`
	Target        Target           `bson:"target,omitempty" json:"target"`
	QuizQuestions QuizQuestionList `bson:"quizQuestions,omitempty" json:"quizQuestions"`
}

// QuizList defines array of quiz objects
type QuizList []*Quiz

/**
CRUD functions
*/

// Create creates a new user record
func (m *Quiz) Create() error {
	_, err := db.Collection(quizCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Quiz by id
func (m *Quiz) FetchByID() error {
	err := db.Collection(quizCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Quiz) FetchAll(cl *QuizList) error {
	cursor, err := db.Collection(quizCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (m *Quiz) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(quizCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *Quiz) Delete() error {
	_, err := db.Collection(quizCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Quiz) DeleteMany() error {
	_, err := db.Collection(quizCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
