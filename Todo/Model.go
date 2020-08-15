package Todo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	Id        primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Completed bool               `bson:"completed" json:"completed"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
