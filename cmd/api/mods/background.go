package mods

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// background is a model for backgrounds table
type Background struct {
	utils.Base
	BornAgain        bool   `json:"born_again"`
	SalvationBrief   string `json:"salvation_brief"`
	GodsWorkings     string `json:"gods_workings"`
	GodsCall         string `json:"gods_call"`
	IntoMinistry     bool   `json:"into_ministry"`
	SpiritualGifts   string `json:"spiritual_gifts"`
	Reason           string `json:"reason"`
	ChurchName       string `json:"church_name"`
	ChurchAddress    string `json:"church_address"`
	PastorName       string `json:"pastor_name"`
	PastorEmail      string `json:"pastor_email"`
	PastorPhone      string `json:"pastor_phone"`
	ChurchInvolved   string `json:"church_involved"`
	WaterBaptism     bool   `json:"water_baptism"`
	BaptismDate      string `json:"baptism_date"`
	HolyGhostBaptism bool   `json:"holyghost_baptism"`
}

// backgroundList defines array of background objects
type BackgroundList []*Background

/**
CRUD functions
*/

// Create creates a new background record
func (m *Background) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches background by id
func (m *Background) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Background) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &m); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given background
func (m *Background) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes background by id
func (m *Background) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Background) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
