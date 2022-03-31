package models

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	courseCol = "courses"
)

// Course is a model for Courses table
type Course struct {
	utils.Base
	Name                string                 `bson:"name,omitempty" json:"name"`
	EndsAt              *time.Time             `bson:"endsAt,omitempty" json:"endsAt"`
	Description         string                 `bson:"description,omitempty" json:"descripption"`
	EnableLeadboard     bool                   `bson:"enableLeadBoard,omitempty" json:"enableLeadBoard"`
	PublicSignup        bool                   `bson:"publicSignUp,omitempty" json:"publicSignUp"`
	Featured            bool                   `bson:"featured,omitempty" json:"featured"`
	About               string                 `bson:"about,omitempty" json:"about"`
	ProgressionBehavior string                 `bson:"progressionBehavior,omitempty" json:"progressionBehavior"`
	ProgressionLimit    int                    `bson:"progressionLimit,omitempty" json:"progressionLimit"`
	Certificates        CertificateList        `bson:"certificates,omitempty" json:"certificates"`
	Authors             CourseAuthorList       `bson:"authors,omitempty" json:"authors"`
	EvaluationCriteria  EvaluationCriteriaList `bson:"evaluationCriteria,omitempty" json:"evaluationCriteria"`
	Levels              LevelList              `bson:"levels,omitempty" json:"levels"`
	Students            StudentCourseList      `bson:"students,omitempty" json:"students"`
}

// CourseList defines array of course objects
type CourseList []*Course

/**
CRUD functions
*/

// Create creates a new course record
func (m *Course) Create() error {
	_, err := db.Collection(courseCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Course by id
func (m *Course) FetchByID() error {
	err := db.Collection(courseCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Course) FetchAll(cl *CourseList) error {
	cursor, err := db.Collection(courseCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given course
func (m *Course) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(courseCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes course by id
func (m *Course) Delete() error {
	_, err := db.Collection(courseCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Course) DeleteMany() error {
	_, err := db.Collection(courseCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
