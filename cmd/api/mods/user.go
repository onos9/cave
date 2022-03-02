package mods

import (
	"time"

	"github.com/cave/pkg/utils"
)

var (
	userTableName = "users"
)

// User is a model for Users table
type User struct {
	utils.Base
	Email                  string `gorm:"type:varchar(100);unique_index" `
	Password               string `gorm:"migration"`
	PasswordSalt           string
	PasswordHash           []byte
	Role                   int
	SiginInCount           int
	CurrentSignInAt        *time.Time
	LastSignInAt           *time.Time
	CurrentSignInIP        string
	LastSignInIP           string
	RememberToken          string
	ConfirmedAt            *time.Time
	ConfirmationMailSentAt *time.Time
	Name                   string
	Phone                  string
	Title                  string
	KeySkills              string
	About                  string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}

// TableName gorm standard table name
func (c *User) TableName() string {
	return userTableName
}

// UserList defines array of user objects
type UserList []*User

// TableName gorm standard table name
func (c *UserList) TableName() string {
	return userTableName
}

/**
* Relationship functions
 */

// GetCertificates returns user certificates
// func (c *User) GetCertificates() error {
// 	return handler.Model(c).Related(&c.Certificates).Error
// }

/**
CRUD functions
*/

// Create creates a new user record
func (c *User) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches User by id
func (c *User) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Users
func (c *User) FetchAll(cl *UserList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given user
func (c *User) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes user by id
func (c *User) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *User) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
