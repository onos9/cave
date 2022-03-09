package mods

import (
	"time"
)

var (
	refreeTableName = "refrees"
)

// Refree is a model for Refrees table
type Refree struct {
	ID       string `json:"id"`
	Date     *time.Time
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	// TableName gorm standard table name
	// func (c *Refree) TableName() string {
	// 	return refreeTableName
}

// RefreeList defines array of refree objects
type RefreeList []*Refree

// TableName gorm standard table name
func (c *RefreeList) TableName() string {
	return refreeTableName
}

/**
* Relationship functions
 */

// GetCertificates returns refree certificates
// func (c *Refree) GetChannel() error {
// 	return handler.Model(c).Related(&c.Refree).Error
// }

// func (c *Refree) GetUser() error {
// 	return handler.Model(c).Related(&c.User).Error
// }

/**
CRUD functions
*/

// Create creates a new refree record
func (c *Refree) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Refree by id
func (c *Refree) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Refrees
func (c *Refree) FetchAll(cl *RefreeList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given refree
func (c *Refree) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes refree by id
func (c *Refree) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Refree) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
