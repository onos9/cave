package mods

import "github.com/cave/pkg/utils"

var (
	backgroundTableName = "backgrounds"
)

// Background is a model for Backgrounds table
type Background struct {
	utils.Base
	BornAgain        bool   `json:"born_again"`
	SalvationBrief   string `json:"salvation_brief"`
	GodsWorkings     string `json:"gods_workings"`
	GodsCall         string `json:"gods_call"`
	IntoMinistry     bool   `json:"into_ministry"`
	SpiritualGifts   string `json:"spiritual_gifts"`
	Reason           string `json:"reason"`
	ChurchName       string `json:"church_name"`
	ChurchAddress    string `json:"church_address"`
	PastorName       string `json:"pastor_name"`
	PastorEmail      string `json:"pastor_email"`
	PastorPhone      string `json:"pastor_phone"`
	ChurchInvolved   string `json:"church_involved"`
	WaterBaptism     bool   `json:"water_baptism"`
	BaptismDate      string `json:"baptism_date"`
	HolyGhostBaptism bool   `json:"holyghost_baptism"`

	// TableName gorm standard table name
	// func (c *Background) TableName() string {
	// 	return backgroundTableName
}

// BackgroundList defines array of background objects
type BackgroundList []*Background

// TableName gorm standard table name
func (c *BackgroundList) TableName() string {
	return backgroundTableName
}

/**
* Relationship functions
 */

// GetCertificates returns background certificates
// func (c *Background) GetChannel() error {
// 	return handler.Model(c).Related(&c.Background).Error
// }

// func (c *Background) GetUser() error {
// 	return handler.Model(c).Related(&c.User).Error
// }

/**
CRUD functions
*/

// Create creates a new background record
func (c *Background) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Background by id
func (c *Background) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Backgrounds
func (c *Background) FetchAll(cl *BackgroundList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given background
func (c *Background) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes background by id
func (c *Background) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Background) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
