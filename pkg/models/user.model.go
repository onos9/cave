package models

import (
	"context"
	"time"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Database table for User
	userCol = "users"
)

// User struct for users table
type User struct {
	utils.Base
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	*BioData       `bson:"bioData,omitempty" json:"bioData"`
	*Qualification `bson:"qualification,omitempty" json:"qualification"`
	*Background    `bson:"background,omitempty" json:"background"`
	*Health        `bson:"health,omitempty" json:"health"`
	*Terms         `bson:"terms,omitempty" json:"terms"`

	RefereeList            []*Referee `bson:"referees,omitempty" json:"referees"`
	Email                  string     `bson:"email,omitempty" json:"email,omitempty"`
	Password               string     `bson:"-" json:"password,omitempty"`
	IsVerified             bool       `bson:"isVerified" json:"isVerified,omitempty"`
	PasswordSalt           string     `bson:"passwordsalt,omitempty" json:"-"`
	PasswordHash           []byte     `bson:"passwordhash,omitempty" json:"-"`
	ExternalID             string     `bson:"external_id,omitempty" json:"externalID,omitempty"`
	Role                   string     `bson:"role,omitempty" json:"role,omitempty"`
	SignInCount            int        `bson:"signInCount,omitempty" json:"signInCount,omitempty"`
	CurrentSignInAt        *time.Time `bson:"currentSignInAt,omitempty" json:"currentSignInAt,omitempty"`
	LastSignInAt           *time.Time `bson:"lastSignInAt,omitempty" json:"lastSignInAt,omitempty"`
	CurrentSignInIP        string     `bson:"currentSignInIp,omitempty" json:"currentSignInIp,omitempty"`
	LastSignInIP           string     `bson:"lastSignInIp,omitempty" json:"lastSignInIp,omitempty"`
	RememberToken          string     `bson:"rememberToken,omitempty" json:"rememberToken,omitempty"`
	ConfirmedAt            *time.Time `bson:"confirmedAt,omitempty" json:"confirmedAt,omitempty"`
	ConfirmationMailSentAt *time.Time `bson:"confirmationMailSentAt,omitempty" json:"confirmationMailSentAt,omitempty"`
	FullName               string     `bson:"fullName,omitempty" json:"fullName,omitempty"`
	Username               string     `bson:"username,omitempty" json:"username,omitempty"`
	Phone                  string     `bson:"phone,omitempty" json:"phone,omitempty"`
	Title                  string     `bson:"title,omitempty" json:"title,omitempty"`
	KeySkills              string     `bson:"keySkills,omitempty" json:"keySkills,omitempty"`
	About                  string     `bson:"about,omitempty" json:"about,omitempty"`
	EnrollProgress         int        `bson:"enrollProgress,omitempty" json:"enrollProgress"`
	Status                 string     `bson:"status,omitempty" json:"status"`
	Program                string     `bson:"program,omitempty" json:"program"`
	ProgramOption          string     `bson:"programOption,omitempty" json:"programOption"`

	TimeZone        *time.Time        `bson:"timeZone,omitempty" json:"timeZone,omitempty"`
	AuthoredCourses CourseAuthorList  `bson:"authoredCourses,omitempty" json:"authoredCourses,omitempty"`
	Courses         StudentCourseList `bson:"courses,omitempty" json:"courses,omitempty"`
}

// Bio is a model for Bios table
type BioData struct {
	FirstName   string `bson:"firstName,omitempty" json:"firstName"`
	LastName    string `bson:"lastName,omitempty" json:"lastName"`
	MiddleName  string `bson:"middleName,omitempty" json:"middleName"`
	Dob         string `bson:"dob,omitempty" json:"dob"`
	Gender      string `bson:"gender,omitempty" json:"gender"`
	Address     string `bson:"address,omitempty" json:"address"`
	City        string `bson:"city,omitempty" json:"city"`
	State       string `bson:"state,omitempty" json:"state"`
	Country     string `bson:"country,omitempty" json:"country"`
	Zip         string `bson:"zipCode,omitempty" json:"zipCode"`
	Phone       string `bson:"phone,omitempty" json:"phone"`
	Nationality string `bson:"nationality,omitempty" json:"nationality"`
	Profession  string `bson:"profession,omitempty" json:"profession"`
}

// Qualification is a model for Qualifications table
type Qualification struct {
	Degree         string `bson:"degree,omitempty" json:"degree" `
	Instution      string `bson:"institution,omitempty" json:"institution"`
	InstutionName  string `bson:"institutionName,omitempty" json:"institutionName"`
	GraduationYear string `bson:"graduationYear,omitempty" json:"graduationYear"`
}

