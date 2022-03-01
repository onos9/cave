package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	viewTableName = "views"
)

// View is a model for Views table
type View struct {
	utils.Base
	Name   string `gorm:"type:varchar(100);unique" `
	Ip     string `gorm:"type:varchar(20)"`
	IsView bool
	Video  Video `gorm:"foreignkey:VideoID"`
}

// TableName gorm standard table name
func (c *View) TableName() string {
	return viewTableName
}

// ViewList defines array of view objects
type ViewList []*View

// TableName gorm standard table name
func (c *ViewList) TableName() string {
	return viewTableName
}

/**
* Relationship functions
 */

// GetCertificates returns view certificates
func (c *View) GetCertificates() error {
	return handler.Model(c).Related(&c.Video).Error
}

/**
CRUD functions
*/

// Create creates a new view record
func (c *View) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches View by id
func (c *View) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Views
func (c *View) FetchAll(cl *ViewList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given view
func (c *View) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes view by id
func (c *View) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *View) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
