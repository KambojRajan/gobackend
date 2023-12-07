package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	NumberOfGuest *int               `json:"number_of_guest" validate:"required,gt=0"`
	TableNumber   *int               `json:"table_number" validate:"required"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	TableId       string             `json:"table_id"`
}
