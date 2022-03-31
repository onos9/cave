package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Health is a model for Healths table
type Health struct {
	utils.Base
	Disability             bool   `json:"disability"`
	Nervousillness         bool   `json:"nervousIll"`
	Anorexia               bool   `json:"anorexia"`
	Diabetese              bool   `json:"diabetese"`
	Epilepsy               bool   `json:"epilepsy"`
	StomachUlcers          bool   `json:"stomach_ulcers"`
	SpecialDiet            bool   `json:"special_diet"`
	LearningDisability     bool   `json:"learning_disability"`
	UsedIllDrug            bool   `json:"usedIllDrug"`
	DrugAddiction          bool   `json:"drug_addiction"`
	HadSurgery             bool   `json:"had_surgery"`
	HealthIssueDescription string `json:"healthIssueDesc"`

	// TableName gorm standard table name
	// func (c *Health) TableName() string {
	// 	return healthTableName
}

// HealthList defines array of health objects
type HealthList []*Health

/**
CRUD functions
*/

// Create creates a new health record
func (m *Health) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Health by id
func (m *Health) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Health) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given health
func (m *Health) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes health by id
func (m *Health) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Health) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
