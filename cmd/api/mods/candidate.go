package mods

import (
	"github.com/cave/pkg/utils"
)

var (
	candidateTableName = "candidates"
)

// Candidate is a model for Candidates table
type Candidate struct {
	utils.Base
	Email         string        `gorm:"type:varchar(100);unique_index" json:"email" `
	PasswordHash  []byte        `json:"password_hash"`
	PasswordSalt  string        `json:"password_salt"`
	IsVerified    bool          `json:"is_verified"`
	Bio           Bio           `gorm:"foreignkey:BioID" json:"bio"`
	Qualification Qualification `gorm:"foreignkey:QualificationID" json:"qualification"`
	Background    Background    `gorm:"foreignkey:BackgroundID" json:"background"`
	Health        Health        `gorm:"foreignkey:HealthID" json:"health"`
	Referee       Ref           `gorm:"foreignkey:RefereeID" json:"referee"`
	Terms         Terms         `gorm:"foreignkey:TermsID" json:"terms"`
}

// TableName gorm standard table name
func (c *Candidate) TableName() string {
	return candidateTableName
}

// CandidateList defines array of candidate objects
type CandidateList []*Candidate

// TableName gorm standard table name
func (c *CandidateList) TableName() string {
	return candidateTableName
}

/**
* Relationship functions
 */

// GetCertificates returns candidate certificates
func (c *Candidate) GetBio() error {
	return handler.Model(c).Related(&c.Bio).Error
}

func (c *Candidate) GetQualification() error {
	return handler.Model(c).Related(&c.Qualification).Error
}

func (c *Candidate) GetBackground() error {
	return handler.Model(c).Related(&c.Background).Error
}

func (c *Candidate) GetHealth() error {
	return handler.Model(c).Related(&c.Health).Error
}

func (c *Candidate) GetTerms() error {
	return handler.Model(c).Related(&c.Terms).Error
}

/**
CRUD functions
*/

// Create creates a new candidate record
func (c *Candidate) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Candidate by id
func (c *Candidate) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchByID fetches User by Email
func (c *Candidate) FetchByEmail() error {
	err := handler.Where("email=?", c.Email).First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Candidates
func (c *Candidate) FetchAll(cl *CandidateList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given candidate
func (c *Candidate) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// UpdateOne updates a given candidate or creates a new one if it doesn't exist'
func (c *Candidate) UpdateOrCreateByEmail() error {
	err := handler.Where("email=?", c.Email).FirstOrCreate(c).Error
	return err
}

// Delete deletes candidate by id
func (c *Candidate) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Candidate) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
