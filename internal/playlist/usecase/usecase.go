package usecase

import (
	"context"

	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/internal/playlist/repository"
	"github.com/xuanhoang/music-library/pkg/mongo"
	"github.com/xuanhoang/music-library/pkg/paginator"

	musicUC "github.com/xuanhoang/music-library/internal/music/usecase"
)

func (uc imlUseCase) Create(ctx context.Context, sc models.Scope, ip CreateInput) (models.Playlist, error) {
	playlist, err := uc.repo.Create(ctx, sc, repository.CreateOpt{
		Name:   ip.Name,
		UserID: ip.UserID,
	})
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.Create.repo.Create", err)
		return models.Playlist{}, err
	}
	return playlist, nil
}

// Err : ErrPlaylistNotFound
func (uc imlUseCase) Update(ctx context.Context, sc models.Scope, ip UpdateInput) (models.Playlist, error) {
	err := uc.repo.Update(ctx, sc, repository.UpdateOpt{
		ID:     ip.ID,
		UserID: ip.UserID,
		Data: repository.UpdateData{
			Name: ip.Data.Name,
		},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Playlist{}, ErrPlaylistNotFound
		}
		uc.l.Error(ctx, "playlist.usecase.Update.repo.Update", err)
		return models.Playlist{}, err
	}

	playlist, err := uc.repo.Detail(ctx, sc, ip.ID)
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.Update.repo.Detail", err)
		return models.Playlist{}, err
	}

	return playlist, nil
}

// Err : ErrPlaylistNotFound
func (uc imlUseCase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.Delete.repo.Delete", err)
		if err == mongo.ErrNoDocuments {
			return ErrPlaylistNotFound
		}
		return err
	}
	return nil
}

// Err : ErrPlaylistNotFound
func (uc imlUseCase) Detail(ctx context.Context, sc models.Scope, id string) (models.Playlist, []models.MusicTrack, error) {
	playlist, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.Detail.repo.Detail", err)
		if err == mongo.ErrNoDocuments {
			return models.Playlist{}, []models.MusicTrack{}, ErrPlaylistNotFound
		}
		return models.Playlist{}, []models.MusicTrack{}, err
	}

	var tracks []models.MusicTrack
	if len(playlist.TrackIDs) > 0 {
		out, err := uc.musicUC.List(ctx, models.Scope{}, musicUC.ListInput{
			PaginatorQuery: paginator.PaginatorQuery{
				Limit: 1000,
				Page:  1,
			},
			Filter: musicUC.Filter{
				IDs: mongo.ObjectIDsToHex(playlist.TrackIDs),
			},
		})
		if err != nil {
			uc.l.Error(ctx, "playlist.usecase.Detail.musicUC.List", err)
			return models.Playlist{}, []models.MusicTrack{}, err
		}
		tracks = out.Tracks
	}

	return playlist, tracks, nil
}

func (uc imlUseCase) List(ctx context.Context, sc models.Scope, ip ListInput) (ListOutput, error) {
	trackIDs := []string{}
	if !ip.TrackFilter.IsEmpty() {
		out, err := uc.musicUC.List(ctx, models.Scope{}, musicUC.ListInput{
			PaginatorQuery: paginator.PaginatorQuery{
				Limit: 1000,
				Page:  1,
			},
			Filter: musicUC.Filter{
				Title:  ip.TrackFilter.Title,
				Artist: ip.TrackFilter.Artist,
				Album:  ip.TrackFilter.Album,
			},
		})
		if err != nil {
			uc.l.Error(ctx, "playlist.usecase.List.musicUC.List", err)
			return ListOutput{}, err
		}

		if len(out.Tracks) == 0 {
			return ListOutput{}, nil
		}

		for _, track := range out.Tracks {
			trackIDs = append(trackIDs, track.ID.Hex())
		}
	}

	playlists, pag, err := uc.repo.List(ctx, sc, repository.ListOpt{
		Filter: repository.Filter{
			TrackIDs: trackIDs,
		},
		PaginatorQuery: ip.PaginatorQuery,
		UserID:         ip.UserID,
	})
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.List.repo.List", err)
		return ListOutput{}, err
	}
	return ListOutput{
		Playlist:  playlists,
		Pagiantor: pag,
	}, nil
}

// Err : ErrTrackNotFound, ErrPlaylistNotFound
func (uc imlUseCase) AddTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error {
	if !uc.isTrackExist(ctx, sc, trackID) {
		return ErrTrackNotFound
	}

	err := uc.repo.AddTrack(ctx, sc, playlistID, trackID)
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.AddTrack.repo.AddTrack", err)
		if err == mongo.ErrNoDocuments {
			return ErrPlaylistNotFound
		}
		return err
	}
	return nil
}

// Err : ErrTrackNotFound, ErrPlaylistNotFound
func (uc imlUseCase) RemoveTrack(ctx context.Context, sc models.Scope, playlistID string, trackID string) error {
	if !uc.isTrackExist(ctx, sc, trackID) {
		return ErrTrackNotFound
	}

	err := uc.repo.RemoveTrack(ctx, sc, playlistID, trackID)
	if err != nil {
		uc.l.Error(ctx, "playlist.usecase.RemoveTrack.repo.RemoveTrack", err)
		if err == mongo.ErrNoDocuments {
			return ErrPlaylistNotFound
		}
		return err
	}
	return nil
}
