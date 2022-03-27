package mods

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Video is a model for Videos table
type Video struct {
	utils.Base
	Date             *time.Time `json:"Date"`
	Title            string     `gorm:"type:varchar(100)" json:"title"`
	ChannelThumbnail string     `json:"channel_thumbnail"`
	ChannelTitle     string     `json:"channel_title"`
	Like             bool       `json:"IsLiked"`
	Dislike          bool       `json:"DisLiked"`
	Comment          string     `json:"comment"`
	Thumbnail        string     `json:"Thumbnail"`
	VideoID          string     `gorm:"type:varchar(100)" json:"video_id"`
	Description      string     `gorm:"type:varchar(100)" json:"description"`
	Channel          Channel    `gorm:"foreignkey:ChannelID" json:"channel"`
	Category         Category   `gorm:"foreignkey:CategoryID" json:"category"`
	View             bool       `gorm:"foreignkey:ViewID" json:"view"`
	User             User       `gorm:"foreignkey:UserID" json:"video"`
}

// VideoList defines array of video objects
type VideoList []*Video

/**
CRUD functions
*/

// Create creates a new video record
func (m *Video) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Video by id
func (m *Video) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Video) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given video
func (m *Video) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes video by id
func (m *Video) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Video) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
