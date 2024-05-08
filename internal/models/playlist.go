package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Name      string               `bson:"name"`
	UserID    primitive.ObjectID   `bson:"user_id"`
	TrackIDs  []primitive.ObjectID `bson:"track_ids"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	DeleteAt  time.Time            `bson:"delete_at,omitempty"`
}
