package repository

import "github.com/quachhoang2002/Music-Library/pkg/paginator"

type Filter struct {
	IDs    []string
	Title  string
	Artist string
	Album  string
}

type ListOpt struct {
	Filter         Filter
	PaginatorQuery paginator.PaginatorQuery
}

type CreateOpt struct {
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int
	MP3FilePath string
}

// -- update

type UpdateOpt struct {
	ID   string
	Data UpdateData
}

type UpdateData struct {
	Title       string
	Artist      string
	Album       string
	Genre       string
	ReleaseYear int
	Duration    int
	MP3FilePath string
}
