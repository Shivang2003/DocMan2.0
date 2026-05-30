package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `bson:"name" json:"username" validate:"required"`
	Email     string        `bson:"email" json:"email" validate:"required"`
	Password  string        `bson:"password" json:"password" validate:"required"`
	Role      string        `bson:"role" json:"role" validate:"required,oneof=admin user viewer"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at" validate:"required"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at" validate:"required"`
}
