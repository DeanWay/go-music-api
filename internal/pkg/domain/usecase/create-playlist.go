package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
	"time"

	"github.com/google/uuid"
)

type CreatePlaylistUseCase struct {
	PlaylistRepository port.HasAddPlaylist
}

func (usecase CreatePlaylistUseCase) CreatePlaylist(
	title string,
	artist string,
	songIds []uuid.UUID,
) (entity.Playlist, error) {
	playlist := entity.Playlist{
		Id:        uuid.New(),
		Title:     title,
		SongIds:   songIds,
		CreatedAt: time.Now().UTC(),
	}
	err := usecase.PlaylistRepository.AddPlaylist(playlist)
	if err != nil {
		return entity.Playlist{}, err
	}
	return playlist, nil
}
