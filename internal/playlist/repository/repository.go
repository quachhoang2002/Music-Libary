package repository

import (
	"context"

	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/pkg/paginator"
)

type Repository interface {
	Create(ctx context.Context, sc models.Scope, opt CreateOpt) (models.Playlist, error)
	Update(ctx context.Context, sc models.Scope, opt UpdateOpt) error
	Delete(ctx context.Context, sc models.Scope, id string) error
	Detail(ctx context.Context, sc models.Scope, id string) (models.Playlist, error)
	List(ctx context.Context, sc models.Scope, opt ListOpt) ([]models.Playlist, paginator.Paginator, error)
	AddTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error
	RemoveTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error
}
