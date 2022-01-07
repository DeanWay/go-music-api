package port

import (
	"go-music-api/internal/pkg/domain/entity"
)

type SongRepository interface {
	AddSong(entity.Song) error
	GetSongById(id string) (entity.Song, error)
}
