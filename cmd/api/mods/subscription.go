package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	subscriptionTableName = "subscriptions"
)

// Subscription is a model for Subscriptions table
type Subscription struct {
	utils.Base
	IsSubscribed   bool    `json:"isSubscribed"`
	Thumbnail      string  `json:"thumbnail"`
	Title          string  `json:"title"`
	ChannelCreated string  `json:"channel_created"`
	Channel        Channel `gorm:"foreignkey:UserID" json:"channel"`
	User           User    `gorm:"foreignkey:UserID" json:"user"`
}

// TableName gorm standard table name
func (c *Subscription) TableName() string {
	return subscriptionTableName
}

// SubscriptionList defines array of subscription objects
type SubscriptionList []*Subscription

// TableName gorm standard table name
func (c *SubscriptionList) TableName() string {
	return subscriptionTableName
}

/**
* Relationship functions
 */

// GetCertificates returns subscription certificates
func (c *Subscription) GetChannel() error {
	return handler.Model(c).Related(&c.Channel).Error
}

func (c *Subscription) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new subscription record
func (c *Subscription) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Subscription by id
func (c *Subscription) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Subscriptions
func (c *Subscription) FetchAll(cl *SubscriptionList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given subscription
func (c *Subscription) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes subscription by id
func (c *Subscription) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Subscription) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
