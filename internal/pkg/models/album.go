package models

import (
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

type Song struct {
	Uuid            uuid.UUID `json:"uuid"`
	AlbumUuid       uuid.UUID `json:"albumUuid"`
	Title           string    `json:"title"`
	DurationSeconds uint      `json:"durationSeconds"`
}
