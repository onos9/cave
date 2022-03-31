package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	quizUserAnswerCol = "quiz_question_user_answers"
)

// QuizQuizUserAnswerAnswer defines a model for questions in a quiz
type QuizUserAnswer struct {
	utils.Base
	QuestionID       string        `bson:"questionId,omitempty" json:"questionId"`
	AnswerID         string        `bson:"answerId,omitempty" json:"answerId"`
	QuizUserAnswerID string        `bson:"quizUserAnswerId,omitempty" json:"quizUserAnswerId"`
	Question         *QuizQuestion `bson:"question,omitempty" json:"question"`
	Answer           *AnswerOption `bson:"answer,omitempty" json:"answer"`
	User             *User         `bson:"user,omitempty" json:"user"`
}

// QuizQuizUserAnswerAnswerList defines array of quiz user answer objects
type QuizUserAnswerList []*QuizUserAnswer

/**
CRUD functions
*/

// Create creates a new user record
func (m *QuizUserAnswer) Create() error {
	_, err := db.Collection(quizUserAnswerCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches QuizUserAnswer by id
func (m *QuizUserAnswer) FetchByID() error {
	err := db.Collection(quizUserAnswerCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *QuizUserAnswer) FetchAll(cl *QuizUserAnswerList) error {
	cursor, err := db.Collection(quizUserAnswerCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (m *QuizUserAnswer) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(quizUserAnswerCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *QuizUserAnswer) Delete() error {
	_, err := db.Collection(quizUserAnswerCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *QuizUserAnswer) DeleteMany() error {
	_, err := db.Collection(quizUserAnswerCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
