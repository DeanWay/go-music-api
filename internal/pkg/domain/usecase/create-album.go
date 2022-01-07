package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
	"time"

	"github.com/google/uuid"
)

type CreateAlbumUseCase struct {
	AlbumRepository port.AlbumRepository
}

func (usecase CreateAlbumUseCase) CreateAlbum(
	title string,
	artist string,
	price float64,
	songIds []uuid.UUID,
) (entity.Album, error) {
	album := entity.Album{
		Id:        uuid.New(),
		Title:     title,
		Artist:    artist,
		Price:     price,
		SongIds:   songIds,
		CreatedAt: time.Now().UTC(),
	}
	err := usecase.AlbumRepository.AddAlbum(album)
	if err != nil {
		return entity.Album{}, err
	}
	return album, nil
}
