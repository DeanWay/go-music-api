package models

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	Uuid      uuid.UUID
	Title     string
	Artist    string
	Price     float64
	CreatedAt time.Time
}
