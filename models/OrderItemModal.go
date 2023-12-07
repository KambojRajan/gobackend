package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Quantity    *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	UnitPrice   *float64           `json:"unit_price" validate:"required"`
	CreatedAtt  time.Time          `josn:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	FoodId      *string            `json:"food_id" validate:"required"`
	OrderItemId string             `json:"order_item_id" validate:"required"`
	OrderId     string             `json:"order_id" validate:"requiired"`
}