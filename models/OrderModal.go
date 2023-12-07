package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OrderDate time.Time          `json:"order_date" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	OrderId   string             `json:"ored_id"`
	TableId   *string            `json:"table_id" validate:"required"`
}
