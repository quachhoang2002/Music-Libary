package http

import "github.com/quachhoang2002/Music-Library/internal/models"

type musicTrackItem struct {
	ID          string `json:"_id,omitempty"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
	Duration    int    `json:"duration"`
	MP3FilePath string `json:"mp3_file_path"`
}

func newMusicTrackItem(track models.MusicTrack) musicTrackItem {
	return musicTrackItem{
		ID:          track.ID.Hex(),
		Title:       track.Title,
		Artist:      track.Artist,
		Album:       track.Album,
		Genre:       track.Genre,
		ReleaseYear: track.ReleaseYear,
		Duration:    track.Duration,
		MP3FilePath: track.MP3FilePath,
	}
}
