package port

import (
	"go-music-api/internal/pkg/domain/entity"
)

type HasAddSong interface {
	AddSong(entity.Song) error
}

type HasGetSongById interface {
	GetSongById(id string) (entity.Song, error)
}

type SongRepository interface {
	HasAddSong
	HasGetSongById
}
