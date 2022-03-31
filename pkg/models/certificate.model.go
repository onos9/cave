package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	certificateCol = "certificates"
)

// Certificate defines a model for student certificates in a model
type Certificate struct {
	utils.Base
	CourseID       string `bson:"courseId,omitempty" json:"courseId"`
	CourseAuthorID string `bson:"courseAuthorId,omitempty" json:"courseAuthorId"`
	QRCorner       string `bson:"qrCorner,omitempty" json:"qrCorner"`
	QRScale        int    `bson:"qrScale,omitempty" json:"qrScale"`
	Margin         int    `bson:"margin,omitempty" json:"margin"`
	NameOffsetTop  int    `bson:"nameOffsetTop,omitempty" json:"nameOffsetTop"`
	FontSize       int    `bson:"fontSize,omitempty" json:"fontSize"`
	Message        string `bson:"message,omitempty" json:"message"`
	Active         bool   `bson:"active,omitempty" json:"active"`

	Course             Course                `bson:"course,omitempty" json:"course"`
	Issuer             CourseAuthor          `bson:"courseAuthor,omitempty" json:"courseAuthor"`
	IssuedCertificates IssuedCertificateList `bson:"certificates,omitempty" json:"certificates"`
}

// CertificateList defines array of certificate objects
type CertificateList []*Certificate

/**
CRUD functions
*/

// Create creates a new certificate record
func (m *Certificate) Create() error {
	_, err := db.Collection(certificateCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches Certificate by id
func (m *Certificate) FetchByID() error {
	err := db.Collection(certificateCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *Certificate) FetchAll(cl *CertificateList) error {
	cursor, err := db.Collection(certificateCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given certificate
func (m *Certificate) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(certificateCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes certificate by id
func (m *Certificate) Delete() error {
	_, err := db.Collection(certificateCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *Certificate) DeleteMany() error {
	_, err := db.Collection(certificateCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
