package repository

import "github.com/xuanhoang/music-library/pkg/paginator"

type Filter struct {
	TrackIDs []string
}

type ListOpt struct {
	UserID         string
	Filter         Filter
	PaginatorQuery paginator.PaginatorQuery
}

type CreateOpt struct {
	Name   string
	UserID string
}

// -- update

type UpdateOpt struct {
	ID     string
	UserID string
	Data   UpdateData
}

type UpdateData struct {
	Name string
}
