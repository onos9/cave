package mods

import "github.com/cave/pkg/utils"

var (
	bioTableName = "bios"
)

// Bio is a model for Bios table
type Bio struct {
	utils.Base
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	Dob         string `json:"dob"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zip         string `json:"zipcode"`
	PhoneNo     string `json:"phone"`
	Nationality string `json:"nationality"`
	Profession  string `json:"profession"`
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
// func (c *Bio) GetChannel() error {
// 	return handler.Model(c).Related(&c.Bio).Error
// }

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
