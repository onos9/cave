package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	categoryTableName = "categorys"
)

// Category is a model for Categorys table
type Category struct {
	utils.Base
	Title    string `gorm:"type:varchar(100)" json:"title"`
	Semester int    `gorm:"type:varchar(100)" json:"semester"`
}

// TableName gorm standard table name
func (c *Category) TableName() string {
	return categoryTableName
}

// CategoryList defines array of category objects
type CategoryList []*Category

// TableName gorm standard table name
func (c *CategoryList) TableName() string {
	return categoryTableName
}

/**
* Relationship functions
 */

// GetCertificates returns category certificates
// func (c *Category) GetCertificates() error {
// 	return handler.Model(c).Related(&c.Certificates).Error
// }

/**
CRUD functions
*/

// Create creates a new category record
func (c *Category) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Category by id
func (c *Category) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Categorys
func (c *Category) FetchAll(cl *CategoryList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given category
func (c *Category) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes category by id
func (c *Category) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Category) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
