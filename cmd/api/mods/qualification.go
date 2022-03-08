package mods

import (
	"time"
)

var (
	qualificationTableName = "qualifications"
)

// Qualification is a model for Qualifications table
type Qualification struct {
	ID             string `json:"id"`
	Date           *time.Time
	Degree         string `gorm:"type:varchar(100);unique_index" json:"degree" `
	Instution      string `json:"university"`
	GraduationYear string `json:"graduationYear"`
	About          string `gorm:"type:text" json:"about" validate:"omitempty"`

	// TableName gorm standard table name
	// func (c *Qualification) TableName() string {
	// 	return qualificationTableName
}

// QualificationList defines array of qualification objects
type QualificationList []*Qualification

// TableName gorm standard table name
func (c *QualificationList) TableName() string {
	return qualificationTableName
}

/**
* Relationship functions
 */

// GetCertificates returns qualification certificates
// func (c *Qualification) GetChannel() error {
// 	return handler.Model(c).Related(&c.Qualification).Error
// }

// func (c *Qualification) GetUser() error {
// 	return handler.Model(c).Related(&c.User).Error
// }

/**
CRUD functions
*/

// Create creates a new qualification record
func (c *Qualification) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Qualification by id
func (c *Qualification) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Qualifications
func (c *Qualification) FetchAll(cl *QualificationList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given qualification
func (c *Qualification) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes qualification by id
func (c *Qualification) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Qualification) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
