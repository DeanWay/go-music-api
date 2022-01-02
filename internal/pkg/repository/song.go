package repository

import (
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"

	"github.com/google/uuid"
)

type SongRepository interface {
	AddSong(
		albumUuid uuid.UUID,
		attrs payloads.SongAttributes,
	) models.Song
}
