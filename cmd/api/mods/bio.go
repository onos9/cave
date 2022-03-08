package mods

import "time"

var (
	bioTableName = "bios"
)

// Bio is a model for Bios table
type Bio struct {
	ID            string `json:"id"`
	Date          *time.Time
	Bio           string `json:"bio"`
	Email         string `gorm:"type:varchar(100);unique_index" json:"email" `
	Name          string `json:"name"`
	Age           string `json:"age"`
	Dob           string `json:"dob"`
	Gender        string `json:"gender"`
	PhoneNumber   string `json:"phoneNumber"`
	Username      string `json:"username"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Title         string `json:"title"`
	MaritalStatus string `json:"maritalStatus"`
	Religion      string `json:"religion"`
	KeySkills     string `json:"keySkills"`
	About         string `gorm:"type:text" json:"about" validate:"omitempty"`
	User          User   `gorm:"foreignkey:UserID" json:"user"`
}

// TableName gorm standard table name
func (c *Bio) TableName() string {
	return bioTableName
}

// BioList defines array of bio objects
type BioList []*Bio

// TableName gorm standard table name
func (c *BioList) TableName() string {
	return bioTableName
}

/**
* Relationship functions
 */

// GetCertificates returns bio certificates
func (c *Bio) GetChannel() error {
	return handler.Model(c).Related(&c.Bio).Error
}

func (c *Bio) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new bio record
func (c *Bio) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Bio by id
func (c *Bio) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Bios
func (c *Bio) FetchAll(cl *BioList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given bio
func (c *Bio) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes bio by id
func (c *Bio) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Bio) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
