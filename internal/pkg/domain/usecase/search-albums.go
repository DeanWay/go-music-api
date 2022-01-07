package usecase

import (
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type SearchAlbumsUseCase struct {
	AlbumRepository port.AlbumRepository
}

func (usecase SearchAlbumsUseCase) SearchAlbums(
	params port.AlbumSearchParams,
) ([]entity.Album, error) {
	return usecase.AlbumRepository.SearchAlbums(params)
}
