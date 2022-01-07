package port

import "go-music-api/internal/pkg/domain/entity"

type AlbumSearchParams struct {
	PriceHigh *float64
	PriceLow  *float64
	Artist    *string
	Title     *string
}

type AlbumRepository interface {
	ListAlbums() ([]entity.Album, error)
	AddAlbum(entity.Album) error
	GetAlbumById(id string) (entity.Album, error)
	SearchAlbums(params AlbumSearchParams) ([]entity.Album, error)
}
