package mods

import "github.com/cave/pkg/utils"

var (
	refTableName = "refs"
)

// Ref is a model for Refs table
type Ref struct {
	utils.Base
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	// TableName gorm standard table name
	// func (c *Ref) TableName() string {
	// 	return refTableName
}

// RefList defines array of refree objects
type RefList []*Ref

// TableName gorm standard table name
func (c *RefList) TableName() string {
	return refTableName
}

/**
* Relationship functions
 */

// GetCertificates returns refree certificates
// func (c *Ref) GetChannel() error {
// 	return handler.Model(c).Related(&c.Ref).Error
// }

// func (c *Ref) GetUser() error {
// 	return handler.Model(c).Related(&c.User).Error
// }

/**
CRUD functions
*/

// Create creates a new refree record
func (c *Ref) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Ref by id
func (c *Ref) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Refs
func (c *Ref) FetchAll(cl *RefList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given refree
func (c *Ref) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes refree by id
func (c *Ref) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Ref) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
