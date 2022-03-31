package models

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// Database table for User
	userCol = "users"
)

// User struct for users table
type User struct {
	utils.Base
	Email                  string     `bson:"email,omitempty" json:"email"`
	Password               string     `bson:"-" json:"password"`
	PasswordSalt           string     `bson:"passwordsalt,omitempty" json:"passwordsalt"`
	PasswordHash           []byte     `bson:"passwordhash,omitempty" json:"passwordhash"`
	ExternalID             string     `bson:"external_id,omitempty" jason:"external_id"`
	Role                   string     `bson:"role,omitempty" jason:"role"`
	SignInCount            int        `bson:"signInCount,omitempty" json:"signInCount"`
	CurrentSignInAt        *time.Time `bson:"currentSignInAt,omitempty" json:"currentSignInAt"`
	LastSignInAt           *time.Time `bson:"lastSignInAt,omitempty" json:"lastSignInAt"`
	CurrentSignInIP        string     `bson:"currentSignInIp,omitempty" json:"currentSignInIp"`
	LastSignInIP           string     `bson:"lastSignInIp,omitempty" json:"lastSignInIp"`
	RememberToken          string     `bson:"rememberToken,omitempty" json:"rememberToken"`
	ConfirmedAt            *time.Time `bson:"confirmedAt,omitempty" json:"confirmedAt"`
	ConfirmationMailSentAt *time.Time `bson:"confirmationMailSentAt,omitempty" json:"confirmationMailSentAt"`
	Name                   string     `bson:"name,omitempty" json:"name"`
	Username               string     `bson:"username,omitempty" json:"username"`
	Phone                  string     `bson:"phone,omitempty" json:"phone"`
	Title                  string     `bson:"title,omitempty" json:"title"`
	KeySkills              string     `bson:"keySkills,omitempty" json:"keySkills"`
	About                  string     `bson:"about,omitempty" json:"about"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`

	AuthoredCourses CourseAuthorList  `gorm:"foreignkey:UserID"`
	Courses         StudentCourseList `gorm:"foreignkey:UserID"`
}

// UserList defines array of user objects
type UserList []*User

/**
CRUD functions
*/

// Create creates a new user record
func (m *User) Create() error {
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

// FetchAll fetchs all Candidates
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

// UpdateOne updates a given user
func (m *User) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(userCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *User) Delete() error {
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
