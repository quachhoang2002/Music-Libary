package usecase

import (
	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/pkg/paginator"
)

type TrackFilter struct {
	Title  string
	Artist string
	Album  string
}

func (f TrackFilter) IsEmpty() bool {
	return f.Title == "" && f.Artist == "" && f.Album == ""
}

type CreateInput struct {
	Name   string
	UserID string
}

// update
type UpdateInput struct {
	ID     string
	UserID string
	Data   UpdateData
}

type UpdateData struct {
	Name string
}

// List
type ListInput struct {
	TrackFilter    TrackFilter
	UserID         string
	PaginatorQuery paginator.PaginatorQuery
}

type ListOutput struct {
	Playlist  []models.Playlist
	Pagiantor paginator.Paginator
}
