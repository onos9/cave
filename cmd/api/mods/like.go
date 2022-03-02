package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	likeTableName = "likes"
)

// Like is a model for Likes table
type Like struct {
	utils.Base
	IsLiked bool  `json:"isLiked"`
	Video   Video `gorm:"foreignkey:VideoID" json:"video"`
	User    User  `gorm:"foreignkey:UserID"  json:"user"`
}

// TableName gorm standard table name
func (c *Like) TableName() string {
	return likeTableName
}

// LikeList defines array of like objects
type LikeList []*Like

// TableName gorm standard table name
func (c *LikeList) TableName() string {
	return likeTableName
}

/**
* Relationship functions
 */

// GetCertificates returns like certificates
func (c *Like) GetVideo() error {
	return handler.Model(c).Related(&c.Video).Error
}

func (c *Like) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new like record
func (c *Like) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Like by id
func (c *Like) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Likes
func (c *Like) FetchAll(cl *LikeList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given like
func (c *Like) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes like by id
func (c *Like) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Like) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
