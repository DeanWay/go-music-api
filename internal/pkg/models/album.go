package models

import (
	"go-music-api/internal/pkg/payloads"
	"time"

	"github.com/google/uuid"
)

type Album struct {
	Uuid      uuid.UUID   `json:"uuid"`
	Title     string      `json:"title"`
	Artist    string      `json:"artist"`
	Price     float64     `json:"price"`
	SongUuids []uuid.UUID `json:"songUuids"`
	CreatedAt time.Time   `json:"createdAt"`
}

func NewAlbum(attrs payloads.AlbumAttributes) Album {
	return Album{
		Uuid:      uuid.New(),
		Title:     attrs.Title,
		Artist:    attrs.Artist,
		Price:     attrs.Price,
		CreatedAt: time.Now().UTC(),
	}
}

type Song struct {
	Uuid            uuid.UUID `json:"uuid"`
	AlbumUuid       uuid.UUID `json:"albumUuid"`
	Title           string    `json:"title"`
	DurationSeconds uint      `json:"durationSeconds"`
}

func NewSong(albumUuid uuid.UUID, attrs payloads.SongAttributes) Song {
	return Song{
		Uuid:            uuid.New(),
		AlbumUuid:       albumUuid,
		Title:           attrs.Title,
		DurationSeconds: attrs.DurationSeconds,
	}
}