// background is a model for backgrounds table
type Background struct {
	BornAgain        bool   `bson:"bornAgain,omitempty" json:"bornAgain,string,omitempty"`
	SalvationBrief   string `bson:"briefSalvation,omitempty" json:"briefSalvation,omitempty"`
	GodsWorkings     string `bson:"godsWorkings,omitempty" json:"godsWorkings"`
	GodsCall         string `bson:"godsCall,omitempty" json:"godsCall"`
	IntoMinistry     bool   `bson:"intoMinistry,omitempty" json:"intoMinistry,string"`
	SpiritualGifts   string `bson:"SpiritualGifts,omitempty" json:"spiritualGifts"`
	Reason           string `bson:"reason,omitempty" json:"reason"`
	ChurchName       string `bson:"churchName,omitempty" json:"churchName"`
	ChurchAddress    string `bson:"churchAdress,omitempty" json:"churchAddress"`
	PastorName       string `bson:"pastorName,omitempty" json:"pastorName"`
	PastorEmail      string `bson:"pastorEmail,omitempty" json:"pastorEmail"`
	PastorPhone      string `bson:"pastorPhone,omitempty" json:"pastorPhone"`
	ChurchInvolved   string `bson:"churchInvolve,omitempty" json:"churchInvolve"`
	WaterBaptism     bool   `bson:"waterBaptism,omitempty" json:"waterBaptism,string"`
	BaptismDate      string `bson:"baptismDate,omitempty" json:"baptismDate"`
	HolyGhostBaptism bool   `bson:"holyghostBaptism,omitempty" json:"holyghostBaptism,string"`
}

// Health is a model for Healths table
type Health struct {
	Disability             bool   `bson:"disability,omitempty" json:"disability,string,omitempty"`
	Nervousillness         bool   `bson:"nervous,omitempty" json:"nervous,string,omitempty"`
	Anorexia               bool   `bson:"anorexic,omitempty" json:"anorexic,string,omitempty"`
	Diabetese              bool   `bson:"diabetic,omitempty" json:"diabetic,string,omitempty"`
	Epilepsy               bool   `bson:"epileptic,omitempty" json:"epileptic,string,omitempty"`
	StomachUlcers          bool   `bson:"stomachUlcer,omitempty" json:"stomachUlcer,string,omitempty"`
	SpecialDiet            bool   `bson:"specilaDiet,omitempty" json:"specialDiet,string,omitempty"`
	LearningDisability     bool   `bson:"learningDisability,omitempty" json:"learningDisability,string,omitempty"`
	UsedIllDrug            bool   `bson:"illegalDrugs,omitempty" json:"illegalDrug,string,omitempty"`
	DrugAddiction          bool   `bson:"drugAddiction,omitempty" json:"drugAddiction,string,omitempty"`
	HadSurgery             bool   `bson:"surgery,omitempty" json:"surgery,string,omitempty"`
	HealthIssueDescription string `bson:"healthIssues,omitempty" json:"healthIssues"`
}

// Ref is a model for Refs table
type Referee struct {
	FullName string `bson:"fullName,omitempty" json:"fullName"`
	Email    string `bson:"email,omitempty" json:"email"`
	Phone    string `bson:"phone,omitempty" json:"phone"`
}

// Terms is a model for Termss table
type Terms struct {
	Scholarship       bool   `bson:"scholarship,omitempty" json:"scholarship"`
	ScholarshipReason string `bson:"scholReason,omitempty" json:"scholReason"`
	Agree             bool   `bson:"agree,omitempty" json:"agree"`
}

// UserList defines array of user objects
type UserList []*User

/**
CRUD functions
*/

// Create creates a new user record
func (m *User) Create() error {
	t := time.Now()
	m.CreatedAt = &t
	m.Id = primitive.NewObjectID()
	result, err := db.Collection(userCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	m.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

// FetchByID fetches User by id
func (m *User) FetchByID() error {
	err := db.Collection(userCol).FindOne(context.TODO(), bson.M{"_id": m.Id}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByEmail fetches User by email
func (m *User) FetchByEmail() error {
	err := db.Collection(userCol).FindOne(context.TODO(), bson.M{"email": m.Email}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all User
func (m *User) FetchAll(ul *UserList) error {
	cursor, err := db.Collection(userCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), ul); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (m *User) UpdateOne() error {
	t := time.Now()
	m.UpdatedAt = &t

	bm, err := bson.Marshal(m)
	if err != nil {
		return err
	}

	var val bson.M
	err = bson.Unmarshal(bm, &val)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": m.Id}
	update := bson.D{{Key: "$set", Value: val}}
	_, err = db.Collection(userCol).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (m *User) Delete() error {
	t := time.Now()
	m.DeletedAt = &t
	_, err := db.Collection(userCol).DeleteOne(context.TODO(), bson.M{"_id": m.Id})
	if err != nil {
		return err
	}
	return nil
}

func (m *User) DeleteMany() error {
	_, err := db.Collection(userCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
