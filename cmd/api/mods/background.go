package mods

var (
	backgroundTableName = "backgrounds"
)

// Background is a model for Backgrounds table
type Background struct {
	ID            string `json:"id"`
	Date           *time.Time
	Background    string `json:"background"`
	BornAgainYear string `gorm:"type:varchar(100);unique_index" json:"bornAgainYear" `
	Baptism       string `json:"baptism"`
	Ministry      string `json:"ministry"`
	Role          string `json:"role"`
	About         string `gorm:"type:text" json:"about" validate:"omitempty"`
	

// TableName gorm standard table name
func (c *Background) TableName() string {
	return backgroundTableName
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
func (c *Background) GetChannel() error {
	return handler.Model(c).Related(&c.Background).Error
}

func (c *Background) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

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
