package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	courseAuthorCol = "course_authors"
)

// CourseAuthor defines a model for course authors
type CourseAuthor struct {
	utils.Base
	CourseAuthorID string          `bson:"courseAuthorId,omitempty" json:"courseAuthorId"`
	UserID         string          `bson:"userId,omitempty" json:"userId"`
	Course         Course          `bson:"course,omitempty" json:"course"`
	User           User            `bson:"user,omitempty" json:"user"`
	Certificates   CertificateList `bson:"certificates,omitempty" json:"certificates"`
}

// CourseAuthorList defines array of course author objects
type CourseAuthorList []*CourseAuthor

/**
CRUD functions
*/

// Create creates a new courseAuthor record
func (m *CourseAuthor) Create() error {
	_, err := db.Collection(courseAuthorCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches CourseAuthor by id
func (m *CourseAuthor) FetchByID() error {
	err := db.Collection(courseAuthorCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *CourseAuthor) FetchAll(cl *CourseAuthorList) error {
	cursor, err := db.Collection(courseAuthorCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given courseAuthor
func (m *CourseAuthor) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(courseAuthorCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes courseAuthor by id
func (m *CourseAuthor) Delete() error {
	_, err := db.Collection(courseAuthorCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *CourseAuthor) DeleteMany() error {
	_, err := db.Collection(courseAuthorCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
