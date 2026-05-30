package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Permission struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	DocumentID bson.ObjectID `bson:"document_id" json:"document_id"`
	UserID     bson.ObjectID `bson:"user_id" json:"user_id"`
	Access     string        `bson:"access" json:"access"`
	GrantedBy  bson.ObjectID `bson:"granted_by" json:"granted_by"`
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`
}
