package port

import "go-music-api/internal/pkg/domain/entity"

type AlbumSearchParams struct {
	PriceHigh *float64
	PriceLow  *float64
	Artist    *string
	Title     *string
}

type HasListAlbums interface {
	ListAlbums() ([]entity.Album, error)
}

type HasAddAlbum interface {
	AddAlbum(entity.Album) error
}

type HasGetAlbumById interface {
	GetAlbumById(id string) (entity.Album, error)
}

type HasSearchAlbums interface {
	SearchAlbums(params AlbumSearchParams) ([]entity.Album, error)
}
type AlbumRepository interface {
	HasListAlbums
	HasAddAlbum
	HasGetAlbumById
	HasSearchAlbums
}
