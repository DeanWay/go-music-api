package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type GetSongUseCase struct {
	SongRepository port.SongRepository
}

func (usecase GetSongUseCase) GetSong(id string) (entity.Song, error) {
	return usecase.SongRepository.GetSongById(id)
}
