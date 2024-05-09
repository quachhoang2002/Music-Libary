package http

import (
	"github.com/xuanhoang/music-library/internal/models"
	"github.com/xuanhoang/music-library/internal/playlist/usecase"
	"github.com/xuanhoang/music-library/pkg/paginator"
)

type detailPlaylistRes struct {
	ID     string           `json:"_id,omitempty"`
	Name   string           `json:"name"`
	Tracks []musicTrackItem `json:"tracks"`
}

func newDetailPlaylistRes(playlist models.Playlist, tracks []models.MusicTrack) detailPlaylistRes {
	trackItems := make([]musicTrackItem, 0, len(tracks))
	for _, v := range tracks {
		trackItems = append(trackItems, newMusicTrackItem(v))
	}

	return detailPlaylistRes{
		ID:     playlist.ID.Hex(),
		Name:   playlist.Name,
		Tracks: trackItems,
	}
}

// -- create
type createPlaylistReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createPlaylistReq) toInput(userID string) usecase.CreateInput {
	return usecase.CreateInput{
		Name:   r.Name,
		UserID: userID,
	}
}

// -- update
type updatePlayListReq struct {
	Name string `json:"name" binding:"required"`
}

func (r updatePlayListReq) toInput(id string, userID string) usecase.UpdateInput {
	return usecase.UpdateInput{
		ID:     id,
		UserID: userID,
		Data: usecase.UpdateData{
			Name: r.Name,
		},
	}
}

// -- list
type listPlaylistReq struct {
	Title  string `form:"title"`
	Artist string `form:"artist"`
	Album  string `form:"album"`
}

type listPlaylistResp struct {
	Data []itemPlaylistResp          `json:"data"`
	Meta paginator.PaginatorResponse `json:"meta"`
}

func newPlaylistsResp(lo usecase.ListOutput) listPlaylistResp {
	items := make([]itemPlaylistResp, 0, len(lo.Playlist))
	for _, v := range lo.Playlist {
		items = append(items, newItemPlaylistResp(v))
	}

	return listPlaylistResp{
		Data: items,
		Meta: lo.Pagiantor.ToResponse(),
	}
}

type itemPlaylistResp struct {
	ID     string `json:"_id,omitempty"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func newItemPlaylistResp(mt models.Playlist) itemPlaylistResp {
	return itemPlaylistResp{
		ID:     mt.ID.Hex(),
		Name:   mt.Name,
		UserID: mt.UserID,
	}
}
