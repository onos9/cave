package lms

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	quizQuestionCol = "quiz_questions"
)

// QuizQuestion defines a model for questions in a quiz
type QuizQuestion struct {
	utils.Base
	QuizID          string `bson:"quizId,omitempty" json:"quizId"`
	Question        string `bson:"question,omitempty" json:"question"`
	Description     string `bson:"description,omitempty" json:"description"`
	CorrectAnswerID string
	Quiz            Quiz              `bson:"quiz,omitempty" json:"quiz"`
	Answer          *AnswerOption     `bson:"answer,omitempty" json:"answer"`
	Answers         AnswerOptionList  `bson:"answers,omitempty" json:"answers"`
}



// QuizQuestionList defines array of quiz question objects
type QuizQuestionList []*QuizQuestion


/**
CRUD functions
*/

// Create creates a new quizQuestionCol record
func (m *QuizQuestion) Create() error {
	_, err := db.Collection(quizQuestionCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches QuizQuestion by id
func (m *QuizQuestion) FetchByID() error {
	err := db.Collection(quizQuestionCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *QuizQuestion) FetchAll(cl *QuizQuestionList) error {
	cursor, err := db.Collection(quizQuestionCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given quizQuestionCol
func (m *QuizQuestion) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(quizQuestionCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes quizQuestionCol by id
func (m *QuizQuestion) Delete() error {
	_, err := db.Collection(quizQuestionCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *QuizQuestion) DeleteMany() error {
	_, err := db.Collection(quizQuestionCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}


