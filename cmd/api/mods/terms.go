package mods

import (
	"time"
)

var (
	termsTableName = "termss"
)

// Terms is a model for Termss table
type Terms struct {
	ID             string `json:"id"`
	Date           *time.Time
	Scholarship bool `json:"scholarship"`
    ScholarshipReason string `json:"scholReason"`
  Agree bool `json:"agree"`
	About          string `gorm:"type:text" json:"about" validate:"omitempty"`
	

// TableName gorm standard table name
func (c *Terms) TableName() string {
	return termsTableName
}

// TermsList defines array of terms objects
type TermsList []*Terms

// TableName gorm standard table name
func (c *TermsList) TableName() string {
	return termsTableName
}

/**
* Relationship functions
 */

// GetCertificates returns terms certificates
// func (c *Terms) GetChannel() error {
// 	return handler.Model(c).Related(&c.Terms).Error
// }

func (c *Terms) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new terms record
func (c *Terms) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Terms by id
func (c *Terms) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Termss
func (c *Terms) FetchAll(cl *TermsList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given terms
func (c *Terms) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes terms by id
func (c *Terms) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Terms) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
