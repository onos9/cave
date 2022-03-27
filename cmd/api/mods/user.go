package mods

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// User is a model for Users table
type User struct {
	utils.Base
	Email                  string     `gorm:"type:varchar(100);unique_index" json:"email" `
	Password               string     `gorm:"migration" json:"password"`
	PasswordSalt           string     `json:"passwordsalt"`
	PasswordHash           []byte     `json:"passwordhash"`
	ExternalID             string     `jason:"external_id"`
	Role                   string     `jason:"role"`
	SiginInCount           int        `json:"signInCount"`
	CurrentSignInAt        *time.Time `json:"currentSignInAt"`
	LastSignInAt           *time.Time `json:"lastSignInAt"`
	CurrentSignInIP        string     `json:"currentSignInAp"`
	LastSignInIP           string     `json:"lastSignInIp"`
	RememberToken          string     `json:"rememberToken"`
	ConfirmedAt            *time.Time `json:"confirmedAt"`
	ConfirmationMailSentAt *time.Time `json:"confirmationMailSentAt"`
	Name                   string     `json:"name"`
	Username               string     `json:"username"`
	Phone                  string     `json:"phone"`
	Title                  string     `json:"title"`
	KeySkills              string     `json:"keySkills"`
	About                  string     `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}

// UserList defines array of user objects
type UserList []*User

/**
CRUD functions
*/

// Create creates a new user record
func (m *User) Create() error {
	_, err := db.Collection(m.Doc).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches User by id
func (m *User) FetchByID() error {
	err := db.Collection(m.Doc).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *User) FetchAll(cl *CandidateList) error {
	cursor, err := db.Collection(m.Doc).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (m *User) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(m.Doc).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *User) Delete() error {
	_, err := db.Collection(m.Doc).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *User) DeleteMany() error {
	_, err := db.Collection(m.Doc).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
