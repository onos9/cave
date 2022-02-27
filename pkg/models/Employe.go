package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// perbedaan menggunakan omitempty dan tidak jika menggunakan 
// omitempty ketika post json kosong tidak terinsert ke document
// jika menggunkana emitpty nanti data akan duplikat juga pas insert

type Employe struct {
	ID     primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string   `json:"name" valid:"required~Name is blank" bson:"name"`
	Salary float64 `json:"salary" bson:"salary" valid:"required~Salary is blank"`
	Age    int64 `json:"age" bson:"age" valid:"required"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt"`
}