package usecase

import (
	"context"
	"os"

	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/internal/music/repository"
	"github.com/xuanhoang/music-library/pkg/file"
	"github.com/xuanhoang/music-library/pkg/mongo"
)

func (uc imlUseCase) Create(ctx context.Context, sc models.Scope, ip CreateInput) (models.MusicTrack, error) {
	//create file
	filePath, err := file.SaveFile(ip.MP3File, fileStorePath)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.Create.file.SaveFile", err)
		return models.MusicTrack{}, err
	}

	track, err := uc.repo.Create(ctx, sc, repository.CreateOpt{
		Title:       ip.Title,
		Artist:      ip.Artist,
		Album:       ip.Album,
		Genre:       ip.Genre,
		ReleaseYear: ip.ReleaseYear,
		Duration:    ip.Duration,
		MP3FilePath: filePath,
	})
	if err != nil {
		uc.l.Error(ctx, "music.usecase.Create.repo.Create", err)
		return models.MusicTrack{}, err
	}
	return track, nil
}

func (uc imlUseCase) Update(ctx context.Context, sc models.Scope, ip UpdateInput) (models.MusicTrack, error) {
	//create file
	filePath, err := file.SaveFile(ip.Data.MP3File, fileStorePath)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.Create.file.SaveFile", err)
		return models.MusicTrack{}, err
	}

	track, err := uc.repo.Update(ctx, sc, repository.UpdateOpt{
		ID: ip.ID,
		Data: repository.UpdateData{
			Title:       ip.Data.Title,
			Artist:      ip.Data.Artist,
			Album:       ip.Data.Album,
			Genre:       ip.Data.Genre,
			ReleaseYear: ip.Data.ReleaseYear,
			Duration:    ip.Data.Duration,
			MP3FilePath: filePath,
		},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.MusicTrack{}, ErrMusicTrackNotFound
		}
		uc.l.Error(ctx, "music.usecase.Update.repo.Update", err)
		return models.MusicTrack{}, err
	}
	return track, nil
}

// Err : ErrMusicTrackNotFound
func (uc imlUseCase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.Delete.repo.Delete", err)
		if err == mongo.ErrNoDocuments {
			return ErrMusicTrackNotFound
		}
		return err
	}
	return nil
}

// Err : ErrMusicTrackNotFound
func (uc imlUseCase) Detail(ctx context.Context, sc models.Scope, id string) (models.MusicTrack, error) {
	track, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.Detail.repo.Detail", err)
		if err == mongo.ErrNoDocuments {
			return models.MusicTrack{}, ErrMusicTrackNotFound
		}
		return models.MusicTrack{}, err
	}
	return track, nil
}

func (uc imlUseCase) List(ctx context.Context, sc models.Scope, ip ListInput) (ListOutput, error) {
	tracks, pag, err := uc.repo.List(ctx, sc, repository.ListOpt{
		Filter: repository.Filter{
			Title:  ip.Filter.Title,
			Artist: ip.Filter.Artist,
			Album:  ip.Filter.Album,
			IDs:    ip.Filter.IDs,
		},
		PaginatorQuery: ip.PaginatorQuery,
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ListOutput{}, nil
		}
		uc.l.Error(ctx, "music.usecase.List.repo.List", err)
		return ListOutput{}, err
	}
	return ListOutput{
		Tracks:    tracks,
		Pagiantor: pag,
	}, nil
}

func (uc imlUseCase) GetFile(ctx context.Context, sc models.Scope, id string) (*os.File, error) {
	track, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.GetFile.repo.Detail", err)
		if err == mongo.ErrNoDocuments {
			return nil, ErrMusicTrackNotFound
		}
		return nil, err
	}
	file, err := os.Open(track.MP3FilePath)
	if err != nil {
		uc.l.Error(ctx, "music.usecase.GetFile.os.Open", err)
		return nil, err
	}
	return file, nil
}
