package repository

import (
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"

	"github.com/google/uuid"
)

type SongRepository interface {
	AddSong(
		albumUuid uuid.UUID,
		attrs payloads.SongAttributes,
	) models.Song
}
