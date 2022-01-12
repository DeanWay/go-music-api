package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
	"net/url"

	"github.com/google/uuid"
)

type CreateSongUseCase struct {
	SongRepository port.HasAddSong
}

func (usecase CreateSongUseCase) CreateSong(
	title string,
	artist string,
	durationSeconds uint,
	audioFileUrl url.URL,
) (entity.Song, error) {
	song := entity.Song{
		Id:              uuid.New(),
		Title:           title,
		Artist:          artist,
		DurationSeconds: durationSeconds,
		AudioFile:       audioFileUrl,
	}
	err := usecase.SongRepository.AddSong(song)
	if err != nil {
		return entity.Song{}, err
	}
	return song, nil
}
