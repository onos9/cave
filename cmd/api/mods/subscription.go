package mods

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Subscription is a model for Subscriptions table
type Subscription struct {
	utils.Base
	IsSubscribed   bool    `json:"isSubscribed"`
	Thumbnail      string  `json:"thumbnail"`
	Title          string  `json:"title"`
	ChannelCreated string  `json:"channel_created"`
	Channel        Channel `gorm:"foreignkey:UserID" json:"channel"`
	User           User    `gorm:"foreignkey:UserID" json:"user"`
}

// SubscriptionList defines array of subscription objects
type SubscriptionList []*Subscription

/**
CRUD functions
*/

// Create creates a new subscription record
func (m *Subscription) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Subscription by id
func (m *Subscription) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Subscription) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given subscription
func (m *Subscription) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes subscription by id
func (m *Subscription) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Subscription) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
