package usecase

import (
	"context"

	"github.com/quachhoang2002/Music-Library/internal/models"
)

func (uc imlUseCase) isTrackExist(ctx context.Context, sc models.Scope, trackID string) bool {
	track, err := uc.musicUC.Detail(ctx, sc, trackID)
	if err != nil {
		return false
	}

	return track.ID.Hex() != ""
}
