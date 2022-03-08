package mods

import "time"

var (
	candidateTableName = "candidates"
)

// Candidate is a model for Candidates table
type Candidate struct {
	ID            string `json:"id"`
	Date          *time.Time
	Bio           string `json:"bio"`
	Qualification string `json:"qualification"`
	Background    string `json:"background"`
	Health        string `json:"health"`
	Referee       string `json:"referee"`
	Terms         bool   `json:"terms"`
	User          User   `gorm:"foreignkey:UserID" json:"user"`
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
func (c *Candidate) GetChannel() error {
	return handler.Model(c).Related(&c.Bio).Error
}

func (c *Candidate) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
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
