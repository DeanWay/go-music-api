package models

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	Uuid      uuid.UUID `json:"uuid"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
