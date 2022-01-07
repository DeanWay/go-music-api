package entity

import "github.com/google/uuid"

type Song struct {
	Id              uuid.UUID `json:"uuid"`
	Title           string    `json:"title"`
	Artist          string    `json:"artist"`
	DurationSeconds uint      `json:"durationSeconds"`
}
