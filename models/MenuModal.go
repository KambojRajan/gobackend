package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `json:"name" validate:"required"`
	Category     string             `json:"category" validate:"required,eq=SEA_FOOD|eq=VEG|eq=NON_VEG"`
	StartingDate *time.Time         `json:"starting_date" validate:"required"`
	EndDate      *time.Time         `json:"end_date"`
	CreateAt     time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	MenuId       string             `json:"food_id"`
}
