package usecase

import (
	"context"
	"os"

	"github.com/quachhoang2002/Music-Library/internal/models"
	"github.com/quachhoang2002/Music-Library/internal/music/repository"
	pkgLog "github.com/quachhoang2002/Music-Library/pkg/log"
)

type UseCase interface {
	// Err : ErrMusicTrackNotFound
	Create(ctx context.Context, sc models.Scope, ip CreateInput) (models.MusicTrack, error)
	// Err : ErrMusicTrackNotFound
	Update(ctx context.Context, sc models.Scope, ip UpdateInput) (models.MusicTrack, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
	// Err : ErrMusicTrackNotFound
	Detail(ctx context.Context, sc models.Scope, id string) (models.MusicTrack, error)
	List(ctx context.Context, sc models.Scope, ip ListInput) (ListOutput, error)
	GetFile(ctx context.Context, sc models.Scope, id string) (*os.File, error)
}

type imlUseCase struct {
	l    pkgLog.Logger
	repo repository.Repository
}

func New(l pkgLog.Logger, repo repository.Repository) UseCase {
	return imlUseCase{
		l:    l,
		repo: repo,
	}
}
