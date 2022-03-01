package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	channelTableName = "channels"
)

// Channel is a model for Channels table
type Channel struct {
	utils.Base
	ID       string `gorm:"type:varchar(100);unique" `
	Name     string `gorm:"type:varchar(100);unique"`
	Thumnail string
	Banner   string
	About    string `gorm:"type:varchar(100)"`
	User     User   `gorm:"foreignkey:CourseID"`
}

// TableName gorm standard table name
func (c *Channel) TableName() string {
	return channelTableName
}

// ChannelList defines array of channel objects
type ChannelList []*Channel

// TableName gorm standard table name
func (c *ChannelList) TableName() string {
	return channelTableName
}

/**
* Relationship functions
 */

// GetCertificates returns channel certificates
func (c *Channel) GetCertificates() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new channel record
func (c *Channel) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Channel by id
func (c *Channel) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Channels
func (c *Channel) FetchAll(cl *ChannelList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given channel
func (c *Channel) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes channel by id
func (c *Channel) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Channel) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
