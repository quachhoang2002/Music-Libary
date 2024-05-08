package http

import (
	"mime/multipart"

	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/internal/music/usecase"
	"github.com/xuanhoang/music-library/pkg/paginator"
)

type detailTrackRes struct {
	ID            string `json:"_id,omitempty"`
	Title         string `json:"title"`
	Artist        string `json:"artist"`
	Album         string `json:"album"`
	Genre         string `json:"genre"`
	ReleaseYear   int    `json:"release_year"`
	Duration      int    `json:"duration"`
	MP3FilePath   string `json:"mp3_file_path"`
	CreatedUserID string `json:"created_user_id"`
}

func newDetailTrackRes(mt models.MusicTrack) detailTrackRes {
	return detailTrackRes{
		ID:            mt.ID.Hex(),
		Title:         mt.Title,
		Artist:        mt.Artist,
		Album:         mt.Album,
		Genre:         mt.Genre,
		ReleaseYear:   mt.ReleaseYear,
		Duration:      mt.Duration,
		MP3FilePath:   mt.MP3FilePath,
		CreatedUserID: mt.CreatedUserID.Hex(),
	}
}

// -- create
type createMusicTrackReq struct {
	Title       string                  `form:"title" binding:"required"`
	Artist      string                  `form:"artist" binding:"required"`
	Album       string                  `form:"album" binding:"required"`
	Genre       string                  `form:"genre" binding:"required"`
	ReleaseYear int                     `form:"release_year" binding:"required"`
	Duration    int                     `form:"duration" binding:"required"`
	MP3Files    []*multipart.FileHeader `form:"mp3_files" binding:"required"`
}

func (r createMusicTrackReq) toInput() usecase.CreateInput {
	return usecase.CreateInput{
		Title:       r.Title,
		Artist:      r.Artist,
		Album:       r.Album,
		Genre:       r.Genre,
		ReleaseYear: r.ReleaseYear,
		Duration:    r.Duration,
		MP3Files:    r.MP3Files,
	}
}

// -- update
type updateMusicTrackReq struct {
	ID          string                  `uri:"id"`
	Name        string                  `form:"name" binding:"required"`
	Aritst      string                  `form:"artist" binding:"required"`
	Album       string                  `form:"album" binding:"required"`
	Genre       string                  `form:"genre" binding:"required"`
	ReleaseYear int                     `form:"release_year" binding:"required"`
	Duration    int                     `form:"duration" binding:"required"`
	MP3Files    []*multipart.FileHeader `form:"mp3_files" binding:"required"`
}

func (r updateMusicTrackReq) toInput() usecase.UpdateInput {
	return usecase.UpdateInput{
		ID: r.ID,
		Data: usecase.UpdateData{
			Title:       r.Name,
			Artist:      r.Aritst,
			Album:       r.Album,
			Genre:       r.Genre,
			ReleaseYear: r.ReleaseYear,
			Duration:    r.Duration,
			MP3Files:    r.MP3Files,
		},
	}
}

// -- list
type listMusicTrackReq struct {
	Title  string `form:"title"`
	Artist string `form:"artist"`
	Album  string `form:"album"`
}

func (r listMusicTrackReq) toInput() usecase.ListInput {
	return usecase.ListInput{
		Filter: usecase.Filter{
			Title:  r.Title,
			Artist: r.Artist,
			Album:  r.Album,
		},
	}
}

type listMusicTrackResp struct {
	Data []itemMusicTrackResp        `json:"data"`
	Meta paginator.PaginatorResponse `json:"meta"`
}

func newListTrackResp(lo usecase.ListOutput) listMusicTrackResp {
	items := make([]itemMusicTrackResp, 0, len(lo.Tracks))
	for _, v := range lo.Tracks {
		items = append(items, newItemMusicTrackResp(v))
	}

	return listMusicTrackResp{
		Data: items,
		Meta: lo.Pagiantor.ToResponse(),
	}
}

type itemMusicTrackResp struct {
	ID          string `json:"_id,omitempty"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
	Duration    int    `json:"duration"`
}

func newItemMusicTrackResp(mt models.MusicTrack) itemMusicTrackResp {
	return itemMusicTrackResp{
		ID:          mt.ID.Hex(),
		Title:       mt.Title,
		Artist:      mt.Artist,
		Album:       mt.Album,
		Genre:       mt.Genre,
		ReleaseYear: mt.ReleaseYear,
		Duration:    mt.Duration,
	}
}
