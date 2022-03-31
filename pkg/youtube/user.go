package youtube

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCol = "users"

// User is a model for Users table
type User struct {
	utils.Base
	Email                  string     `bson:"email" json:"email" `
	Username               string     `json:"username"`
	Password               string     `bson:"-" json:"omitemty"`
	PasswordSalt           string     `json:"passwordsalt"`
	PasswordHash           []byte     `json:"passwordhash"`
	IsVerified             bool       `json:"is_verified"`
	Role                   string     `jason:"role"`
	SiginInCount           int        `json:"signInCount"`
	CurrentSignInAt        *time.Time `json:"currentSignInAt"`
	LastSignInAt           *time.Time `json:"lastSignInAt"`
	CurrentSignInIP        string     `json:"currentSignInAp"`
	LastSignInIP           string     `json:"lastSignInIp"`
	RememberToken          string     `json:"rememberToken"`
	ConfirmedAt            *time.Time `json:"confirmedAt"`
	ConfirmationMailSentAt *time.Time `json:"confirmationMailSentAt"`

	Phone     string `json:"phone"`
	Title     string `json:"title"`
	KeySkills string `json:"keySkills"`
	About     string `bson:"about" json:"about"`

	Bio           Bio           `bson:"bio" json:"bio"`
	Qualification Qualification `bson:"bio" json:"qualification"`
	Background    Background    `bson:"bio" json:"background"`
	Health        Health        `bson:"bio" json:"health"`
	Referee       Referee       `bson:"bio" json:"referee"`
	Terms         Terms         `bson:"bio" json:"terms"`

	TimeZone *time.Time `json:"timezone"`
}

// UserList defines array of user objects
type UserList []*User

/**
CRUD functions
*/

// Create creates a new candidate record
func (m *User) Create() error {
	m.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	m.CreatedAt = time.Now()
	_, err := db.Collection(userCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches User by id
func (m *User) FetchByID() error {
	err := db.Collection(userCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches User by email
func (m *User) FetchByEmail() error {
	err := db.Collection(userCol).FindOne(context.TODO(), bson.M{"email": m.Email}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Users
func (m *User) FetchAll(cl *UserList) error {
	cursor, err := db.Collection(userCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given candidate
func (m *User) UpdateOne() error {
	m.UpdatedAt = time.Now()
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(userCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes candidate by id
func (m *User) Delete() error {
	m.DeletedAt = time.Now()
	_, err := db.Collection(userCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *User) DeleteMany() error {
	_, err := db.Collection(userCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
