package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	commentTableName = "comments"
)

// Comment is a model for Comments table
type Comment struct {
	utils.Base
	Described string `json:"described"`
	Video     Video  `gorm:"foreignkey:VideoID" json:"video"`
	User      User   `gorm:"foreignkey:UserID" json:"user"`
}

// TableName gorm standard table name
func (c *Comment) TableName() string {
	return commentTableName
}

// CommentList defines array of comment objects
type CommentList []*Comment

// TableName gorm standard table name
func (c *CommentList) TableName() string {
	return commentTableName
}

/**
* Relationship functions
 */

// GetCertificates returns comment certificates
func (c *Comment) GetVideo() error {
	return handler.Model(c).Related(&c.Video).Error
}

func (c *Comment) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new comment record
func (c *Comment) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Comment by id
func (c *Comment) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Comments
func (c *Comment) FetchAll(cl *CommentList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given comment
func (c *Comment) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes comment by id
func (c *Comment) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Comment) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
