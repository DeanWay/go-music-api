package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type GetAlbumUseCase struct {
	AlbumRepository port.AlbumRepository
}

func (usecase GetAlbumUseCase) GetAlbum(id string) (entity.Album, error) {
	return usecase.AlbumRepository.GetAlbumById(id)
}
