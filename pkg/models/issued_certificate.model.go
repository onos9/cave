package models

import (
	"context"

	"github.com/cave/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	issuedCertificateCol = "issued_certificates"
)

// IssuedCertificate defines a model for IssuedCertificate certificates for a course
type IssuedCertificate struct {
	utils.Base
	CertificateID       string      `bson:"certificateId,omitempty" json:"certificateId"`
	IssuedCertificateID string      `bson:"userId,omitempty" json:"userId"`
	SerialNumber        string      `bson:"serialNumber,omitempty" json:"serialNumber"`
	Certificate         Certificate `bson:"certificate,omitempty" json:"certificate"`
}

// IssuedCertificateList defines array of certificate objects
type IssuedCertificateList []*IssuedCertificate

/**
CRUD functions
*/

// Create creates a new IssuedCertificate record
func (m *IssuedCertificate) Create() error {
	_, err := db.Collection(issuedCertificateCol).InsertOne(context.TODO(), &m)
	if err != nil {
		return err
	}
	return nil
}

// FetchByID fetches IssuedCertificate by id
func (m *IssuedCertificate) FetchByID() error {
	err := db.Collection(issuedCertificateCol).FindOne(context.TODO(), bson.M{"_id": m.ID}).Decode(&m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all Candidates
func (m *IssuedCertificate) FetchAll(cl *IssuedCertificateList) error {
	cursor, err := db.Collection(issuedCertificateCol).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &cl); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given IssuedCertificate
func (m *IssuedCertificate) UpdateOne() error {
	update := bson.M{
		"$inc": bson.M{"copies": 1},
	}
	_, err := db.Collection(issuedCertificateCol).UpdateOne(context.TODO(), bson.M{"_id": m.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes IssuedCertificate by id
func (m *IssuedCertificate) Delete() error {
	_, err := db.Collection(issuedCertificateCol).DeleteOne(context.TODO(), bson.M{"_id": m.ID})
	if err != nil {
		return err
	}
	return nil
}

func (m *IssuedCertificate) DeleteMany() error {
	_, err := db.Collection(issuedCertificateCol).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
