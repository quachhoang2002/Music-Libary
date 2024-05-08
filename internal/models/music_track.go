package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MusicTrack struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title"`
	Artist        string             `bson:"artist"`
	Album         string             `bson:"album"`
	Genre         string             `bson:"genre"`
	ReleaseYear   int                `bson:"release_year"`
	Duration      int                `bson:"duration"` // Duration in seconds
	MP3FilePath   string             `bson:"mp3_file_path"`
	CreatedUserID primitive.ObjectID `bson:"created_user_id"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	DeleteAt      *time.Time         `bson:"delete_at,omitempty"`
}
