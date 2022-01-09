package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type ListAlbumsUseCase struct {
	AlbumRepository port.HasListAlbums
}

func (usecase ListAlbumsUseCase) ListAlbums() ([]entity.Album, error) {
	return usecase.AlbumRepository.ListAlbums()
}
