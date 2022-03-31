package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Base serves as a base model for other models
type Base struct {
	ID        primitive.ObjectID ` json:"id" bson:"_id" validate:"omitempty,uuid,required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"update_at"`
	DeletedAt time.Time          `json:"-" bson:"deleted_at,omitempty"`
	Doc       string             `bson:"-"`
}

// GetID returns Id of the model
func (base *Base) GetID() primitive.ObjectID {
	return base.ID
}

// SetID sets Id of the model
func (base *Base) SetID(id primitive.ObjectID) {
	base.ID = id
}

// SetCreatedAt sets field createdAt, should only be used in mongodb
func (base *Base) SetCreatedAt(t time.Time) {
	base.CreatedAt = t
}

// SetUpdatedAt sets field UpdatedAt
func (base *Base) SetUpdatedAt(t time.Time) {
	base.UpdatedAt = t
}

// SetDeletedAt sets field DeletedAt
func (base *Base) SetDeletedAt(t time.Time) {
	base.DeletedAt = t
}
