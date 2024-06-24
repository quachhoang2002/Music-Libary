package usecase

import (
	"context"

	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/internal/playlist/repository"
	pkgLog "github.com/quachhoang2002/Music-Library/pkg/log"

	musicUC "github.com/quachhoang2002/Music-Library/internal/music/usecase"
)

type UseCase interface {
	Create(ctx context.Context, sc models.Scope, ip CreateInput) (models.Playlist, error)
	// Err :ErrPlaylistNotFound
	Update(ctx context.Context, sc models.Scope, ip UpdateInput) (models.Playlist, error)
	// Err :ErrPlaylistNotFound
	Delete(ctx context.Context, sc models.Scope, id string) error
	// Err :ErrPlaylistNotFound
	Detail(ctx context.Context, sc models.Scope, id string) (models.Playlist, []models.MusicTrack, error)
	List(ctx context.Context, sc models.Scope, ip ListInput) (ListOutput, error)
	// Err : ErrTrackNotFound, ErrPlaylistNotFound
	AddTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error
	// Err : ErrTrackNotFound, ErrPlaylistNotFound
	RemoveTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error
}

type imlUseCase struct {
	l       pkgLog.Logger
	repo    repository.Repository
	musicUC musicUC.UseCase
}

func New(l pkgLog.Logger, repo repository.Repository, musicUC musicUC.UseCase) UseCase {
	return imlUseCase{
		l:       l,
		repo:    repo,
		musicUC: musicUC,
	}
}
