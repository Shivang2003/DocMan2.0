package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DocumentVersion struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"id"`
	DocumentID    bson.ObjectID `bson:"document_id" json:"document_id"`
	VersionNumber int           `bson:"version_number" json:"version_number"`
	Content       interface{}   `bson:"content" json:"content"`
	UpdatedBy     bson.ObjectID `bson:"updated_by" json:"updated_by"`
	ChangeSummary string        `bson:"change_summary" json:"change_summary"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
}
