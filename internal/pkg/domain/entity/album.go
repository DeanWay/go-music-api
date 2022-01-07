package entity

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	Id        uuid.UUID   `json:"uuid"`
	Title     string      `json:"title"`
	Artist    string      `json:"artist"`
	Price     float64     `json:"price"`
	SongIds   []uuid.UUID `json:"songIds"`
	CreatedAt time.Time   `json:"createdAt"`
}
