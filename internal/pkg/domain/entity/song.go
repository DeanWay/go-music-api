package entity

import (
	"net/url"

	"github.com/google/uuid"
)

type Song struct {
	Id              uuid.UUID `json:"uuid"`
	Title           string    `json:"title"`
	Artist          string    `json:"artist"`
	DurationSeconds uint      `json:"durationSeconds"`
	AudioFile       url.URL   `json:"audioFile"`
}
