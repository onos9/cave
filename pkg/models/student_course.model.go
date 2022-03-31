package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	studentCourseCol = "student_courses"
)

// StudentCourse defines a model for course students
type StudentCourse struct {
	utils.Base
	UserID   string `bson:"userId,omitempty" json:"userId"`
	CourseID string `bson:"courseId,omitempty" json:"courseId"`
	Course   Course `bson:"course,omitempty" json:"course"`
	User     User   `bson:"user,omitempty" json:"user"`
}

// StudentCourseList defines array of course student objects
type StudentCourseList []*StudentCourse

/**
CRUD functions
*/

// Create creates a new user record
func (m *StudentCourse) Create() error {
	_, err := db.Collection(studentCourseCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches StudentCourse by id
func (m *StudentCourse) FetchByID() error {
	err := db.Collection(studentCourseCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *StudentCourse) FetchAll(cl *StudentCourseList) error {
	cursor, err := db.Collection(studentCourseCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (m *StudentCourse) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(studentCourseCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *StudentCourse) Delete() error {
	_, err := db.Collection(studentCourseCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *StudentCourse) DeleteMany() error {
	_, err := db.Collection(studentCourseCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
