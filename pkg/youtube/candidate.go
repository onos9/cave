package youtube

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Candidate is a model for Candidates table
type Candidate struct {
	utils.Base
	Email        string `gorm:"type:varchar(100);unique_index" json:"email" `
	PasswordHash []byte `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`

	Data          string        `gorm:"type:text" json:"data"`
	Qualification Qualification `json:"qualification"`
	Background    Background    `json:"background"`
	Health        Health        `json:"health"`
	Referee       Health        `json:"referee"`
	Terms         Health        `json:"terms"`
}

// CandidateList defines array of candidate objects
type CandidateList []*Candidate

/**
CRUD functions
*/

// Create creates a new candidate record
func (m *Candidate) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Candidate by id
func (m *Candidate) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Candidate) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given candidate
func (m *Candidate) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes candidate by id
func (m *Candidate) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Candidate) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
