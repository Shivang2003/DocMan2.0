package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AuditLog struct {
	ID         bson.ObjectID          `bson:"_id,omitempty" json:"id"`
	UserID     bson.ObjectID          `bson:"user_id" json:"user_id"`
	DocumentID bson.ObjectID          `bson:"document_id" json:"document_id"`
	Action     string                 `bson:"action" json:"action"`
	Metadata   map[string]interface{} `bson:"metadata" json:"metadata"`
	Timestamp  time.Time              `bson:"timestamp" json:"timestamp"`
}
