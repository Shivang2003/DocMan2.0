package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Document struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string        `bson:"title" json:"title"`
	Content        interface{}   `bson:"content" json:"content"`
	OwnerID        bson.ObjectID `bson:"owner_id" json:"owner_id"`
	CurrentVersion int           `bson:"current_version" json:"current_version"`
	Visibility     string        `bson:"visibility" json:"visibility"`
	CreatedAt      time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at" json:"updated_at"`
}
