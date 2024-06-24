package usecase

import (
	"mime/multipart"

	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/pkg/paginator"
)

type Filter struct {
	IDs    []string
	Title  string
	Artist string
	Album  string
}

type CreateInput struct {
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int
	MP3File     *multipart.FileHeader
}

// update
type UpdateInput struct {
	ID   string
	Data UpdateData
}

type UpdateData struct {
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int // Duration in seconds
	MP3File     *multipart.FileHeader
}

// List
type ListInput struct {
	Filter         Filter
	PaginatorQuery paginator.PaginatorQuery
}

type ListOutput struct {
	Tracks    []models.MusicTrack
	Pagiantor paginator.Paginator
}
