package repository

import (
	"context"

	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/pkg/paginator"
)

type Repository interface {
	Create(ctx context.Context, sc models.Scope, opt CreateOpt) (models.MusicTrack, error)
	Update(ctx context.Context, sc models.Scope, opt UpdateOpt) (models.MusicTrack, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
	Detail(ctx context.Context, sc models.Scope, id string) (models.MusicTrack, error)
	List(ctx context.Context, sc models.Scope, opt ListOpt) ([]models.MusicTrack, paginator.Paginator, error)
}
