package mods

import (
	"time"
)

var (
	healthTableName = "healths"
)

// Health is a model for Healths table
type Health struct {
	ID                     string `json:"id"`
	Date                   *time.Time
	Disability             bool   `json:"disability"`
	Nervousillness         bool   `json:"nervousIll"`
	Anorexia               bool   `json:"anorexia"`
	Diabetese              bool   `json:"diabetese"`
	Epilepsy               bool   `json:"epilepsy"`
	StomachUlcers          bool   `json:"stomach_ulcers"`
	SpecialDiet            bool   `json:"special_diet"`
	LearningDisability     bool   `json:"learning_disability"`
	UsedIllDrug            bool   `json:"usedIllDrug"`
	DrugAddiction          bool   `json:"drug_addiction"`
	HadSurgery             bool   `json:"had_surgery"`
	HealthIssueDescription string `json:"healthIssueDesc"`

	// TableName gorm standard table name
	// func (c *Health) TableName() string {
	// 	return healthTableName
}

// HealthList defines array of health objects
type HealthList []*Health

// TableName gorm standard table name
func (c *HealthList) TableName() string {
	return healthTableName
}

/**
* Relationship functions
 */

// GetCertificates returns health certificates
// func (c *Health) GetChannel() error {
// 	return handler.Model(c).Related(&c.Health).Error
// }

// func (c *Health) GetUser() error {
// 	return handler.Model(c).Related(&c.User).Error
// }

/**
CRUD functions
*/

// Create creates a new health record
func (c *Health) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Health by id
func (c *Health) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Healths
func (c *Health) FetchAll(cl *HealthList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given health
func (c *Health) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes health by id
func (c *Health) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Health) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
