package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	dislikeTableName = "dislikes"
)

// Dislike is a model for Dislikes table
type Dislike struct {
	utils.Base
	IsDisliked bool
	Video      Video `gorm:"foreignkey:VideoID"`
	User       User  `gorm:"foreignkey:UserID"`
}

// TableName gorm standard table name
func (c *Dislike) TableName() string {
	return dislikeTableName
}

// DislikeList defines array of dislike objects
type DislikeList []*Dislike

// TableName gorm standard table name
func (c *DislikeList) TableName() string {
	return dislikeTableName
}

/**
* Relationship functions
 */

// GetCertificates returns dislike certificates
func (c *Dislike) GetVideo() error {
	return handler.Model(c).Related(&c.Video).Error
}

func (c *Dislike) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new dislike record
func (c *Dislike) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Dislike by id
func (c *Dislike) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Dislikes
func (c *Dislike) FetchAll(cl *DislikeList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given dislike
func (c *Dislike) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes dislike by id
func (c *Dislike) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Dislike) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
