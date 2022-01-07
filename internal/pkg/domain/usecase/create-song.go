package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"

	"github.com/google/uuid"
)

type CreateSongUseCase struct {
	SongRepository port.SongRepository
}

func (usecase CreateSongUseCase) CreateSong(
	title string,
	artist string,
	durationSeconds uint,
) (entity.Song, error) {
	song := entity.Song{
		Id:              uuid.New(),
		Title:           title,
		Artist:          artist,
		DurationSeconds: durationSeconds,
	}
	err := usecase.SongRepository.AddSong(song)
	if err != nil {
		return entity.Song{}, err
	}
	return song, nil
}
