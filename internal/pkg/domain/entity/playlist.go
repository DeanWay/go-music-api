package entity

import (
	"time"

	"github.com/google/uuid"
)

type Playlist struct {
	Id        uuid.UUID   `json:"uuid"`
	Title     string      `json:"title"`
	SongIds   []uuid.UUID `json:"songIds"`
	CreatedAt time.Time   `json:"createdAt"`
}
