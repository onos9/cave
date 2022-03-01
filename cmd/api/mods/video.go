package mods

import (
	"time"

	"github.com/cave/pkg/utils"
)

var (
	videoTableName = "videos"
)

// Video is a model for Videos table
type Video struct {
	utils.Base
	Date        *time.Time
	Title       bool
	Thumnail    string
	VideoID     string   `gorm:"type:varchar(100)"`
	Ip          string   `gorm:"type:varchar(20)"`
	Description string   `gorm:"type:varchar(100)"`
	Channel     Channel  `gorm:"foreignkey:ChannelID"`
	Category    Category `gorm:"foreignkey:CategoryID"`
}

// TableName gorm standard table name
func (c *Video) TableName() string {
	return videoTableName
}

// VideoList defines array of video objects
type VideoList []*Video

// TableName gorm standard table name
func (c *VideoList) TableName() string {
	return videoTableName
}

/**
* Relationship functions
 */

// GetCertificates returns video certificates
func (c *Video) GetChannel() error {
	return handler.Model(c).Related(&c.Channel).Error
}

func (c *Video) GetCategory() error {
	return handler.Model(c).Related(&c.Category).Error
}

/**
CRUD functions
*/

// Create creates a new video record
func (c *Video) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Video by id
func (c *Video) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Videos
func (c *Video) FetchAll(cl *VideoList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given video
func (c *Video) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes video by id
func (c *Video) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Video) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
